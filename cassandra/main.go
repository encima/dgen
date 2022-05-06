package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gocql/gocql"
	"github.com/joho/godotenv"
)

func createSchema(session gocql.Session) {
	if err := session.Query(
		"CREATE KEYSPACE IF NOT EXISTS example_keyspace WITH REPLICATION = {'class': 'NetworkTopologyStrategy', 'aiven': 3}",
	).Exec(); err != nil {
		log.Fatal(err)
	}
	if err := session.Query(
		"CREATE TABLE IF NOT EXISTS example_keyspace.example_go (id int PRIMARY KEY, message text)",
	).Exec(); err != nil {
		log.Fatal(err)
	}
}

func writeData(session gocql.Session) {
	for i := 1; i <= 10; i++ {
		if err := session.Query(
			"INSERT INTO example_keyspace.example_go (id, message) VALUES (?, ?)", i, "Hello from golang!",
		).Exec(); err != nil {
			log.Fatal(err)
		}
	}
}

func readData(session gocql.Session) {
	iter := session.Query("SELECT id, message FROM example_keyspace.example_go").Iter()
	var id int
	var message string
	for iter.Scan(&id, &message) {
		fmt.Printf("Row: id = %d, message = %s\n", id, message)
	}
	if err := iter.Close(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	err := godotenv.Load()
	cluster := gocql.NewCluster(os.Getenv("SVC_URI"))
	cluster.Authenticator = gocql.PasswordAuthenticator{
		Username: os.Getenv("SVC_USER"),
		Password: os.Getenv("SVC_PASS"),
	}
	cluster.SslOpts = &gocql.SslOptions{
		CaPath:                 os.Getenv("SVC_CA"),
		EnableHostVerification: false,
	}
	cluster.ProtoVersion = 4
	cluster.Consistency = gocql.Quorum
	session, err := cluster.CreateSession()
	if err != nil {
		log.Fatal(err)
	}
	defer session.Close()

	createSchema(*session)
	writeData(*session)
	readData(*session)

}
