package main

import (
	"log"
	"net/http"
)

func systemsHandler(w http.ResponseWriter, r *http.Request) {
	var u *User
	if u = authenticated(r); u == nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	switch r.Method {
	case http.MethodGet:
	case http.MethodPut:
		m := r.FormValue(PARAM_METHOD)
		if checkLoginMethod(m) {
			if err := callLogin(m); err != nil {
				log.Println(err)
				w.WriteHeader(http.StatusInternalServerError)
			}
		}
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
