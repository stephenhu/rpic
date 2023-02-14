package main

import (
	"fmt"
	"log"
)

func addr() string {
	return fmt.Sprintf("%s:%s", APP_ADDRESS, APP_PORT)
}

func checkParam(params ...string) bool {
	for _, p := range params {
		if p == "" || len(p) == 0 {
			return false
		}
	}

	return true
}

func checkAdmin() {
	if u := getUserByName(APP_ADMIN); u == nil {
		if err := addUser(APP_ADMIN, APP_ADMIN_PASSWORD); err != nil {
			log.Println(err)
			log.Fatal(ERR_USER_ADMIN_NOT_EXIST)
		}
	}
}
