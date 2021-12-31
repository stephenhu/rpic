package main

import (
	"log"
	"net/http"
)


func systemHandler(w http.ResponseWriter, r *http.Request) {

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

} // systemHandler
