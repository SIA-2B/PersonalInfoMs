package commons

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func ConexionDB() (db *sql.DB) {
	driver := "mysql"
	user := "root"
	contrasenia := "eO07IOI2MbBo"
	nombreDB := "personalInfoDb"

	//direccionIP := "172.17.0.2" //MySQL in  Docker
	direccionIP := "34.174.47.78" //googleSQLCLoud

	port := "3306" //Google & Docker

	db, err := sql.Open(driver, user+":"+contrasenia+"@tcp("+direccionIP+":"+port+")/"+nombreDB)
	if err != nil {
		panic(err.Error())
	}

	return db
}
