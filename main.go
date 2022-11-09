package main

import (
	"fmt"
	"log"
	"net/http"
	"personalInfoMs/commons"
	"personalInfoMs/routes"
	"strconv"
	"time"

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

		f := true
		i := 1
		for f {
			log.Println("Estableciendo conexión con RabbitMQ... (Intento " + strconv.Itoa(i) + ")")
			commons.RabbitMQConsumer()
			time.Sleep(5 * time.Second)
			i++
		}
	}()

	rest()
}
