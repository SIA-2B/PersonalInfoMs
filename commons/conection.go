package commons

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func ConexionDB() (db *sql.DB) {
	driver := "mysql"
	user := "root"
	//contrasenia := "poiuasdoiweq123we1232130asd"
	contrasenia := "hCQf5uePrYH2ErqKI6YJ"
	//nombreDB := "personalInfoDb"
	nombreDB := "railway"

	//direccionIP := "172.17.0.2" //MySQL in  Docker
	direccionIP := "containers-us-west-102.railway.app" //MySQL in  railway
	//direccionIP := "34.95.213.9" //googleSQLCLoud

	port := "6497" //Railway
	//port := "3306"		//Google & Docker

	db, err := sql.Open(driver, user+":"+contrasenia+"@tcp("+direccionIP+":"+port+")/"+nombreDB)
	if err != nil {
		panic(err.Error())
	}

	return db
}
