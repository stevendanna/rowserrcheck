package a

import "github.com/jmoiron/sqlx"

func CheckSqlx(db *sqlx.DB) {
	rows, _ := db.Queryx("")
	for rows.Next() {

	}
	_ = rows.Err() // OK
}

func CheckSqlxBad(db *sqlx.DB) {
	rows, _ := db.Queryx("") // want "rows err must be checked"
	for rows.Next() {

	}
}
