package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"personalInfoMs/commons"
	"personalInfoMs/models"
	"strconv"

	"github.com/gorilla/mux"
)

func CreateEPS(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newEPS models.EPS
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("Insert a Valid EPS")
		fmt.Fprintf(w, "Insert a Valid EPS")
	}
	json.Unmarshal(reqBody, &newEPS)
	newEPS.IdEPS = 1

	//Consulta DB
	db := commons.ConexionDB()
	insertarRegistro, err := db.Prepare("INSERT INTO `EPS`(`razonSocial`, `estadoEPS`) VALUES (?,?);")
	if err != nil {
		panic(err.Error())
	}
	defer insertarRegistro.Close()

	_, err = insertarRegistro.Exec(newEPS.RazonSocial, newEPS.EstadoEPS)
	if err != nil {
		log.Println("Error al insertar la EPS en la DB")
		panic(err.Error())
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "EPS Creada")
}

func GetEPSs(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") //Le indicamos el tipo de contenido que tiene que procesar

	//Consulta DB
	db := commons.ConexionDB()
	consultarRegistros, err := db.Query("SELECT * FROM EPS;")
	if err != nil {
		panic(err.Error())
	}
	EPS := models.EPS{}
	EPSs := []models.EPS{}

	for consultarRegistros.Next() {
		var idEPS int
		var nombreEPS string
		var estadoEPS bool
		err = consultarRegistros.Scan(&idEPS, &nombreEPS, &estadoEPS)
		if err != nil {
			panic(err.Error())
		}
		EPS.IdEPS = idEPS
		EPS.RazonSocial = nombreEPS
		EPS.EstadoEPS = estadoEPS

		EPSs = append(EPSs, EPS)
	}

	defer consultarRegistros.Close()
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(EPSs) //devolvemos por el w todo el paquete que el cliente desea ver
}

func GetEPS(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") //Le indicamos el tipo de contenido que tiene que procesar
	vars := mux.Vars(r)                                //Extraemos los parametros de la Request

	EPSID, err := strconv.Atoi(vars["id"])
	fmt.Println(EPSID)
	if err != nil {
		fmt.Fprintf(w, "Invalid ID")
		return
	}

	//Consulta DB
	db := commons.ConexionDB()
	consultarRegistro, err := db.Query("SELECT * FROM EPS WHERE idEPS = ? ;", EPSID)
	if err != nil {
		w.WriteHeader(http.StatusConflict)
		log.Fatalf("Invalid ID %v", err)
		fmt.Fprintf(w, "Invalid ID")
		return
	}

	defer consultarRegistro.Close()

	for consultarRegistro.Next() {

		EPS := models.EPS{}
		var idEPS int
		var nombreEPS string
		var estadoEPS bool
		err = consultarRegistro.Scan(&idEPS, &nombreEPS, &estadoEPS)
		if err != nil {
			panic(err.Error())
		}
		EPS.IdEPS = idEPS
		EPS.RazonSocial = nombreEPS
		EPS.EstadoEPS = estadoEPS

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(EPS)
		return
	}
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintf(w, "Ciudad no encontrado")
}

func DeleteEPS(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") //Le indicamos el tipo de contenido que tiene que procesar
	vars := mux.Vars(r)                                //Extraemos los parametros de la Request

	EPSID, err := strconv.Atoi(vars["id"])
	if err != nil {
		json.NewEncoder(w).Encode(false)
		panic(err.Error())
	}

	//Consulta DB
	db := commons.ConexionDB()
	_, err = db.Query("DELETE FROM EPS WHERE idEPS = ?;", EPSID)
	if err != nil {
		json.NewEncoder(w).Encode(false)
		panic(err.Error())
	}

	w.WriteHeader(http.StatusOK)
	//fmt.Fprintf(w, "Ciudad Eliminado")
	json.NewEncoder(w).Encode(true)

	defer db.Close()
}

func UpdateEPS(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r) //Extraemos los parametros de la Request

	EPSID, err := strconv.Atoi(vars["id"])
	// fmt.Println(paisID)
	if err != nil {
		fmt.Fprintf(w, "Invalid ID")
		return
	}

	var newEPS models.EPS
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Insert a Valid EPS")
		return
	}
	json.Unmarshal(reqBody, &newEPS)

	//Consulta DB
	db := commons.ConexionDB()
	actualizarRegistro, err := db.Prepare("UPDATE `EPS` SET `razonSocial`=?, `estadoEPS`=? WHERE idEPS = ?;")
	if err != nil {
		fmt.Fprintf(w, "Error in Prepare Query")
		return
	}

	_, err = actualizarRegistro.Query(newEPS.RazonSocial, newEPS.EstadoEPS, EPSID)
	if err != nil {
		fmt.Fprintf(w, "Error in Prepare Query")
		return
	}
	defer actualizarRegistro.Close()

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "EPS Actualizada")
}
