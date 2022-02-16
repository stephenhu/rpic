package main

import (
	"encoding/json"
	"log"
	"net/http"
)


type App struct {
	Version				string					`json:"version"`
	Name          string          `json:"name"`
}


func versionHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:

		app := App{
			Version:				APP_VERSION,
			Name:						APP_NAME,
		}

		j, err := json.Marshal(app)

		if err != nil {

			log.Printf("versionHandler: %s", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			
		} else {
			w.Header().Set(HTTP_CONTENT_TYPE, CONTENT_TYPE_JSON)
			w.Write(j)
		}

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}

} // versionHandler
