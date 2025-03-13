package database

import (
  "database/sql"
  "fmt"
  "log"

  _ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func ConnectDBSQL() {

  dsn := "root:123456789@tcp(localhost:3306)/auth"
  var err error;

  DB, err = sql.Open("mysql", dsn);
  if err != nil {
    log.Fatal("Error connecting to database: ",err);
  }

  err = DB.Ping() 
  if err != nil {
    log.Fatal("Database is not reachable: ", err);
  }

  fmt.Println("SQL Database is working !");
}
