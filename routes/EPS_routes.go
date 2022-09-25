package routes

import (
	"personalInfoMs/controllers"

	"github.com/gorilla/mux"
)

func SetEPSRoutes(router *mux.Router) {
	subRoute := router.PathPrefix("/api").Subrouter()
	subRoute.HandleFunc("/EPS/{id}", controllers.GetEPS).Methods("GET")
	subRoute.HandleFunc("/EPSs", controllers.GetEPSs).Methods("GET")
	subRoute.HandleFunc("/EPS", controllers.CreateEPS).Methods("POST")
	subRoute.HandleFunc("/EPS/{id}", controllers.DeleteEPS).Methods("DELETE")
	subRoute.HandleFunc("/EPS/{id}", controllers.UpdateEPS).Methods("PUT")
}
