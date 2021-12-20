package main

const (
	APP_ADDRESS     			= "0.0.0.0"
	APP_NAME							= "rpic"
	APP_PORT        			= "9008"
	APP_VERSION						= "1.0"
)

const (
	PARAM_NAME          	= "name"
	PARAM_OPERATION				= "operation"
)

const (
	OPERATION_LS        	= "ls"
	OPERATION_REBOOT    	= "reboot"
	OPERATION_SHUTDOWN		= "shutdown"
	OPERATION_SUSPEND   	= "suspend"
)

const (
	CMD_LS              	= "dir"
	CMD_REBOOT          	= "reboot"
	CMD_SHUTDOWN					= "poweroff"
	CMD_WIREGUARD       	= "wg-quick@wg0"
	CMD_SYSTEMCTL       	= "systemctl"
)

const (
	CMD_OPTION_RESTART  	= "restart"
	CMD_OPTION_START    	= "start"
	CMD_OPTION_STATUS   	= "status"
	CMD_OPTION_STOP     	= "stop"
	CMD_OPTION_NONE     	= "none"
	CMD_OPTION_IS_ACTIVE	= "is-active"
)

const (
  SYSTEMD								= "org.freedesktop.systemd1"
  SYSTEMD_MANAGE   			= "Manage"
	SYSTEMD_PATH     			= "/org/freedesktop/systemd1"
	SYSTEMD_POWEROFF 			= "poweroff"
	SYSTEMD_REBOOT        = "reboot"
	SYSTEMD_START    			= "StartUnit"
	SYSTEMD_STOP     			= "StopUnit"
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