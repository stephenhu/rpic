package main

const (
	APP_ADDRESS     			= "0.0.0.0"
  APP_ADMIN             = "admin"
  APP_ADMIN_PASSWORD    = "rooster2#"
  APP_CONFIG            = "config.json"
	APP_NAME							= "rpic"
	APP_PORT        			= "9008"
	APP_VERSION						= "1.0"
)

const (
  DEFAULT_DATABASE            = "./rpic.db"
  DEFAULT_DATABASE_DRIVER     = "sqlite3"
)


const (
	PARAM_METHOD				  = "method"
	PARAM_NAME          	= "name"
  PARAM_PASS            = "pass"
  PARAM_USER            = "user"
)

const (
  DBUS_LOGIN            = "org.freedesktop.login1"
  DBUS_LOGIN_MANAGER    = "org.freedesktop.login1.Manager"
  DBUS_LOGIN_PATH       = "/org/freedesktop/login1"
  DBUS_MANAGER          = "Manager"
  DBUS_PROPERTIES_GET   = "org.freedesktop.DBus.Properties.Get"
  DBUS_PROPERTIES_GETALL   = "org.freedesktop.DBus.Properties.GetAll"
  DBUS_SYSTEMD          = "org.freedesktop.systemd1"
  DBUS_SYSTEMD_MANAGER  = "org.freedesktop.systemd1.Manager"
	DBUS_SYSTEMD_UNIT  		= "org.freedesktop.systemd1.Unit"
  DBUS_SYSTEMD_PATH     = "/org/freedesktop/systemd1"
  DBUS_UNIT             = "Unit"
)

const (
	LOGIN_POWEROFF 			  = "PowerOff"
	LOGIN_REBOOT          = "Reboot"
)

const (
  SYSTEMD_UNIT_RESTART  = "RestartUnit"
	SYSTEMD_UNIT_START    = "StartUnit"
	SYSTEMD_UNIT_STOP     = "StopUnit"
	SYSTEMD_UNIT_GET      = "GetUnit"
  SYSTEMD_GET           = "Get"
)

const (
  SYSTEMD_UNIT_MODE_REPLACE     = "replace"
)

const (
	PROPERTY_ACTIVESTATE 	= "ActiveState"
)

const (
  UNIT_STATE_ACTIVATING = "activating"
  UNIT_STATE_ACTIVE     = "active"
  UNIT_STATE_FAILED     = "failed"
  UNIT_STATE_INACTIVE   = "inactive"
  UNIT_STATE_RELOADING  = "reloading"
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
	STR_SPACE							= " "
	STR_EMPTY             = ""
  STR_PERIOD            = "."
)

const (
  MAX_TIMEOUT           = 5
)

const (
  CSS_DIR               = "/css"
  FORWARD_SLASH         = "/"
  INDEX_PAGE            = "index"
  JS_DIR                = "/js"
  PWD                   = "."
  ROOT_DIR              = "www"
  WEB_ASSETS            = "/assets/"
)


const (
  BLOCK_KEY             = "0123456789654321"
  HASH_LENGTH           = 32
  HMAC_KEY              = "i love raspberry pi"
  IV                    = "abcdef ghijklmno"
  SALT                  = "random words in a row"
  SALT2                 = "abcdefghijkm"
  TOKEN_LENGTH          = 48
)


const (
  ERR_CONFIG_NOT_FOUND =
    "config.json not found"
  ERR_DATABASE_CONNECTION =
    "Unable to connect to database"
  ERR_DATABASE_NOT_INITIALIZED =
    "Database has not been initialized"
  ERR_EMPTY_USER_NAME =
    "User query cannot search for an empty name"
  ERR_EMPTY_USER_TOKEN =
    "User query cannot search for an empty token"
  ERR_INVALID_USER_NAME =
    "Invalid user name"  
  ERR_USER_INVALID =
    "Inavlid user"
  ERR_INVALID_PASSWORD =
    "Invalid password"
  ERR_USER_ADMIN_NOT_EXIST =
    "admin user does not exist"
)
