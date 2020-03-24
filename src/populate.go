package main

import (
	"database/sql"
	"fmt"
)

func populateTables(db *sql.DB) {
	vendors := []string{
		"3M Corp",
		"AKM Semiconductor Inc.",
		"Asahi Glass Co Ltd.",
		"Daikin Industries Ltd.",
		"Dynacast International Inc.",
		"Foster Electric Co. Ltd.",
		"Murata Manufacturing Co. Ltd.",
	}
	populateVendors(db, vendors)

	parts := []string{
		"SIM Tray",
		"Speaker",
		"Vibrator",
		"Antenna",
		"Home Button",
		"LTE Modem",
	}
	populateParts(db, parts)

	partIds := []int{4, 4, 5, 5, 6, 6, 1, 1, 2, 2, 3, 3}
	vendorIds := []int{6, 7, 5, 1, 5, 1, 2, 1, 4, 3, 5, 6}
	mergeVendorsParts(db, partIds, vendorIds)
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

func mergeVendorsParts(db *sql.DB, partIds []int, vendorIds []int) {
	size := 0
	if len(partIds) == len(vendorIds) {
		size = len(partIds)
	} else {
		fmt.Println("Error: size: partIds, vendorIds need to have same length")
	}

	for i := 0; i < size; i++ {
		_, err := db.Exec(`INSERT INTO vendor_parts(part_id, vendor_id) VALUES ($1, $2)`, partIds[i], vendorIds[i])
		if err != nil {
			panic(err)
		}
	}
}
