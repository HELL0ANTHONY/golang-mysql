package connection

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func DBConnection() (connection *sql.DB) {
	// driver := "mysql"
	// user := "root"
	// pass := "password"
	// dbName := "go_users"

	// connection, err := sql.Open(driver, user+":"+password+"@tcp(172.17.0.2)/"+dbName)
	connection, err := sql.Open("mysql", "root:password@tcp(172.17.0.2)/go_users")
	if err != nil {
		panic(err.Error()) // con .Error() el error es más especifico
	}
	connection.SetConnMaxLifetime(time.Minute * 3)
	connection.SetMaxOpenConns(10)
	connection.SetMaxIdleConns(10)
	return connection
}
