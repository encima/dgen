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

	db, err := sql.Open("postgres", os.Getenv("SVC_URI"))
	CheckError(err)

	defer db.Close()
	db.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`)
	db.Exec(`CREATE TABLE IF NOT EXISTS person (
			 id uuid DEFAULT uuid_generate_v1(),
			 email varchar NOT NULL,
			 first_name varchar,
			 last_name varchar,
			 CONSTRAINT person_pkey PRIMARY KEY (id));`)
	fmt.Println("Connected!")

	for {
		size := getdbsize(db)
		if size > 3891776803 {
			break
		}
		db.Exec("insert into person (email, first_name, last_name) select random()::text, random()::text, random()::text from generate_series(1, 80000);")
	}

}

func getdbsize(db *sql.DB) float64 {
	row := db.QueryRow("SELECT pg_database_size('defaultdb');")
	var size float64
	err := row.Scan(&size)
	CheckError(err)
	return size
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
