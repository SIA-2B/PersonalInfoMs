package main

import (
	"fmt"
	"net/http"

	"log"
	"personalInfoMs/commons"
	"personalInfoMs/routes"

	"github.com/gorilla/mux"
)

func indexRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to my REST API")

}

func rest() {
	router := mux.NewRouter().StrictSlash(true)
	routes.SetPersonaRoutes(router)
	routes.SetPaisRoutes(router)
	routes.SetCiudadRoutes(router)
	routes.SetEPSRoutes(router)
	routes.SetEtniaRoutes(router)
	router.HandleFunc("/", indexRoute)              // (URL raiz, funion que quiero ejecutar)
	log.Fatal(http.ListenAndServe(":3000", router)) //Port and router
}

func main() {

	go func() {
		for i := 0; i < 3; i++ {
			commons.RabbitMQConsumer()
		}
	}()

	rest()
}
