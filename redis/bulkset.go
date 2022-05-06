package main

import (
	"context"
	"fmt"
	"github.com/encima/dgen/lib"
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
	pp := lib.PromPull{URI: os.Getenv("PROM_URI"), USER: os.Getenv("PROM_USER"), PASS: os.Getenv("PROM_PASS")}
	addr, err := rds.ParseURL(redisURI)
	if err != nil {
		panic(err)
	}
	rdb := rds.NewClient(addr)
	fmt.Println("connected")
	for {
		disk := pp.Pull("disk_used_percent")
		if disk > 80 {
			break
		}
		massImport(rdb)
		fmt.Println("uuids inserted")
	}
	rdb.Close()
}
