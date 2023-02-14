package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	_ "modernc.org/sqlite"
)

var (
	data     *sql.DB
	conf     = flag.String("conf", DEFAULT_CONF, "config file path")
	database = flag.String("database", DEFAULT_DATABASE, "database file path")
)

func connectDatabase() {
	var err error
	if _, err = os.Stat(*database); err != nil || os.IsNotExist(err) {
		log.Println(err)
		log.Fatal(ERR_DATABASE_NOT_INITIALIZED)
		// automatically initialize database?
	}

	if data, err = sql.Open(DEFAULT_DATABASE_DRIVER, *database); err != nil {
		log.Println(err)
		log.Fatal(ERR_DATABASE_CONNECTION)
	}
}

func initRouter() *mux.Router {
	router := mux.NewRouter()

	router.PathPrefix(CSS_DIR).Handler(http.FileServer(http.Dir(ROOT_DIR)))
	router.PathPrefix(JS_DIR).Handler(http.FileServer(http.Dir(ROOT_DIR)))

	router.HandleFunc("/", pageHandler)
	router.HandleFunc("/auth", authHandler)
	router.HandleFunc("/api/services", servicesHandler)
	router.HandleFunc("/api/services/{name:[0-9a-z]+}", servicesHandler)
	router.HandleFunc("/api/systems", systemsHandler)
	router.HandleFunc("/api/users", usersHandler)
	router.HandleFunc("/api/version", versionHandler)

	return router
}

func main() {
	flag.Parse()

	if conf == nil || *conf == STR_EMPTY {
		*conf = DEFAULT_CONF
	}

	connectDatabase()
	checkAdmin()

	var (
		router  = initRouter()
		address = addr()
	)

	log.Printf("%s starting on %s...", APP_NAME, address)
	if err := http.ListenAndServe(address, router); err != nil {
		log.Fatal(err)
	}
}
