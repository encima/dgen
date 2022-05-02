package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	err := godotenv.Load()
	uri := os.Getenv("SVC_URI")

	db, err := sql.Open("postgres", uri)
	CheckError(err)

	// close database
	defer db.Close()

	res, err := db.Query("SELECT pg_database_size('testdb');")
	CheckError(err)
	fmt.Println(res)
	fmt.Println("Connected!")
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
