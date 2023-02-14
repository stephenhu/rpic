package main

import (
	"database/sql"
	"errors"
	"log"

	"github.com/stephenhu/webtools"
)

const (
	GET_USER_BY_NAME  = `SELECT id, name, hash, token, created, updated FROM users WHERE name=?`
	GET_USER_BY_TOKEN = `SELECT id, name, hash, token, created, updated FROM users WHERE token=?`
	ADD_USER          = `INSERT into users(name, hash) VALUES(?, ?)`
	UPDATE_TOKEN      = `UPDATE users SET token=?, updated=CURRENT_TIMESTAMP WHERE id=?`
	DELETE_TOKEN      = `UPDATE users SET token='' WHERE id=?`
)

func deleteToken(user *User) {
	if user == nil {
		log.Println(ERR_USER_INVALID)
		return
	}

	if _, err := data.Exec(DELETE_TOKEN, user.ID); err != nil {
		log.Println(err)
	}
}

func updateToken(user *User) string {
	if user == nil {
		log.Println(ERR_USER_INVALID)
		return STR_EMPTY
	}

	var (
		token string
		err   error
	)

	if token, err = webtools.GenerateToken(HMAC_KEY, TOKEN_LENGTH); err != nil {
		return STR_EMPTY
	}

	if _, err = data.Exec(UPDATE_TOKEN, token, user.ID); err != nil {
		return STR_EMPTY
	}

	return token
}

func getUserByName(name string) (user *User) {
	if name == STR_EMPTY {
		log.Println(ERR_EMPTY_USER_NAME)
		return
	}

	user = new(User)
	if err := data.QueryRow(GET_USER_BY_NAME, name).Scan(
		&user.ID,
		&user.Name,
		&user.Hash,
		&user.Token,
		&user.Created,
		&user.Updated,
	); err != nil || err == sql.ErrNoRows {
		log.Println(err)
		return nil
	}

	return
}

func getUserByToken(token string) (user *User) {
	user = new(User)

	if token == STR_EMPTY {
		log.Println(ERR_EMPTY_USER_TOKEN)
		return
	}

	if err := data.QueryRow(GET_USER_BY_TOKEN, token).Scan(
		&user.ID,
		&user.Name,
		&user.Hash,
		&user.Token,
		&user.Created,
		&user.Updated,
	); err != nil || err == sql.ErrNoRows {
		log.Println(err)
		return nil
	}

	return
}

func addUser(name string, pass string) (err error) {
	if name == STR_EMPTY {
		return errors.New(ERR_INVALID_USER_NAME)
	}

	if pass == STR_EMPTY {
		return errors.New(ERR_INVALID_PASSWORD)
	}

	var hash string
	if hash, err = webtools.GenerateHash(pass, SALT, HMAC_KEY, SALT2, HASH_LENGTH); err != nil {
		log.Println(err)
		return err
	}

	if _, err := data.Exec(ADD_USER, name, hash); err != nil {
		log.Println(err)
		return err
	}

	return nil
}
