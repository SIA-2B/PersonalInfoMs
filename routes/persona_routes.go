package routes

import (
	"personalInfoMs/controllers"

	"github.com/gorilla/mux"
)

func SetPersonaRoutes(router *mux.Router) {
	subRoute := router.PathPrefix("/api").Subrouter()
	subRoute.HandleFunc("/persona/{id}", controllers.GetPersona).Methods("GET")
	subRoute.HandleFunc("/persona/user/{username}", controllers.GetPersonaByUsername).Methods("GET")
	subRoute.HandleFunc("/persona/nuip/{nuip}", controllers.GetPersonaByNUIP).Methods("GET")
	subRoute.HandleFunc("/personas", controllers.GetPersonas).Methods("GET")
	subRoute.HandleFunc("/persona", controllers.CreatePersona).Methods("POST")
	subRoute.HandleFunc("/persona/{id}", controllers.DeletePersona).Methods("DELETE")
	subRoute.HandleFunc("/persona/{id}", controllers.UpdatePersona).Methods("PUT")
	subRoute.HandleFunc("/persona/file/{id}", controllers.UploadPhoto).Methods("POST")
}
