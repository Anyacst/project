package handler

import (
	"context"
	"database/sql"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"

	"go.com/auth/database"
	"go.com/auth/model"
)

// Register User
func Register(c *gin.Context) {
	var user model.User

	// Parse JSON input
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Password hashing failed"})
		return
	}

	// Store user in MySQL
	_, err = database.DB.Exec("INSERT INTO users (username, passwd_hash) VALUES (?, ?)", user.Username, string(hashedPassword))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database insert failed", "details": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User registered"})
}

// Login User
func Login(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx := context.Background()
	sessionKey := "session:" + user.Username
  
  sessionVal, err := database.Redisclient.Get(ctx, sessionKey).Result()
  if err == nil && sessionVal == "loggedin" { 
      c.JSON(http.StatusOK, gin.H{"message": "Logged in by redis."})
      return
  }

	var storedHash string
	err = database.DB.QueryRow("SELECT passwd_hash FROM users WHERE username = ?", user.Username).Scan(&storedHash)
	if err == sql.ErrNoRows {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error", "details": err.Error()})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(storedHash), []byte(user.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	err = database.Redisclient.Set(ctx, sessionKey, "loggedin", 30*time.Minute).Err()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to set session"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
}

