package main

import (
	"errors"
	"fmt"
	"log"

	"github.com/godbus/dbus/v5"
)

var conn *dbus.Conn

func connectDbus() error {

	c, err := dbus.ConnectSessionBus()

	if err != nil {
		
		log.Println(err)
		return err

	}

	defer c.Close()

	conn = c

	return nil

} // connectDbus


func objectName(service string) string {
	return fmt.Sprintf("%s.%s", service, SYSTEMD_MANAGE)
} // dbusObject


func methodName(service string, name string) string {
	return fmt.Sprintf("%s.%s", objectName(service), name)
} // methodName


func startService() error {
	return callService(SYSTEMD_MANAGE, SYSTEMD_PATH,
		methodName(SYSTEMD_MANAGE, SYSTEMD_START))
} // startService


func stopService() error {
	return callService(SYSTEMD_MANAGE, SYSTEMD_PATH,
		methodName(SYSTEMD_MANAGE, SYSTEMD_STOP))
} // stopService


func checkService() bool {
	return true
} // checkService


func callService(service string, path string, method string) error {

	if len(service) == 0 || len(path) == 0 || len(method) == 0 {
		return errors.New("Invalid dbus call")
	}

	var out string

	obj := conn.Object(service, dbus.ObjectPath(path))

	err := obj.Call(method, 0).Store(&out)

	if err != nil {
	
		log.Println(err)
		return err
	
	}

	return nil

} // callService
