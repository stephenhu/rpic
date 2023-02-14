package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
)

type User struct {
	ID      string         `json:"id"`
	Name    string         `json:"name"`
	Hash    string         `json:"hash"`
	Token   sql.NullString `json:"token"`
	Created string         `json:"created"`
	Updated string         `json:"updated"`
}

func usersHandler(w http.ResponseWriter, r *http.Request) {
	var u *User

	switch r.Method {
	case http.MethodGet:
		if u = authenticated(r); u == nil {
			w.WriteHeader(http.StatusUnauthorized)
		}

		u.Hash = STR_EMPTY
		u.Token.String = STR_EMPTY
		u.Token.Valid = false

		j, err := json.Marshal(u)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Set(HTTP_CONTENT_TYPE, CONTENT_TYPE_JSON)
		w.Write(j)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
