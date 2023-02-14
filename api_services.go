package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

func servicesHandler(w http.ResponseWriter, r *http.Request) {
	var (
		vars = mux.Vars(r)
		err  error
		name = vars[PARAM_NAME]
	)

	if authenticated(r) == nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	switch r.Method {
	case http.MethodGet:
		// TODO: check service name
		out, err := getUnitProperty(r.Context(), SERVICE_WIREGUARD, PROPERTY_ACTIVESTATE)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			var j []byte
			if j, err = json.Marshal(Property{Active: out}); err != nil {
				log.Println(err)
			} else {
				w.Write(j)
			}
		}
	case http.MethodPut:
		if !checkParam(name) {
			w.WriteHeader(http.StatusBadRequest)
		} else {
			// TODO: add sqlite3 to store services
			if strings.ToLower(name) == WIREGUARD {
				method := r.FormValue(PARAM_METHOD)
				if checkSystemdMethod(method) {
					if _, err = callSystemd(method, SERVICE_WIREGUARD); err != nil {
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
}
