package main

import (
	"database/sql"
)

func fetchMany(db *sql.DB) {
	command := `SELECT part_name, vendor_name
	FROM parts
	INNER JOIN vendor_parts ON vendor_parts.part_id = parts.part_id
	INNER JOIN vendors ON vendors.vendor_id = vendor_parts.vendor_id
	ORDER BY part_name`

	rows, err := db.Query(command)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	printTable(rows)

	err = rows.Err()
	if err != nil {
		panic(err)
	}
}
