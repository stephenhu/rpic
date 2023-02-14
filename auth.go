package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/stephenhu/webtools"
)

func checkToken(c *http.Cookie) (user *User) {
	if c == nil {
		return nil
	}

	clearBuf, err := webtools.Decrypt(c.Value, BLOCK_KEY, IV)
	if err != nil {
		log.Println(err)
		return nil
	}

	if len(clearBuf) == 0 {
		return nil
	}

	user = new(User)
	if err := json.Unmarshal(clearBuf, &user); err != nil {
		log.Println(err)
		return nil
	}

	if !user.Token.Valid {
		return nil
	}

	return getUserByToken(user.Token.String)
}

func authenticate(name string, pass string) *User {
	if name == STR_EMPTY || pass == STR_EMPTY {
		return nil
	}

	u := getUserByName(name)
	if u == nil {
		return nil
	}

	hash, err := webtools.GenerateHash(pass, SALT, HMAC_KEY, SALT2, HASH_LENGTH)
	if err != nil {
		log.Println(err)
		return nil
	}

	if hash == u.Hash {
		return u
	}

	return nil
}

func authenticated(r *http.Request) *User {
	if r == nil {
		return nil
	}

	cookie, err := r.Cookie(APP_NAME)
	if err != nil {
		log.Println(err)
		return nil
	}

	return checkToken(cookie)
}

func authHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodDelete:
		http.SetCookie(w, &http.Cookie{
			Name:   APP_NAME,
			Value:  STR_EMPTY,
			Path:   FORWARD_SLASH,
			MaxAge: -1,
		})

		if u := authenticated(r); u != nil {
			deleteToken(u)
		}
	case http.MethodGet:
	case http.MethodPut:
		var (
			user = r.FormValue(PARAM_USER)
			pass = r.FormValue(PARAM_PASS)
			u    = authenticate(user, pass)
		)

		if u == nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		token := updateToken(u)
		if token != STR_EMPTY {
			u.Token.String = token
			u.Token.Valid = true

			j, err := json.Marshal(u)
			if err != nil {
				log.Println(err)
				return
			}

			encData, err := webtools.Encrypt(j, BLOCK_KEY, IV)
			if err != nil {
				log.Println(err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			http.SetCookie(w, &http.Cookie{
				Name:  APP_NAME,
				Value: encData,
				Path:  FORWARD_SLASH,
			})
		}
	case http.MethodPost:
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
