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

	/* populate tables for suppliers database */
	populateTables(db)
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

func populateTables(db *sql.DB) {
	// vendors := []string{
	// 	"AKM Semiconductor Inc.",
	// 	"Asahi Glass Co Ltd.",
	// 	"Daikin Industries Ltd.",
	// 	"Dynacast International Inc.",
	// 	"Foster Electric Co. Ltd.",
	// 	"Murata Manufacturing Co. Ltd.",
	// }
	// populateVendors(db, vendors)

	// parts := []string{
	// 	"SIM Tray",
	// 	"Speaker",
	// 	"Vibrator",
	// 	"Antenna",
	// 	"Home Button",
	// 	"LTE Modem",
	// }
	// populateParts(db, parts)

	part_ids := []int{4, 4, 5, 5, 6, 6, 1, 1, 2, 2, 3, 3}
	vendor_ids := []int{6, 7, 5, 1, 5, 1, 2, 1, 4, 3, 5, 6}
	mergeVendorsParts(db, part_ids, vendor_ids)
}

func populateVendors(db *sql.DB, vendors []string) {
	/* Useful SQL statements
	SELECT * FROM vendors;
	DELETE FROM vendors;
	ALTER SEQUENCE vendors_vendor_id_seq RESTART WITH 1;
	*/

	sql := `INSERT INTO vendors(vendor_name) VALUES($1)`

	for _, vendor := range vendors {
		_, err := db.Exec(sql, vendor)
		if err != nil {
			panic(err)
		}
	}
}

func populateParts(db *sql.DB, parts []string) {
	sql := `INSERT INTO parts(part_name) VALUES($1)`

	for _, part := range parts {
		_, err := db.Exec(sql, part)
		if err != nil {
			panic(err)
		}
	}
}

func mergeVendorsParts(db *sql.DB, part_ids []int, vendor_ids []int) {
	size := 0
	if len(part_ids) == len(vendor_ids) {
		size = len(part_ids)
	} else {
		fmt.Println("Error: size: part_ids, vendor_ids need to have same length")
	}

	/* FIXME
		panic: pq: duplicate key value violates unique constraint "vendor_parts_pkey"
	this is caused because both columns are primary key. SOLUTION ?
	*/
	for i := 0; i < size; i++ {
		_, err := db.Exec(`INSERT INTO vendor_parts(part_id, vendor_id) VALUES ($1, $2)`, part_ids[i], vendor_ids[i])
		if err != nil {
			panic(err)
		}
	}
}
