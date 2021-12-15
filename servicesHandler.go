package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)


func servicesHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:
	case http.MethodPut:

		vars := mux.Vars(r)

		name := vars[PARAM_NAME]

		if name == "" {
			w.WriteHeader(http.StatusBadRequest)
		} else {

			// TODO: add sqlite3 to store services
			if name == "wireguard" {
				
				err := systemCtl(CMD_WIREGUARD, CMD_OPTION_STOP)

				if err != nil {
					
					log.Println(err)
					w.WriteHeader(http.StatusInternalServerError)

				}

			}

		}

	case http.MethodPost:

		// TODO: add service

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}

} // servicesHandler
