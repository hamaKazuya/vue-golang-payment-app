package db

import (
	"database/sql"

	// hoge
	_ "github.com/go-sql-driver/mysql"
)

// Conn - sql connection handler
var Conn *sql.DB

// newSQLHandler - init sql handler
func init() {
	user := "root"
	pass := "pwpw"
	name := "itemsDB"

	dbconf := user + ":" + pass + "@/" + name
	conn, err := sql.Open("mysql", dbconf)
	if err != nil {
		panic(err.Error)
	}
	Conn = conn
}
