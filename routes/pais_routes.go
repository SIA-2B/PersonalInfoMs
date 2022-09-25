package routes

import (
	"personalInfoMs/controllers"

	"github.com/gorilla/mux"
)

func SetPaisRoutes(router *mux.Router) {
	subRoute := router.PathPrefix("/api").Subrouter()
	subRoute.HandleFunc("/pais/{id}", controllers.GetPais).Methods("GET")
	subRoute.HandleFunc("/paises", controllers.GetPaises).Methods("GET")
	subRoute.HandleFunc("/pais", controllers.CreatePais).Methods("POST")
	subRoute.HandleFunc("/pais/{id}", controllers.DeletePais).Methods("DELETE")
	subRoute.HandleFunc("/pais/{id}", controllers.UpdatePais).Methods("PUT")
}
