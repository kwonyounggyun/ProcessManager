package dbtask

import (
	"database/sql"
)

type IDBTask interface {
	Execute(db *sql.DB) error
	Complete()
}
