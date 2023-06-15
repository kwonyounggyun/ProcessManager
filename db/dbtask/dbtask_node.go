package dbtask

import (
	"database/sql"
)

type AddNodeTask struct {
	IP       string
	NodeName string

	ID int

	CallFunc func()
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
	t.CallFunc()
}

type Node struct {
	ID   int64  `json:"id"`
	IP   string `json:"ip"`
	Name string `json:"name"`
}

type GetNodesTask struct {
	Nodes []Node

	CallFunc func()
}

func (t *GetNodesTask) Execute(db *sql.DB) error {
	stmt, err := db.Prepare("call Select_Nodes()")
	if err != nil {
		return err
	}

	result, err := stmt.Query()
	if err != nil {
		return err
	}
	defer result.Close()
	for result.Next() {
		var node Node
		err := result.Scan(&node.ID, &node.IP, &node.Name)
		if err != nil {
			return err
		}
		t.Nodes = append(t.Nodes, node)
	}

	return nil
}

func (t *GetNodesTask) Complete() {
	t.CallFunc()
}
