package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/encima/loadgen/lib"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	pp := lib.PromPull{URI: os.Getenv("PROM_URI"), USER: os.Getenv("PROM_USER"), PASS: os.Getenv("PROM_PASS")}
	uri := os.Getenv("SVC_URI")

	db, err := sql.Open("mysql", uri)

	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	pp.Pull()

	// crt, err := db.Prepare("CREATE TABLE person (id int not null primary key auto_increment, email varchar(200), first_name varchar(200), last_name varchar(200));")
	// defer crt.Close()
	// _, err = db.Exec("CREATE TABLE person (id int not null primary key auto_increment, email varchar(200), first_name varchar(200), last_name varchar(200));")
	fmt.Println("Table created")
	// db.Exec(`DELIMITER $$
	//   CREATE PROCEDURE IF NOT EXISTS InsertRand(IN NumRows INT)
	// 	  BEGIN
	// 		  DECLARE i INT;
	// 		  SET i = 1;
	// 		  START TRANSACTION;
	// 		  WHILE i <= NumRows DO
	// 			  INSERT INTO person (email, first_name, last_name) select MD5(RAND()), MD5(RAND()), MD5(RAND());
	// 			  SET i = i + 1;
	// 		  END WHILE;
	// 		  COMMIT;
	// 	  END$$
	//   DELIMITER ;
	//   `)

	if err != nil {
		panic(err.Error())
	}

}
