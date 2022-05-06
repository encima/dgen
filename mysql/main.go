package main

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/encima/dgen/lib"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func insert(db *sql.Stmt) {
	_, err := db.Exec()
	if err != nil {
		panic(err.Error())
	}
	time.Sleep(2000)
}

func main() {
	err := godotenv.Load()
	pp := lib.PromPull{URI: os.Getenv("PROM_URI"), USER: os.Getenv("PROM_USER"), PASS: os.Getenv("PROM_PASS")}
	uri := os.Getenv("SVC_URI")

	db, err := sql.Open("mysql", uri)
	db.SetMaxOpenConns(50)
	db.SetMaxIdleConns(20)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	_, err = db.Exec("CREATE TABL IF NOT EXISTS person (id int not null primary key auto_increment, email varchar(200), first_name varchar(200), last_name varchar(200));")
	fmt.Println("Table created")
	_, err = db.Exec(`DROP PROCEDURE IF EXISTS InsertRand;`)
	_, err = db.Exec(`CREATE PROCEDURE InsertRand(IN NumRows INT)
	BEGIN
	DECLARE i INT;
	SET i = 1;
	START TRANSACTION; 
	WHILE i <= NumRows DO 
		INSERT INTO person (email, first_name, last_name) select MD5(RAND()), MD5(RAND()), MD5(RAND());
	SET i = i + 1;
	END WHILE;
	COMMIT;
	END`)
	if err != nil {
		panic(err.Error())
	}
	for {
		disk := pp.Pull("disk_used_percent")
		if disk > 80 {
			break
		}
		stmtOut, err := db.Prepare("CALL InsertRand(6000);")
		if err != nil {
			panic(err.Error())
		}
		defer stmtOut.Close()
		for i := 0; i < 25; i++ {
			go insert(stmtOut)
		}
	}
}
