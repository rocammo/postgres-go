package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "suppliers"
)

func main() {
	/* connection to the postgres host and suppliers database */
	db := connectDb(host, port, user, password, dbname)
	defer db.Close()

	err := db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected!\n--")

	/* create tables for suppliers database */
	// createTables(db)
	/* populate tables for suppliers database */
	// populateTables(db)

	fetchMany(db)
}

func connectDb(host string, port int, user string, password string, dbname string) *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	return db
}
