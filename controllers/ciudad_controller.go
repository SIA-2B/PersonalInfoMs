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

func CreateCiudad(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newCiudad models.Ciudad
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Insert a Valid City")
	}
	json.Unmarshal(reqBody, &newCiudad)
	newCiudad.IdCiudad = 1

	//Consulta DB
	db := commons.ConexionDB()
	insertarRegistro, err := db.Prepare("INSERT INTO `Ciudad`(`nombreCiudad`, `Pais_idPais`, `estadoCiudad`) VALUES (?,?,?);")
	if err != nil {
		panic(err.Error())
	}
	defer insertarRegistro.Close()

	_, err = insertarRegistro.Exec(newCiudad.NombreCiudad, newCiudad.Pais_idPais, newCiudad.EstadoCiudad)
	if err != nil {
		panic(err.Error())
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Ciudad Creada")
}

func GetCiudades(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") //Le indicamos el tipo de contenido que tiene que procesar

	//Consulta DB
	db := commons.ConexionDB()
	consultarRegistros, err := db.Query("SELECT * FROM Ciudad;")
	if err != nil {
		panic(err.Error())
	}
	Ciudad := models.Ciudad{}
	Ciudades := []models.Ciudad{}

	for consultarRegistros.Next() {
		var idCiudad, Pais_idPais int
		var nombreCiudad string
		var estadoCiudad bool
		err = consultarRegistros.Scan(&idCiudad, &nombreCiudad, &Pais_idPais, &estadoCiudad)
		if err != nil {
			panic(err.Error())
		}
		Ciudad.IdCiudad = idCiudad
		Ciudad.NombreCiudad = nombreCiudad
		Ciudad.Pais_idPais = Pais_idPais
		Ciudad.EstadoCiudad = estadoCiudad

		Ciudades = append(Ciudades, Ciudad)
	}

	defer consultarRegistros.Close()
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Ciudades) //devolvemos por el w todo el paquete que el cliente desea ver
}

func GetCiudad(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") //Le indicamos el tipo de contenido que tiene que procesar
	vars := mux.Vars(r)                                //Extraemos los parametros de la Request

	ciudadID, err := strconv.Atoi(vars["id"])
	fmt.Println(ciudadID)
	if err != nil {
		fmt.Fprintf(w, "Invalid ID")
		return
	}

	//Consulta DB
	db := commons.ConexionDB()
	consultarRegistro, err := db.Query("SELECT * FROM Ciudad WHERE idCiudad = ? ;", ciudadID)
	if err != nil {
		w.WriteHeader(http.StatusConflict)
		log.Fatalf("Invalid ID %v", err)
		fmt.Fprintf(w, "Invalid ID")
		return
	}

	defer consultarRegistro.Close()

	for consultarRegistro.Next() {

		Ciudad := models.Ciudad{}
		var idCiudad, Pais_idPais int
		var nombreCiudad string
		var estadoCiudad bool
		err = consultarRegistro.Scan(&idCiudad, &nombreCiudad, &Pais_idPais, &estadoCiudad)
		if err != nil {
			panic(err.Error())
		}
		Ciudad.IdCiudad = idCiudad
		Ciudad.NombreCiudad = nombreCiudad
		Ciudad.Pais_idPais = Pais_idPais
		Ciudad.EstadoCiudad = estadoCiudad

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(Ciudad)
		return
	}
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintf(w, "Ciudad no encontrado")
}

func DeleteCiudad(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") //Le indicamos el tipo de contenido que tiene que procesar
	vars := mux.Vars(r)                                //Extraemos los parametros de la Request

	ciudadID, err := strconv.Atoi(vars["id"])
	if err != nil {
		panic(err.Error())
	}

	//Consulta DB
	db := commons.ConexionDB()
	_, err = db.Query("DELETE FROM Ciudad WHERE idCiudad = ?;", ciudadID)
	if err != nil {
		panic(err.Error())
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Ciudad Eliminado")

	defer db.Close()
}

func UpdateCiudad(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r) //Extraemos los parametros de la Request

	ciudadID, err := strconv.Atoi(vars["id"])
	// fmt.Println(paisID)
	if err != nil {
		fmt.Fprintf(w, "Invalid ID")
		return
	}

	var newCiudad models.Ciudad
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Insert a Valid Country")
		return
	}
	json.Unmarshal(reqBody, &newCiudad)

	//Consulta DB
	db := commons.ConexionDB()
	actualizarRegistro, err := db.Prepare("UPDATE `Ciudad` SET `nombreCiudad`=?, `Pais_idPais`=?, `estadoCiudad`=? WHERE idPais = ?;")
	if err != nil {
		fmt.Fprintf(w, "Error in Prepare Query")
		return
	}

	_, err = actualizarRegistro.Query(newCiudad.NombreCiudad, newCiudad.Pais_idPais, newCiudad.EstadoCiudad, ciudadID)
	if err != nil {
		fmt.Fprintf(w, "Error in Prepare Query")
		return
	}
	defer actualizarRegistro.Close()

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Ciudad Actualizada")
}
