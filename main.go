package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)


func initRouter() *mux.Router {

	router := mux.NewRouter()

  router.PathPrefix(ROOT_DIR).Handler(http.StripPrefix(ROOT_DIR,
		http.FileServer(http.Dir(fmt.Sprintf("%s%s/*", PWD, ROOT_DIR)))))

	router.PathPrefix(CSS_DIR).Handler(http.FileServer(
		http.Dir(ROOT_DIR)))

	router.PathPrefix(JS_DIR).Handler(http.FileServer(
		http.Dir(ROOT_DIR)))
		
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
