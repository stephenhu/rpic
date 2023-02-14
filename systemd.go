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
	Active string `json:"active"`
}

func methodName(s string, m string) string {
	return fmt.Sprintf("%s.%s", s, m)
}

func checkLoginMethod(m string) bool {
	var (
		s = strings.ToLower(m)
		r = strings.ToLower(LOGIN_REBOOT)
		p = strings.ToLower(LOGIN_POWEROFF)
	)

	if len(m) == 0 || (s != r && s != p) {
		return false
	} else {
		return true
	}
}

func checkSystemdMethod(m string) bool {
	var (
		s = strings.ToLower(m)
		r = strings.ToLower(SYSTEMD_UNIT_RESTART)
		t = strings.ToLower(SYSTEMD_UNIT_START)
		p = strings.ToLower(SYSTEMD_UNIT_STOP)
		g = strings.ToLower(SYSTEMD_UNIT_GET)
	)

	if len(m) == 0 || (s != r && s != t && s != p && s != g) {
		return false
	} else {
		return true
	}
}

func callLogin(m string) (err error) {
	var (
		conn   *dbus.Conn
		out    string
		method = strings.ToLower(m)
		power  = strings.ToLower(LOGIN_POWEROFF)
		reboot = strings.ToLower(LOGIN_REBOOT)
	)

	if !checkLoginMethod(m) {
		return errors.New("invalid login method call")
	}

	if conn, err = dbus.ConnectSystemBus(); err != nil {
		log.Println(err)
		return
	}

	obj := conn.Object(DBUS_LOGIN, dbus.ObjectPath(DBUS_LOGIN_PATH))
	ctx, cancelCtx := context.WithTimeout(context.Background(), MAX_TIMEOUT*time.Second)
	defer cancelCtx()

	switch method {
	case power:
		if err = obj.CallWithContext(ctx, methodName(DBUS_LOGIN_MANAGER, LOGIN_POWEROFF), 0, false).Store(&out); err != nil {
			log.Println(err)
			return
		}

		log.Println(out)
	case reboot:
		if err = obj.CallWithContext(ctx, methodName(DBUS_LOGIN_MANAGER, LOGIN_REBOOT), 0, false).Store(&out); err != nil {
			log.Println(err)
			return
		}
	default:
		log.Println("Unknown org.freedesktop.login1 method")
	}

	return
}

func callSystemd(m string, s string) (string, error) {
	var (
		conn    *dbus.Conn
		call    *dbus.Call
		out     string
		err     error
		method  = strings.ToLower(m)
		restart = strings.ToLower(SYSTEMD_UNIT_RESTART)
		start   = strings.ToLower(SYSTEMD_UNIT_START)
		stop    = strings.ToLower(SYSTEMD_UNIT_STOP)
		get     = strings.ToLower(SYSTEMD_UNIT_GET)
	)

	if !checkParam(m, s) {
		return STR_EMPTY, errors.New("systemd service name cannot be blank")
	}

	if !checkSystemdMethod(m) {
		return STR_EMPTY, errors.New("invalid systemd method call")
	}

	if conn, err = dbus.ConnectSystemBus(); err != nil {
		return STR_EMPTY, err
	}

	ctx, cancelCtx := context.WithTimeout(context.Background(), MAX_TIMEOUT*time.Second)
	defer cancelCtx()

	obj := conn.Object(DBUS_SYSTEMD, dbus.ObjectPath(DBUS_SYSTEMD_PATH))
	switch method {
	case restart:
		call = obj.CallWithContext(ctx, methodName(DBUS_SYSTEMD_MANAGER, SYSTEMD_UNIT_RESTART), 0, s, SYSTEMD_UNIT_MODE_REPLACE)
		if err = call.Store(&out); err != nil {
			return STR_EMPTY, err
		}
	case start:
		call = obj.CallWithContext(ctx, methodName(DBUS_SYSTEMD_MANAGER, SYSTEMD_UNIT_START), 0, s, SYSTEMD_UNIT_MODE_REPLACE)
		if err = call.Store(&out); err != nil {
			return STR_EMPTY, err
		}
	case stop:
		call = obj.CallWithContext(ctx, methodName(DBUS_SYSTEMD_MANAGER, SYSTEMD_UNIT_STOP), 0, s, SYSTEMD_UNIT_MODE_REPLACE)
		if err = call.Store(&out); err != nil {
			return STR_EMPTY, err
		}
	case get:
		if err = obj.CallWithContext(ctx, methodName(DBUS_SYSTEMD_MANAGER, SYSTEMD_UNIT_GET), 0, s).Store(&out); err != nil {
			return STR_EMPTY, err
		}
		return out, nil
	default:
		log.Println("Unknown org.freedesktop.systemd1 method")
	}

	return STR_EMPTY, nil
}

func getUnitProperty(ctx context.Context, service string, property string) (res string, err error) {
	var (
		out, path string
		conn      *dbus.Conn
	)

	if !checkParam(service, property) {
		return res, fmt.Errorf("invalid parameter: (%s) (%s)", service, property)
	}

	if path, err = callSystemd(SYSTEMD_UNIT_GET, service); err != nil {
		return
	}

	if conn, err = dbus.ConnectSystemBus(); err != nil {
		return
	}

	ctx, cancelCtx := context.WithTimeout(ctx, MAX_TIMEOUT*time.Second)
	defer cancelCtx()

	obj := conn.Object(DBUS_SYSTEMD, dbus.ObjectPath(path))
	call := obj.CallWithContext(ctx, DBUS_PROPERTIES_GET, 0, DBUS_SYSTEMD_UNIT, property)
	if err = call.Store(&out); err != nil {
		return
	}

	return
}
