package main

import (
	"database/sql"
	"errors"
	"log"

	"github.com/stephenhu/webtools"
)


const (

	GET_USER_BY_NAME = "SELECT " +
	  "id, name, hash, token, created, updated " +
		"FROM users " +
		"WHERE name=?"
	
	GET_USER_BY_TOKEN = "SELECT " +
	  "id, name, hash, token, created, updated " +
		"FROM users " +
		"WHERE token=?"

	ADD_USER = "INSERT into users(name, hash) " +
	  "VALUES(?, ?)"

	UPDATE_TOKEN = "UPDATE users " +
	  "SET token=?, updated=CURRENT_TIMESTAMP " +
		"WHERE id=?"
	
	DELETE_TOKEN = "UPDATE users " +
	  "SET token='' " +
		"WHERE id=?"

)


func deleteToken(u *User) {
  
	if u == nil {
		log.Println(ERR_USER_INVALID)
	} else {

		_, err := data.Exec(
			DELETE_TOKEN, u.ID,
		)

		if err != nil {
			log.Println(err)
		}

	}

} // deleteToken


func updateToken(u *User) string {

	if u == nil {
		log.Println(ERR_USER_INVALID)
		return STR_EMPTY
	} else {

		token, err := webtools.GenerateToken(HMAC_KEY, TOKEN_LENGTH)

		if err != nil {
			
			log.Println(err)
			return STR_EMPTY

		} else {

			_, err := data.Exec(
				UPDATE_TOKEN, token, u.ID,
			)

			if err != nil {
				
				log.Println(err)
				return STR_EMPTY

			} else {
				return token
			}

		}

	}

} // updateToken


func getUserByName(name string) *User {

	if name == STR_EMPTY {

		log.Println(ERR_EMPTY_USER_NAME)
		return nil

	} else {

		row := data.QueryRow(
			GET_USER_BY_NAME, name,
		)

		u := User{}

		err := row.Scan(&u.ID, &u.Name, &u.Hash, &u.Token,
			&u.Created, &u.Updated)

		if err != nil || err == sql.ErrNoRows {
			
			log.Println(err)
			return nil

		}

		return &u

	}

} // getUserByName


func getUserByToken(token string) *User {

	if token == STR_EMPTY {

		log.Println(ERR_EMPTY_USER_TOKEN)
		return nil

	} else {

		row := data.QueryRow(
			GET_USER_BY_TOKEN, token,
		)

		u := User{}

		err := row.Scan(&u.ID, &u.Name, &u.Hash, &u.Token, &u.Created, &u.Updated)

		if err != nil || err == sql.ErrNoRows {
			
			log.Println(err)
			return nil

		}

		return &u

	}

} // getUserByToken


func addUser(name string, pass string) error {

	if name == STR_EMPTY {
		return errors.New(ERR_INVALID_USER_NAME)
	}

	if pass == STR_EMPTY {
		return errors.New(ERR_INVALID_PASSWORD)
	}

	hash, err := webtools.GenerateHash(pass, SALT, HMAC_KEY,
		SALT2, HASH_LENGTH)

	if err != nil {

		log.Println(err)
		return err

		} else {

		_, err := data.Exec(
			ADD_USER, name, hash,
		)
	
		if err != nil {
			log.Println(err)
			return err
		}
	
		return nil
	
	}

} // addUser
