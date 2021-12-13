package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"syscall"
)


func addr() string {
  return fmt.Sprintf("%s:%s", APP_ADDRESS, APP_PORT)
} // addr


func command(s string) error {

	if s == "" {
		return errors.New("Invalid command, please enter a valid command.")
	}

	args := strings.Split(s, CH_SPACE)

	if len(args) == 0 {
		return errors.New("Invalid command, please enter a valid command.")
	}

	index := 1

	if len(args) == 1 {
		index = 0
	}

	bin, err := exec.LookPath(args[index])

	if err != nil {
		return err
	} else {

		env := os.Environ()

		err := syscall.Exec(bin, args, env)

		if err != nil {
			return err	
		}
		
		return nil

	}

} // command
