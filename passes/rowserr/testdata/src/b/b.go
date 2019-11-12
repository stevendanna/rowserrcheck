package b

import "database/sql"

func RowsErrCheck(db *sql.DB) {
	rows, _ := db.Query("")
	defer func() {
		err := rows.Err() // ok
		println(err)
	}()

	for rows.Next() {

	}
	// _ = rows.Err() // OK
}

func RowsErrNotCheck(db *sql.DB) {
	rows, _ := db.Query("") // want "rows err must be checked"
	defer func() {
		err := rows.Close()
		println(err)
	}()

	for rows.Next() {

	}
	// _ = rows.Err() // OK
}
