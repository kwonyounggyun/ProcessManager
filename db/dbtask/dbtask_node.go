package dbtask

import "database/sql"

type AddNodeTask struct {
	IP       string
	NodeName string

	ID int
}

func (t *AddNodeTask) Execute(db *sql.DB) error {
	stmt, err := db.Prepare("call Insert_AddNode(?, ?)")
	if err != nil {
		return err
	}

	result := stmt.QueryRow(t.IP, t.NodeName)

	result.Scan(&t.ID)

	return nil
}

func (t *AddNodeTask) Complete() {

}
