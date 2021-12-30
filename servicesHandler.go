package main

import (
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)


func servicesHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	name := vars[PARAM_NAME]

	switch r.Method {
	case http.MethodGet:

		out, err := getUnitProperty(SERVICE_WIREGUARD, PROPERTY_ACTIVESTATE)

		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
		}

		log.Println(out)

	case http.MethodPut:

		if !checkParam(name) {
			w.WriteHeader(http.StatusBadRequest)
		} else {

			// TODO: add sqlite3 to store services
			if strings.ToLower(name) == WIREGUARD {

        method := r.FormValue(PARAM_METHOD)

        if checkSystemdMethod(method) {
        
				  _, err := callSystemd(method, SERVICE_WIREGUARD)

				  if err != nil {
						
				    log.Println(err)
				    w.WriteHeader(http.StatusInternalServerError)

          }

        } else {

          log.Println("Invalid systemd method")
          w.WriteHeader(http.StatusBadRequest)

				}

			}

		}

	case http.MethodPost:

		// TODO: add service, need to consider security

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}

} // servicesHandler

