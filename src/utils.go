package main

import (
	"database/sql"
	"fmt"
)

func printTable(rows *sql.Rows) {
	// rowN := 0
	cols, _ := rows.Columns()
	data := make(map[string]string)

	// for i, col := range cols {
	// 	if i != len(cols)-1 {
	// 		fmt.Printf("%17s | ", col)
	// 	} else {
	// 		fmt.Printf("%s", col)
	// 	}
	// }
	// fmt.Println()
	// for i := 0; i < len(cols); i++ {
	// 	for j := 0; j < len(cols[i]); j++ {
	// 		fmt.Printf("--")
	// 	}
	// 	if i != len(cols)-1 {
	// 		fmt.Printf("+")
	// 	}
	// }
	// fmt.Println()

	for rows.Next() {
		columns := make([]string, len(cols))
		columnPointers := make([]interface{}, len(cols))
		for i, _ := range columns {
			columnPointers[i] = &columns[i]
		}

		rows.Scan(columnPointers...)

		for i, colName := range cols {
			data[colName] = columns[i]
			fmt.Printf("%*s", len(data[colName])+10/2, data[colName])
			if i != len(cols)-1 {
				fmt.Printf(" | ")
			}
			// if i != len(cols)-1 {
			// 	fmt.Printf("%17s | ", data[colName])
			// } else {
			// 	fmt.Printf("%s", data[colName])
			// }
		}
		fmt.Println()

		// rowN++
	}
	// fmt.Printf("(%d rows)\n", rowN)
}
