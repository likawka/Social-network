package queries

import "database/sql"

var DBWrapper *sql.DB

func NewDBWrapper(db *sql.DB) {
	DBWrapper = db
}
