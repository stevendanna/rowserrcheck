package a

import "database/sql"

func f(master_db *sql.DB, slave_db *sql.DB) {

}

func XX() {
	var (
		master_db = new(sql.DB)
		slave_db  = new(sql.DB)
	)

	f(master_db, master_db) // want bad slave
	f(master_db, slave_db)
}
