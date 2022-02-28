package main

import (
	"database/sql"
	"flag"
	//"fmt"
	"log"
	"net/http"
	"os"

	//_ "github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "modernc.org/sqlite"
	"github.com/gorilla/mux"
)


var data *sql.DB
var appConf Config


var (
	conf        = flag.String("conf", DEFAULT_CONF, "config file path")
	database		= flag.String("database", DEFAULT_DATABASE, "database file path")
)


func connectDatabase() {

	_, err := os.Stat(*database)

	if err != nil || os.IsNotExist(err) {
		
		log.Println(err)
		log.Fatal(ERR_DATABASE_NOT_INITIALIZED)

		// automatically initialize database?
	}

	data, err = sql.Open(DEFAULT_DATABASE_DRIVER, *database)

	if err != nil {

		log.Println(err)
		log.Fatal(ERR_DATABASE_CONNECTION)

	}
	
} // connectDatabase


func initRouter() *mux.Router {

	router := mux.NewRouter()

	router.PathPrefix(CSS_DIR).Handler(http.FileServer(
		http.Dir(ROOT_DIR)))

	router.PathPrefix(JS_DIR).Handler(http.FileServer(
		http.Dir(ROOT_DIR)))
		
  router.HandleFunc("/", pageHandler)

	router.HandleFunc("/auth", authHandler)

	router.HandleFunc("/api/services", servicesHandler)
	router.HandleFunc("/api/services/{name:[0-9a-z]+}", servicesHandler)
	router.HandleFunc("/api/systems", systemsHandler)
	router.HandleFunc("/api/users", usersHandler)
	router.HandleFunc("/api/version", versionHandler)

	return router

} // initRouter


func main() {

	flag.Parse()

	if conf == nil || *conf == STR_EMPTY {
		*conf = DEFAULT_CONF
	}

	//parseConfig()

	connectDatabase()

	checkAdmin()

	router := initRouter()

	address := addr()

	log.Printf("%s starting on %s...", APP_NAME, address)
	log.Fatal(http.ListenAndServe(address, router))

} // main
