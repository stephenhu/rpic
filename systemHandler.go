package main

import (
	"log"
	"net/http"
)


func systemHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:
	case http.MethodPut:

		operation := r.FormValue(PARAM_OPERATION)

		if operation == "" {
			w.WriteHeader(http.StatusBadRequest)
		} else {

			switch(operation) {
			case CMD_LS:
				
				err := command(CMD_LS)

				if err != nil {
					log.Println(err)
				}
			default:
				log.Println("%s operation not found", operation)
			}

		}

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}

} // systemHandler
