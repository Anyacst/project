package database

import(
  "log"
  "fmt"
  "context"
  "os"

  "github.com/redis/go-redis/v9"
)

var Redisclient *redis.Client

func Redissetup() {

  Redisclient = redis.NewClient(&redis.Options{ Addr : os.Getenv("REDIS_HOST"), })  //"REDIS_HOST : "localhost/6379"

  //Creating the context 
  ctx := context.Background();

  //Ping the redis server
  _, err := Redisclient.Ping(ctx).Result()

  if err != nil {
    log.Fatal("Redis database is not reachabel: ",err);
  }

  fmt.Println("Redis database is Working!");
}
