package routes

import (
	"personalInfoMs/controllers"

	"github.com/gorilla/mux"
)

func SetCiudadRoutes(router *mux.Router) {
	subRoute := router.PathPrefix("/api").Subrouter()
	subRoute.HandleFunc("/ciudad/{id}", controllers.GetCiudad).Methods("GET")
	subRoute.HandleFunc("/ciudades", controllers.GetCiudades).Methods("GET")
	subRoute.HandleFunc("/ciudad", controllers.CreateCiudad).Methods("POST")
	subRoute.HandleFunc("/ciudad/{id}", controllers.DeleteCiudad).Methods("DELETE")
	subRoute.HandleFunc("/ciudad/{id}", controllers.UpdateCiudad).Methods("PUT")
}
