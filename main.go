package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)


func initRouter() *mux.Router {

	router := mux.NewRouter()

  router.PathPrefix(WEB_ASSETS).Handler(http.StripPrefix(WEB_ASSETS,
		http.FileServer(http.Dir(fmt.Sprintf("%s%s", PWD, ROOT_DIR)))))
		
  router.HandleFunc("/", pageHandler)

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
