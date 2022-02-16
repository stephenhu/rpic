package main

import (
	"log"
	"net/http"
)


func systemsHandler(w http.ResponseWriter, r *http.Request) {


	u := authenticated(r)

	if u == nil {
		w.WriteHeader(http.StatusUnauthorized)
	} else {

		switch r.Method {
		case http.MethodGet:
		case http.MethodPut:

			m := r.FormValue(PARAM_METHOD)

			if checkLoginMethod(m) {
				
				err := callLogin(m)

				if err != nil {
						
					log.Println(err)
					w.WriteHeader(http.StatusInternalServerError)

				}

			}

		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}

	}

} // systemsHandler
