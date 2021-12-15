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
			case OPERATION_REBOOT:
				
				err := systemCtl(CMD_REBOOT, CMD_OPTION_NONE)

				if err != nil {
					
					log.Println(err)
					w.WriteHeader(http.StatusInternalServerError)

				}

			case OPERATION_SHUTDOWN:

				err := systemCtl(CMD_SHUTDOWN, CMD_OPTION_NONE)

				if err != nil {
					
					log.Println(err)
					w.WriteHeader(http.StatusInternalServerError)

				}

			default:
				log.Printf("%s operation not found", operation)
			}

		}

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}

} // systemHandler
