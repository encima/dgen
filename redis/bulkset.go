package main

import (
	"context"
	"fmt"
	rds "github.com/go-redis/redis/v8"
	"github.com/joho/godotenv"
	"github.com/twinj/uuid"
	"os"
	"runtime"
)

var ctx = context.Background()

const numKeys = 10000

var uuids [numKeys]string

func massImport(client *rds.Client) {
	for i := 0; i < numKeys; i++ {
		uid := uuid.NewV4().String()
		client.Set(ctx, uid, uid, 0)
	}
}

func main() {
	err := godotenv.Load()
	runtime.GOMAXPROCS(16)
	redisURI := os.Getenv("SVC_URI")

	addr, err := rds.ParseURL(redisURI)
	if err != nil {
		panic(err)
	}
	rdb := rds.NewClient(addr)
	fmt.Println("connected")

	massImport(rdb)
	fmt.Println("uuids inserted")
	rdb.Close()
}
