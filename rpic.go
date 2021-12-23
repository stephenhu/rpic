package main

const (
	APP_ADDRESS     			= "0.0.0.0"
	APP_NAME							= "rpic"
	APP_PORT        			= "9008"
	APP_VERSION						= "1.0"
)

const (
	PARAM_NAME          	= "name"
	PARAM_METHOD				  = "method"
)

const (
	ACTION_REBOOT    	    = "reboot"
	ACTION_SHUTDOWN		    = "shutdown"
	ACTION_SUSPEND   	    = "suspend"
)

const (
  DBUS_LOGIN            = "org.freedesktop.login1"
  DBUS_LOGIN_MANAGER    = "org.freedesktop.login1.Manager"
  DBUS_LOGIN_PATH       = "/org/freedesktop/login1"
  DBUS_MANAGER          = "Manager"
  DBUS_SYSTEMD          = "org.freedesktop.systemd1"
  DBUS_SYSTEMD_MANAGER  = "org.freedesktop.systemd1.Manager"
  DBUS_SYSTEMD_PATH     = "/org/freedesktop/systemd1"
)

const (
	LOGIN_POWEROFF 			  = "PowerOff"
	LOGIN_REBOOT          = "Reboot"
)

const (
	SYSTEMD_UNIT_START    = "StartUnit"
	SYSTEMD_UNIT_STOP     = "StopUnit"
)

const (
  WIREGUARD             = "wireguard"
)

const (
  SERVICE_WIREGUARD     = "wg-quick@wg0.service"
)

const (
	HTTP_CONTENT_TYPE			= "Content-Type"
)

const (
	CONTENT_TYPE_JSON			= "application/json"
)

const (
	CH_SPACE							= " "
)

const (
  MAX_TIMEOUT           = 5
)

