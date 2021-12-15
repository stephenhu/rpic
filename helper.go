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


func isValidOption(o string) bool {

	if o == CMD_OPTION_RESTART || o == CMD_OPTION_START ||
	  o == CMD_OPTION_STOP || o == CMD_OPTION_STATUS {
			return true
	} else {
		return false
	}

} // isValidOption


func systemCtl(s string, option string) error {

	if s == "" {
		return errors.New("Please input valid command.")
	}

	if option == CMD_OPTION_NONE {
	
		cmd := exec.Command(CMD_SYSTEMCTL, s)

		out, err := cmd.Output()

		log.Println(string(out))

		return err

	} else {
		
		if isValidOption(option) {

			cmd := exec.Command(CMD_SYSTEMCTL, option, s)

			out, err := cmd.Output()

			log.Println(string(out))

			return err
			
		} else {
			return errors.New(fmt.Sprintf("Invalid option: %s", option))
		}

	}

} // systemCtl
