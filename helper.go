package main

import (
	"errors"	
	"fmt"
	"log"
	"os/exec"	
)


func addr() string {
  return fmt.Sprintf("%s:%s", APP_ADDRESS, APP_PORT)
} // addr


func command(s string) error {

	if s == "" {
		return errors.New("Please input valid command.")
	}

	cmd := exec.Command(s)

	out, err := cmd.Output()

	log.Println(string(out))

	return err

} // command
