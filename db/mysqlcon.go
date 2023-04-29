package db

import (
	"database/sql"
	"fmt"
	"log"
)

type MySQLHelper struct {
	db *sql.DB
}

func (my *MySQLHelper) Connect(ip, id, pw, db_name string, port int) {
	con_str := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", id, pw, ip, port, db_name)
	var err error
	my.db, err = sql.Open("mysql", con_str)
	if err != nil {
		log.Fatal(err)
	}
	defer my.db.Close()

	// // 하나의 Row를 갖는 SQL 쿼리
	// var name string
	// err = my.db.QueryRow("SELECT name FROM test1 WHERE id = 1").Scan(&name)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(name)
}

func (my *MySQLHelper) CreatePrepare(sql string) (*sql.Stmt, error) {
	stmt, err := my.db.Prepare(sql)

	if err != nil {
		return nil, err
	}

	return stmt, nil
}

func (my *MySQLHelper) Close() {
	my.db.Close()
}
