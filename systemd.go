package main

import (
  "context"
	"errors"
	"fmt"
	"log"
  "strings"
  "time"

	"github.com/godbus/dbus/v5"
)


type Property struct {
  Active string       `json:"active"`
}


func methodName(s string, m string) string {
  return fmt.Sprintf("%s.%s", s, m)
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
  g := strings.ToLower(SYSTEMD_UNIT_GET)
  
  if len(m) == 0 || (s != r && s != t && s != p && s != g) {
    return false
  } else {
    return true
  }

} // checkSystemMethod


func callLogin(m string) error {

  if !checkLoginMethod(m) {
    return errors.New("Invalid login method call")
  }

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

  ctx, cancel := context.WithTimeout(context.Background(),
  MAX_TIMEOUT * time.Second)

  defer cancel()

  switch(method) {

  case power:
  
    err := obj.CallWithContext(ctx, methodName(DBUS_LOGIN_MANAGER,
      LOGIN_POWEROFF), 0, false).Store(&out)

    if err != nil {
      log.Println(err)
      return err
    }

    log.Println(out)

  case reboot:

    err := obj.CallWithContext(ctx, methodName(DBUS_LOGIN_MANAGER,
      LOGIN_REBOOT), 0, false).Store(&out)

    if err != nil {
      log.Println(err)
      return err
    }

  default:
    log.Println("Unknown org.freedesktop.login1 method")

  }

  return nil

} // callLogin


func callSystemd(m string, s string) (string, error) {

  if !checkParam(m, s) {
    return STR_EMPTY, errors.New("systemd service name cannot be blank")
  }

  if !checkSystemdMethod(m) {
    return STR_EMPTY, errors.New("Invalid systemd method call")
  }

  conn, err := dbus.ConnectSystemBus()

  if err != nil {
    log.Println(err)
    return STR_EMPTY, err
  }

  obj := conn.Object(DBUS_SYSTEMD, dbus.ObjectPath(DBUS_SYSTEMD_PATH))

  var out string

  method    := strings.ToLower(m)
  restart   := strings.ToLower(SYSTEMD_UNIT_RESTART)
  start     := strings.ToLower(SYSTEMD_UNIT_START)
  stop      := strings.ToLower(SYSTEMD_UNIT_STOP)
  get       := strings.ToLower(SYSTEMD_UNIT_GET)

  ctx, cancel := context.WithTimeout(context.Background(),
  MAX_TIMEOUT * time.Second)

  defer cancel()

  switch(method) {
  case restart:

    err := obj.CallWithContext(ctx, methodName(DBUS_SYSTEMD_MANAGER,
      SYSTEMD_UNIT_RESTART), 0, s, SYSTEMD_UNIT_MODE_REPLACE).Store(&out)

    if err != nil {
      log.Println(err)
      return STR_EMPTY, err
    }

  case start:

    err := obj.CallWithContext(ctx, methodName(DBUS_SYSTEMD_MANAGER,
      SYSTEMD_UNIT_START), 0, s, SYSTEMD_UNIT_MODE_REPLACE).Store(&out)

    if err != nil {
      log.Println(err)
      return STR_EMPTY, err
    }

  case stop:

    err := obj.CallWithContext(ctx, methodName(DBUS_SYSTEMD_MANAGER,
      SYSTEMD_UNIT_STOP), 0, s, SYSTEMD_UNIT_MODE_REPLACE).Store(&out)

    if err != nil {
      log.Println(err)
      return STR_EMPTY, err
    }

  case get:

    err := obj.CallWithContext(ctx, methodName(DBUS_SYSTEMD_MANAGER,
      SYSTEMD_UNIT_GET), 0, s).Store(&out)

    if err != nil {
      log.Println(err)
      return STR_EMPTY, err
    }

    return out, nil

  default:
    log.Println("Unknown org.freedesktop.systemd1 method")

  }

  return STR_EMPTY, nil

} // callSystemd


func getUnitProperty(s string, p string) (string, error) {

  if !checkParam(s, p) {
    return STR_EMPTY, errors.New(fmt.Sprintf(
      "Invalid parameter: (%s) (%s)", s, p))
  }

  path, err := callSystemd(SYSTEMD_UNIT_GET, s)

  if err != nil {
    log.Println(err)
    return STR_EMPTY, err
  }

  conn, err := dbus.ConnectSystemBus()

  if err != nil {
    log.Println(err)
    return STR_EMPTY, err
  }

  ctx, cancel := context.WithTimeout(context.Background(),
    MAX_TIMEOUT * time.Second)

  defer cancel()

  obj := conn.Object(DBUS_SYSTEMD, dbus.ObjectPath(path))

  var out string

  err = obj.CallWithContext(ctx, DBUS_PROPERTIES_GET, 0,
    DBUS_SYSTEMD_UNIT, p).Store(&out)

  if err != nil {
    log.Println(err)
    return STR_EMPTY, err
  }

  return out, nil

} // getUnitProperty

