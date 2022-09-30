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

func CreateEtnia(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newEtnia models.Etnia
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Insert a Valid Etnia")
	}
	json.Unmarshal(reqBody, &newEtnia)
	newEtnia.IdEtnia = 1

	//Consulta DB
	db := commons.ConexionDB()
	insertarRegistro, err := db.Prepare("INSERT INTO `Etnia`(`categoriaEtnica`, `estadoEtnia`) VALUES (?,?);")
	if err != nil {
		panic(err.Error())
	}
	defer insertarRegistro.Close()

	_, err = insertarRegistro.Exec(newEtnia.CategoriaEtnica, newEtnia.EstadoEtnia)
	if err != nil {
		panic(err.Error())
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Etnia Creada")
}

func GetEtnias(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") //Le indicamos el tipo de contenido que tiene que procesar

	//Consulta DB
	db := commons.ConexionDB()
	consultarRegistros, err := db.Query("SELECT * FROM Etnia;")
	if err != nil {
		panic(err.Error())
	}
	Etnia := models.Etnia{}
	Etnias := []models.Etnia{}

	for consultarRegistros.Next() {
		var idEtnia int
		var categoriaEtnica string
		var estadoEtnia bool
		err = consultarRegistros.Scan(&idEtnia, &categoriaEtnica, &estadoEtnia)
		if err != nil {
			panic(err.Error())
		}
		Etnia.IdEtnia = idEtnia
		Etnia.CategoriaEtnica = categoriaEtnica
		Etnia.EstadoEtnia = estadoEtnia

		Etnias = append(Etnias, Etnia)
	}

	defer consultarRegistros.Close()
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Etnias) //devolvemos por el w todo el paquete que el cliente desea ver
}

func GetEtnia(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") //Le indicamos el tipo de contenido que tiene que procesar
	vars := mux.Vars(r)                                //Extraemos los parametros de la Request

	EtniaID, err := strconv.Atoi(vars["id"])
	fmt.Println(EtniaID)
	if err != nil {
		fmt.Fprintf(w, "Invalid ID")
		return
	}

	//Consulta DB
	db := commons.ConexionDB()
	consultarRegistro, err := db.Query("SELECT * FROM Etnia WHERE idEtnia = ? ;", EtniaID)
	if err != nil {
		w.WriteHeader(http.StatusConflict)
		log.Fatalf("Invalid ID %v", err)
		fmt.Fprintf(w, "Invalid ID")
		return
	}

	defer consultarRegistro.Close()

	for consultarRegistro.Next() {

		Etnia := models.Etnia{}
		var idEtnia int
		var categoriaEtnica string
		var estadoEtnia bool
		err = consultarRegistro.Scan(&idEtnia, &categoriaEtnica, &estadoEtnia)
		if err != nil {
			panic(err.Error())
		}
		Etnia.IdEtnia = idEtnia
		Etnia.CategoriaEtnica = categoriaEtnica
		Etnia.EstadoEtnia = estadoEtnia

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(Etnia)
		return
	}
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintf(w, "Etnia no encontrada")
}

func DeleteEtnia(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") //Le indicamos el tipo de contenido que tiene que procesar
	vars := mux.Vars(r)                                //Extraemos los parametros de la Request

	EtniaID, err := strconv.Atoi(vars["id"])
	if err != nil {
		panic(err.Error())
	}

	//Consulta DB
	db := commons.ConexionDB()
	_, err = db.Query("DELETE FROM Etnia WHERE idEtnia = ?;", EtniaID)
	if err != nil {
		panic(err.Error())
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Etnia Eliminada")

	defer db.Close()
}

func UpdateEtnia(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r) //Extraemos los parametros de la Request

	EtniaID, err := strconv.Atoi(vars["id"])
	// fmt.Println(paisID)
	if err != nil {
		fmt.Fprintf(w, "Invalid ID")
		return
	}

	var newEtnia models.Etnia
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Insert a Valid Etnia")
		return
	}
	json.Unmarshal(reqBody, &newEtnia)

	//Consulta DB
	db := commons.ConexionDB()
	actualizarRegistro, err := db.Prepare("UPDATE `Etnia` SET `categoriaEtnica`=?, `estadoEtnia`=? WHERE idEtnia = ?;")
	if err != nil {
		fmt.Fprintf(w, "Error in Prepare Query")
		return
	}

	_, err = actualizarRegistro.Query(newEtnia.CategoriaEtnica, newEtnia.EstadoEtnia, EtniaID)
	if err != nil {
		fmt.Fprintf(w, "Error in Prepare Query")
		return
	}
	defer actualizarRegistro.Close()

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Etnia Actualizada")
}
