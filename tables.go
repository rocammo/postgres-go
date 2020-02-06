// Credits: https://www.calhoun.io/connecting-to-a-postgresql-database-with-gos-database-sql-package/
// Docs: http://go-database-sql.org/retrieving.html

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
	fmt.Println("Successfully connected!")

	/* create tables for suppliers database */
	// createTables(db)

	/* select * query from testtable */
	// querySelectAll(db)
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

func createTables(db *sql.DB) {
	commands := []string{`
		CREATE TABLE vendors (
			vendor_id SERIAL PRIMARY KEY,
			vendor_name VARCHAR(255) NOT NULL
		)
		`,
		`
		CREATE TABLE parts (
			part_id SERIAL PRIMARY KEY,
			part_name VARCHAR(255) NOT NULL
		)
		`,
		`
		CREATE TABLE part_drawings (
			part_id INTEGER PRIMARY KEY,
			file_extension VARCHAR(5) NOT NULL,
			drawing_data BYTEA NOT NULL,
			FOREIGN KEY (part_id)
			REFERENCES parts (part_id)
			ON UPDATE CASCADE ON DELETE CASCADE
		)
		`,
		`
		CREATE TABLE vendor_parts (
			vendor_id INTEGER NOT NULL,
			part_id INTEGER NOT NULL,
			PRIMARY KEY (vendor_id , part_id),
			FOREIGN KEY (vendor_id)
			REFERENCES vendors (vendor_id)
			ON UPDATE CASCADE ON DELETE CASCADE,
			FOREIGN KEY (part_id)
			REFERENCES parts (part_id)
			ON UPDATE CASCADE ON DELETE CASCADE
		)
		`}

	for _, command := range commands {
		_, err := db.Exec(command)
		if err != nil {
			panic(err)
		}
	}
}

// func querySelectAll(db *sql.DB) {
// 	rows, err := db.Query("SELECT * FROM testtable")
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer rows.Close()

// 	var (
// 		name, value, updatedt string
// 	)

// 	for rows.Next() {
// 		err := rows.Scan(&name, &value, &updatedt)
// 		if err != nil {
// 			panic(err)
// 		}
// 		fmt.Println(name, value, updatedt)
// 	}
// 	err = rows.Err()
// 	if err != nil {
// 		panic(err)
// 	}
// }
