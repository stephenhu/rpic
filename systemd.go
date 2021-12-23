package main

import (
  //"context"
	"errors"
	"fmt"
	"log"
  "strings"
  //"time"

	"github.com/godbus/dbus/v5"
)


func methodName(s string, m string) string {
  return fmt.Sprintf("%s.%s.%s", s, DBUS_MANAGER, m)
} // methodName


func checkLoginMethod(m string) bool {

  s := strings.ToLower(m)
  r := strings.ToLower(LOGIN_REBOOT)
  p := strings.ToLower(LOGIN_POWEROFF)

  if len(m) == 0 || (s != r && s != p) {
    return false
  } else {
    return true
  }

} // checkLoginMethod


func checkSystemdMethod(m string) bool {

  s := strings.ToLower(m)
  r := strings.ToLower(SYSTEMD_UNIT_RESTART)
  t := strings.ToLower(SYSTEMD_UNIT_START)
  p := strings.ToLower(SYSTEMD_UNIT_STOP)
  
  if len(m) == 0 || (s != r && s != t && s != p) {
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

  method := strings.ToLower(m)
  power  := strings.ToLower(LOGIN_POWEROFF)
  reboot := strings.ToLower(LOGIN_REBOOT)

  switch(method) {

  case power:
  
    err := obj.Call(methodName(DBUS_LOGIN, LOGIN_POWEROFF), 0, false).Store(&out)

    if err != nil {
      log.Println(err)
      return err
    }

    log.Println(out)

  case reboot:

    err := obj.Call(methodName(DBUS_LOGIN, LOGIN_REBOOT), 0, false).Store(&out)

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

  conn, err := dbus.ConnectSystemBus()

  if err != nil {
    log.Println(err)
    return err
  }

  obj := conn.Object(DBUS_SYSTEMD, dbus.ObjectPath(DBUS_SYSTEMD_PATH))

  var out string

  method    := strings.ToLower(m)
  restart   := strings.ToLower(SYSTEMD_UNIT_RESTART)
  start     := strings.ToLower(SYSTEMD_UNIT_START)
  stop      := strings.ToLower(SYSTEMD_UNIT_STOP)

  switch(method) {
  case restart:

    err := obj.Call(methodName(DBUS_SYSTEMD, SYSTEMD_UNIT_RESTART), 0, s,
      "replace").Store(&out)

    if err != nil {
      log.Println(err)
      return err
    }

  case start:

    err := obj.Call(methodName(DBUS_SYSTEMD, SYSTEMD_UNIT_START), 0, s,
      "replace").Store(&out)

    if err != nil {
      log.Println(err)
      return err
    }

  case stop:

    err := obj.Call(methodName(DBUS_SYSTEMD, SYSTEMD_UNIT_STOP), 0, s,
      "replace").Store(&out)

    if err != nil {
      log.Println(err)
      return err
    }


  default:
    log.Println("Unknown org.freedesktop.systemd1 method")

  }

  return nil

} // callSystemd

