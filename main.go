package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)


func initRouter() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/api/services", servicesHandler)
	router.HandleFunc("/api/services/{name:[0-9a-z]+}", servicesHandler)
	router.HandleFunc("/api/systems", systemHandler)
	router.HandleFunc("/api/version", versionHandler)

	return router

} // initRouter


func main() {

	router := initRouter()

	address := addr()

	log.Printf("%s starting on %s...", APP_NAME, address)
	log.Fatal(http.ListenAndServe(address, router))

} // main
