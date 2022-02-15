package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/stephenhu/webtools"
)


type User struct {
	ID					string								`json:"id"`
	Name      	string								`json:"name"`
	Hash        string        				`json:"hash"`
	Token      	sql.NullString				`json:"token"`
	Created   	string								`json:"created"`
  Updated   	string								`json:"updated"`	
}


func checkToken(c *http.Cookie) *User {

	if c == nil {
		return nil
	}

	clearBuf, err := webtools.Decrypt(c.Value, BLOCK_KEY, IV)

	if err != nil {

		log.Println(err)
		return nil

	} else {

		if len(clearBuf) == 0 {
			return nil
		} else {
	
			user := User{}
	
			err := json.Unmarshal(clearBuf, &user)

			if err != nil {
				
				log.Println(err)
				return nil
	
			} else {
	
				if user.Token.Valid {

					u := getUserByToken(user.Token.String)

					if u == nil {
						return nil
					} else {
						return u
					}
	
				} else {
					return nil
				}

			}
	
		}
	
	}
	
} // checkToken


func authenticate(name string, pass string) *User {

	if name == STR_EMPTY || pass == STR_EMPTY {
		return nil
	} else {

		u := getUserByName(name)

		if u == nil {
			return nil
		} else {

			hash, err := webtools.GenerateHash(pass, SALT, HMAC_KEY,
				SALT2, HASH_LENGTH)
			
			if err != nil {
				
				log.Println(err)
				return nil

			} else {

				if hash == u.Hash {
					return u
				} else {
					return nil
				}
	
			}

		}

	}

} // authenticate


func authHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodDelete:

		cookie := &http.Cookie{
			Name: APP_NAME,
			Value: STR_EMPTY,
			Path: FORWARD_SLASH,
			MaxAge: -1,
		}

		http.SetCookie(w, cookie)

		c, err := r.Cookie(APP_NAME)
		
		if err != nil {
			log.Println(err)
		} else {

			u := checkToken(c)

			if u != nil {
				deleteToken(u)
			}
	
		}


	case http.MethodGet:
  case http.MethodPut:

		user := r.FormValue(PARAM_USER)
		pass := r.FormValue(PARAM_PASS)

		u := authenticate(user, pass)

		if u == nil {
			w.WriteHeader(http.StatusUnauthorized)
		} else {

			token := updateToken(u)

			if token != STR_EMPTY {

				u.Token.String = token
				u.Token.Valid  = true

				j, err := json.Marshal(u)

				if err != nil {
					log.Println(err)
				} else {
	
					encData, err := webtools.Encrypt(j, BLOCK_KEY, IV)
	
					if err != nil {

						log.Println(err)
						w.WriteHeader(http.StatusInternalServerError)

					} else {

						cookie := &http.Cookie{
							Name: APP_NAME,
							Value: encData,
							Path: FORWARD_SLASH,
						}
		
						http.SetCookie(w, cookie)
			
					}

				}

			}

		}

	case http.MethodPost:
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}

} // authHandler
