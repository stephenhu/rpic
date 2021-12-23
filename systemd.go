package main

import (
  //"context"
	"errors"
	"fmt"
	"log"
  //"time"

  //"github.com/coreos/go-systemd/v22/dbus"
	"github.com/godbus/dbus/v5"
)


func methodName(s string, m string) string {
  return fmt.Sprintf("%s.%s.%s", s, DBUS_MANAGER, m)
} // methodName


func checkLoginMethod(m string) bool {

  if len(m) == 0 || (m != LOGIN_REBOOT &&
    m != LOGIN_POWEROFF) {
    return false
  } else {
    return true
  }

} // checkLoginMethod


func checkSystemdMethod(m string) bool {

  if len(m) == 0 || (m != SYSTEMD_UNIT_START &&
    m != SYSTEMD_UNIT_STOP) {
    return false
  } else {
    return true
  }

} // checkSystemMethod


func callLogin(m string) error {

  if !checkLoginMethod(m) {
    return errors.New("Invalid login method call")
  }

  //ctx, cancel := context.WithTimeout(context.Background(),
  //  MAX_TIMEOUT * time.Second)

  //defer cancel()

  conn, err := dbus.ConnectSystemBus()

  if err != nil {
    log.Println(err)
    return err
  }

  obj := conn.Object(DBUS_LOGIN, dbus.ObjectPath(DBUS_LOGIN_PATH))

  var out string

  switch(m) {

  case LOGIN_POWEROFF:
  
    err := obj.Call(methodName(DBUS_LOGIN, m), 0, m, false).Store(&out)

    if err != nil {
      log.Println(err)
      return err
    }

    log.Println(out)

  case LOGIN_REBOOT:

    err := obj.Call(methodName(DBUS_LOGIN, m), 0, m, false).Store(&out)

    if err != nil {
      log.Println(err)
      return err
    }

  default:
    log.Println("Unknown org.freedesktop.login1 method")

  }

  return nil

} // callLogin


func callSystemd(m string, s string) error {

  if s == "" {
    return errors.New("systemd service name cannot be blank")
  }

  if !checkSystemdMethod(m) {
    return errors.New("Invalid systemd method call")
  }

  //ctx, cancel := context.WithTimeout(context.Background(),
  //  MAX_TIMEOUT * time.Second)

  //defer cancel()

  conn, err := dbus.ConnectSessionBus()

  if err != nil {
    log.Println(err)
    return err
  }

  obj := conn.Object(DBUS_SYSTEMD, dbus.ObjectPath(DBUS_SYSTEMD_PATH))

  var out string

  switch(m) {
  case SYSTEMD_UNIT_START:

    err := obj.Call(methodName(DBUS_SYSTEMD, m), 0, "snapd.service", "replace").Store(&out)

    if err != nil {
      log.Println(err)
      return err
    }

  case SYSTEMD_UNIT_STOP:

    err := obj.Call(methodName(DBUS_SYSTEMD, m), 0, "snapd.service", "replace").Store(&out)

    if err != nil {
      log.Println(err)
      return err
    }


  default:
    log.Println("Unknown org.freedesktop.systemd1 method")

  }

  return nil

} // callSystemd

