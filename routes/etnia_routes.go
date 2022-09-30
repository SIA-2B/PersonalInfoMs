package routes

import (
	"personalInfoMs/controllers"

	"github.com/gorilla/mux"
)

func SetEtniaRoutes(router *mux.Router) {
	subRoute := router.PathPrefix("/api").Subrouter()
	subRoute.HandleFunc("/etnia/{id}", controllers.GetEtnia).Methods("GET")
	subRoute.HandleFunc("/etnias", controllers.GetEtnias).Methods("GET")
	subRoute.HandleFunc("/etnia", controllers.CreateEtnia).Methods("POST")
	subRoute.HandleFunc("/etnia/{id}", controllers.DeleteEtnia).Methods("DELETE")
	subRoute.HandleFunc("/etnia/{id}", controllers.UpdateEtnia).Methods("PUT")
}
