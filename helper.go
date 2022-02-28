package main

import (
	"fmt"
	"log"
	"os"
)


func addr() string {
  return fmt.Sprintf("%s:%s", APP_ADDRESS, APP_PORT)
} // addr


func checkParam(params... string) bool {
	
	for _, p := range params {
		if p == "" || len(p) == 0 {
			return false
		}
	}

	return true

} // checkParam


func checkAdmin() {

	u := getUserByName(APP_ADMIN)

	if u == nil {
		
		err := addUser(APP_ADMIN, APP_ADMIN_PASSWORD)

		if err != nil {
			log.Println(err)
			log.Fatal(ERR_USER_ADMIN_NOT_EXIST)
		}
	}

} // checkAdmin


func fileExists(f string) bool {

	_, err := os.Stat(f)

	if err != nil || os.IsNotExist(err) {
		log.Println(err)
		return false
	}

	return true

} // fileExists
