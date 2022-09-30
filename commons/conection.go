package commons

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func ConexionDB() (db *sql.DB) {

	db, err := sql.Open("mysql", "root:poiuasdoiweq123we1232130asd@tcp(172.17.0.2:3306)/personalInfoDb")

	if err != nil {
		panic(err.Error())
	}

	return db
}
