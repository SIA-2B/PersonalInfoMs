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

func CreatePersona(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newPersona models.Persona
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Insert a Valid Person")
	}
	json.Unmarshal(reqBody, &newPersona)
	newPersona.IDPersona = 1
	// tasks = append(tasks, newPersona)

	//Consulta DB
	db := commons.ConexionDB()
	insertarRegistro, err := db.Prepare("INSERT INTO `Persona`(`nombrePersona`, `apellidoPersona`, `tipoDocumento`, `NUIPPersona`, `lugarNacimiento_idCiudad`, `lugarExpDocumento_idCiudad`, `estadoCivil`, `sexoBio`, `Etnia_idEtnia`, `correoPersonal`, `telefonoMovil`, `telefonoFijo`, `fechaNacimiento`, `EPS_idEPS`, `grupoSangre`, `nivelAcademico`, `factorRH`, `direccionResidencia`, `lugarResidencia_idCiudad`, `estratoSocioeconomico`, `libretaMilitar`, `fechaRegistroSistema`, `estadoPersona`) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
	if err != nil {
		panic(err.Error())
	}
	defer insertarRegistro.Close()

	_, err = insertarRegistro.Exec(newPersona.NombrePersona, newPersona.ApellidoPersona, newPersona.TipoDocumento, newPersona.NUIPPersona, newPersona.LugarNacimiento, newPersona.LugarExpDocumento, newPersona.EstadoCivil, newPersona.SexoBio, newPersona.Etnia, newPersona.CorreoPersonal, newPersona.TelefonoMovil, newPersona.TelefonoFijo, newPersona.FechaNacimiento, newPersona.EPS, newPersona.GrupoSangre, newPersona.NivelAcademico, newPersona.FactorRH, newPersona.DirResidencia, newPersona.LugarResidencia, newPersona.EstratoSocioeconomico, newPersona.LibretaMilitar, newPersona.FechaRegistroSistema, newPersona.EstadoPersona)
	if err != nil {
		panic(err.Error())
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Persona Creada")
}

func GetPersonas(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") //Le indicamos el tipo de contenido que tiene que procesar

	//Consulta DB
	db := commons.ConexionDB()
	consultarRegistros, err := db.Query("SELECT * FROM Persona;")
	if err != nil {
		panic(err.Error())
	}
	persona := models.Persona{}
	personas := []models.Persona{}

	for consultarRegistros.Next() {
		var idPersona, NUIPPersona, lugarNacimiento, lugarExpDocumento, estadoCivil, etnia, telefonoMovil, telefonoFijo, EPS, lugarResidencia, estratoSocioeconomico int
		var nombrePersona, apellidoPersona, tipoDocumento, sexoBio, correoPersonal, fechaNacimiento, grupoSangre, nivelAcademico, factorRH, dirResidencia, fechaRegistroSistema string
		var libretaMilitar, estadoPersona bool
		err = consultarRegistros.Scan(&idPersona, &nombrePersona, &apellidoPersona, &tipoDocumento, &NUIPPersona, &lugarNacimiento, &lugarExpDocumento, &estadoCivil, &sexoBio, &etnia, &correoPersonal, &telefonoMovil, &telefonoFijo, &fechaNacimiento, &EPS, &grupoSangre, &nivelAcademico, &factorRH, &dirResidencia, &lugarResidencia, &estratoSocioeconomico, &libretaMilitar, &fechaRegistroSistema, &estadoPersona)
		if err != nil {
			panic(err.Error())
		}
		persona.IDPersona = idPersona
		persona.NombrePersona = nombrePersona
		persona.ApellidoPersona = apellidoPersona
		persona.TipoDocumento = tipoDocumento
		persona.NUIPPersona = NUIPPersona
		persona.LugarNacimiento = lugarNacimiento
		persona.LugarExpDocumento = lugarExpDocumento
		persona.EstadoCivil = estadoCivil
		persona.SexoBio = sexoBio
		persona.Etnia = etnia
		persona.CorreoPersonal = correoPersonal
		persona.TelefonoMovil = telefonoMovil
		persona.TelefonoFijo = telefonoFijo
		persona.FechaNacimiento = fechaNacimiento
		persona.EPS = EPS
		persona.GrupoSangre = grupoSangre
		persona.NivelAcademico = nivelAcademico
		persona.FactorRH = factorRH
		persona.DirResidencia = dirResidencia
		persona.LugarResidencia = lugarResidencia
		persona.EstratoSocioeconomico = estratoSocioeconomico
		persona.LibretaMilitar = libretaMilitar
		persona.FechaRegistroSistema = fechaRegistroSistema
		persona.EstadoPersona = estadoPersona

		personas = append(personas, persona)
	}

	//fmt.Println(personas)

	defer consultarRegistros.Close()
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(personas) //devolvemos por el w todo el paquete que el cliente desea ver
}

func GetPersona(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") //Le indicamos el tipo de contenido que tiene que procesar
	vars := mux.Vars(r)                                //Extraemos los parametros de la Request

	personaID, err := strconv.Atoi(vars["id"])
	fmt.Println(personaID)
	if err != nil {
		fmt.Fprintf(w, "Invalid ID")
		return
	}

	//Consulta DB
	db := commons.ConexionDB()
	consultarRegistro, err := db.Query("SELECT * FROM Persona WHERE idPersona = ? ;", personaID)
	if err != nil {
		w.WriteHeader(http.StatusConflict)
		log.Fatalf("Invalid ID %v", err)
		fmt.Fprintf(w, "Invalid ID")
		return
	}

	defer consultarRegistro.Close()

	for consultarRegistro.Next() {

		persona := models.Persona{}
		var idPersona, NUIPPersona, lugarNacimiento, lugarExpDocumento, estadoCivil, etnia, telefonoMovil, telefonoFijo, EPS, lugarResidencia, estratoSocioeconomico int
		var nombrePersona, apellidoPersona, tipoDocumento, sexoBio, correoPersonal, fechaNacimiento, grupoSangre, nivelAcademico, factorRH, dirResidencia, fechaRegistroSistema string
		var libretaMilitar, estadoPersona bool

		err = consultarRegistro.Scan(&idPersona, &nombrePersona, &apellidoPersona, &tipoDocumento, &NUIPPersona, &lugarNacimiento, &lugarExpDocumento, &estadoCivil, &sexoBio, &etnia, &correoPersonal, &telefonoMovil, &telefonoFijo, &fechaNacimiento, &EPS, &grupoSangre, &nivelAcademico, &factorRH, &dirResidencia, &lugarResidencia, &estratoSocioeconomico, &libretaMilitar, &fechaRegistroSistema, &estadoPersona)
		if err != nil {
			w.WriteHeader(http.StatusConflict)
			fmt.Fprintf(w, "Error al convertir de SQL Row a Objeto GO")
			return
		}

		persona.IDPersona = idPersona
		persona.NombrePersona = nombrePersona
		persona.ApellidoPersona = apellidoPersona
		persona.TipoDocumento = tipoDocumento
		persona.NUIPPersona = NUIPPersona
		persona.LugarNacimiento = lugarNacimiento
		persona.LugarExpDocumento = lugarExpDocumento
		persona.EstadoCivil = estadoCivil
		persona.SexoBio = sexoBio
		persona.Etnia = etnia
		persona.CorreoPersonal = correoPersonal
		persona.TelefonoMovil = telefonoMovil
		persona.TelefonoFijo = telefonoFijo
		persona.FechaNacimiento = fechaNacimiento
		persona.EPS = EPS
		persona.GrupoSangre = grupoSangre
		persona.NivelAcademico = nivelAcademico
		persona.FactorRH = factorRH
		persona.DirResidencia = dirResidencia
		persona.LugarResidencia = lugarResidencia
		persona.EstratoSocioeconomico = estratoSocioeconomico
		persona.LibretaMilitar = libretaMilitar
		persona.FechaRegistroSistema = fechaRegistroSistema
		persona.EstadoPersona = estadoPersona

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(persona) //devolvemos por el w la Persona que el cliente desea ver
		return
	}
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintf(w, "Persona no encontrada")
}

func DeletePersona(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") //Le indicamos el tipo de contenido que tiene que procesar
	vars := mux.Vars(r)                                //Extraemos los parametros de la Request

	personaID, err := strconv.Atoi(vars["id"])
	fmt.Println(personaID)
	if err != nil {
		fmt.Fprintf(w, "Invalid ID")
		return
	}

	//Consulta DB
	db := commons.ConexionDB()
	eliminarRegistro, err := db.Query("DELETE FROM Persona WHERE idPersona = ? ;", personaID)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Persona no encontrada")
		return
	}

	defer eliminarRegistro.Close()
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Persona Eliminada")
}

