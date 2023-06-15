package dbtask

import "database/sql"

type LoginTask struct {
	ID string
	PW string

	KEY int

	CallFunc func()
}

func (t *LoginTask) Execute(db *sql.DB) error {
	stmt, err := db.Prepare("call Select_User(?, ?)")
	if err != nil {
		return err
	}

	result := stmt.QueryRow(t.ID, t.PW)

	result.Scan(&t.KEY)

	return nil
}

func (t *LoginTask) Complete() {
	t.CallFunc()
}
