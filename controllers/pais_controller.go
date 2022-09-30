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

func CreatePais(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newPais models.Pais
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Insert a Valid Country")
	}
	json.Unmarshal(reqBody, &newPais)
	newPais.IdPais = 1

	//Consulta DB
	db := commons.ConexionDB()
	insertarRegistro, err := db.Prepare("INSERT INTO `Pais`(`nombrePais`, `estadoPais`) VALUES (?,?);")
	if err != nil {
		panic(err.Error())
	}
	defer insertarRegistro.Close()

	_, err = insertarRegistro.Exec(newPais.NombrePais, newPais.EstadoPais)
	if err != nil {
		panic(err.Error())
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Pais Creado")
}

func GetPaises(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") //Le indicamos el tipo de contenido que tiene que procesar

	//Consulta DB
	db := commons.ConexionDB()
	consultarRegistros, err := db.Query("SELECT * FROM Pais;")
	if err != nil {
		panic(err.Error())
	}
	Pais := models.Pais{}
	Paises := []models.Pais{}

	for consultarRegistros.Next() {
		var idPais int
		var nombrePais string
		var estadoPais bool
		err = consultarRegistros.Scan(&idPais, &nombrePais, &estadoPais)
		if err != nil {
			panic(err.Error())
		}
		Pais.IdPais = idPais
		Pais.NombrePais = nombrePais
		Pais.EstadoPais = estadoPais

		Paises = append(Paises, Pais)
	}

	//fmt.Println(personas)

	defer consultarRegistros.Close()
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Paises) //devolvemos por el w todo el paquete que el cliente desea ver
}

func GetPais(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") //Le indicamos el tipo de contenido que tiene que procesar
	vars := mux.Vars(r)                                //Extraemos los parametros de la Request

	paisID, err := strconv.Atoi(vars["id"])
	fmt.Println(paisID)
	if err != nil {
		fmt.Fprintf(w, "Invalid ID")
		return
	}

	//Consulta DB
	db := commons.ConexionDB()
	consultarRegistro, err := db.Query("SELECT * FROM Pais WHERE idPais = ? ;", paisID)
	if err != nil {
		w.WriteHeader(http.StatusConflict)
		log.Fatalf("Invalid ID %v", err)
		fmt.Fprintf(w, "Invalid ID")
		return
	}

	defer consultarRegistro.Close()

	for consultarRegistro.Next() {

		pais := models.Pais{}
		var idPais int
		var nombrePais string
		var estadoPais bool

		err = consultarRegistro.Scan(&idPais, &nombrePais, &estadoPais)
		if err != nil {
			w.WriteHeader(http.StatusConflict)
			fmt.Fprintf(w, "Error al convertir de SQL Row a Objeto GO")
			return
		}

		pais.IdPais = idPais
		pais.NombrePais = nombrePais
		pais.EstadoPais = estadoPais

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(pais)
		return
	}
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintf(w, "Pais no encontrado")
}

func DeletePais(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") //Le indicamos el tipo de contenido que tiene que procesar
	vars := mux.Vars(r)                                //Extraemos los parametros de la Request

	paisID, err := strconv.Atoi(vars["id"])
	if err != nil {
		panic(err.Error())
	}

	//Consulta DB
	db := commons.ConexionDB()
	_, err = db.Query("DELETE FROM Pais WHERE idPais = ?;", paisID)
	if err != nil {
		panic(err.Error())
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Pais Eliminado")

	defer db.Close()
}

func UpdatePais(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r) //Extraemos los parametros de la Request

	paisID, err := strconv.Atoi(vars["id"])
	// fmt.Println(paisID)
	if err != nil {
		fmt.Fprintf(w, "Invalid ID")
		return
	}

	var newPais models.Pais
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Insert a Valid Country")
		return
	}
	json.Unmarshal(reqBody, &newPais)

	//Consulta DB
	db := commons.ConexionDB()
	actualizarRegistro, err := db.Prepare("UPDATE `Pais` SET `nombrePais`=?,`estadoPais`=? WHERE idPais = ?;")
	if err != nil {
		fmt.Fprintf(w, "Error in Prepare Query")
		return
	}

	_, err = actualizarRegistro.Query(newPais.NombrePais, newPais.EstadoPais, paisID)
	if err != nil {
		fmt.Fprintf(w, "Error in Prepare Query")
		return
	}
	defer actualizarRegistro.Close()

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Pais Actualizado")
}