func UpdatePersona(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r) //Extraemos los parametros de la Request

	personaID, err := strconv.Atoi(vars["id"])
	fmt.Println(personaID)
	if err != nil {
		fmt.Fprintf(w, "Invalid ID")
		return
	}

	var newPersona models.Persona
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Insert a Valid Person")
	}
	json.Unmarshal(reqBody, &newPersona)

	//Consulta DB
	db := commons.ConexionDB()
	actualizarRegistro, err := db.Prepare("UPDATE `Persona` SET `nombrePersona`=?,`apellidoPersona`=?,`tipoDocumento`=?,`NUIPPersona`=?,`lugarNacimiento_idCiudad`=?,`lugarExpDocumento_idCiudad`=?,`estadoCivil`=?,`sexoBio`=?,`Etnia_idEtnia`=?,`correoPersonal`=?,`telefonoMovil`=?,`telefonoFijo`=?,`fechaNacimiento`=?,`EPS_idEPS`=?,`grupoSangre`=?,`nivelAcademico`=?,`factorRH`=?,`direccionResidencia`=?,`lugarResidencia_idCiudad`=?,`estratoSocioeconomico`=?,`libretaMilitar`=?,`fechaRegistroSistema`=?,`estadoPersona`=? WHERE idPersona = ?")
	if err != nil {
		panic(err.Error())
	}
	defer actualizarRegistro.Close()

	_, err = actualizarRegistro.Exec(newPersona.NombrePersona, newPersona.ApellidoPersona, newPersona.TipoDocumento, newPersona.NUIPPersona, newPersona.LugarNacimiento, newPersona.LugarExpDocumento, newPersona.EstadoCivil, newPersona.SexoBio, newPersona.Etnia, newPersona.CorreoPersonal, newPersona.TelefonoMovil, newPersona.TelefonoFijo, newPersona.FechaNacimiento, newPersona.EPS, newPersona.GrupoSangre, newPersona.NivelAcademico, newPersona.FactorRH, newPersona.DirResidencia, newPersona.LugarResidencia, newPersona.EstratoSocioeconomico, newPersona.LibretaMilitar, newPersona.FechaRegistroSistema, newPersona.EstadoPersona, personaID)
	if err != nil {
		panic(err.Error())
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Persona Actualizada")
}

