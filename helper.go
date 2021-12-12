package main

import (
	"fmt"
	"log"
	"os/exec"
)


func addr() string {
  return fmt.Sprintf("%s:%s", APP_ADDRESS, APP_PORT)
} // addr


func command(s string) error {

	cmd := exec.Command(s)

	out, err := cmd.Output()

	if err != nil {
		return err
	} else {
		
		log.Println(out)
		return nil

	}

} // command