func UploadPhoto(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") //Le indicamos el tipo de contenido que tiene que procesar
	vars := mux.Vars(r)                                //Extraemos los parametros de la Request

	var IDName int = -1

	personaID, err := strconv.Atoi(vars["id"])
	fmt.Println(personaID)
	if err != nil {
		fmt.Fprintf(w, "Invalid ID")
		return
	}

	//Consulta DB
	db := commons.ConexionDB()
	consultarRegistro, err := db.Query("SELECT idPersona FROM Persona WHERE idPersona = ? ;", personaID)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Not Found")
		return
	}

	defer consultarRegistro.Close()

	for consultarRegistro.Next() {
		var idPersona int

		err = consultarRegistro.Scan(&idPersona)
		if err != nil {
			w.WriteHeader(http.StatusConflict)
			fmt.Fprintf(w, "Error al convertir de SQL Row a Objeto GO")
			return
		}

		if idPersona != 0 {
			IDName = idPersona
		} else {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "Persona no encontrada")
			return
		}

	}

	if IDName == -1 {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Persona no encontrada")
		return
	}

	file, _, err := r.FormFile("myFile")
	if err != nil {
		log.Printf("Error al cargar el archivo %v", err)
		fmt.Fprintf(w, "Error al cargar el archivo %v", err)
		return
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Printf("Error al leer el archivo %v", err)
		fmt.Fprintf(w, "Error al leer el archivo %v", err)
		return
	}

	err = ioutil.WriteFile("./files/images/UserPictures/"+strconv.Itoa(IDName)+".png", data, 0666)
	if err != nil {
		log.Printf("Error al escribir el archivo %v", err)
		fmt.Fprintf(w, "Error al escribir el archivo %v", err)
		return
	}

	fmt.Fprintf(w, "Imagen Cargada exitosamente")
}
