package main

const (
	APP_ADDRESS     		= "0.0.0.0"
	APP_NAME						= "rpic"
	APP_PORT        		= "9008"
	APP_VERSION					= "1.0"
)

const (
	PARAM_OPERATION			= "operation"
)

const (
	OPERATION_LS        = "ls"
	OPERATION_REBOOT    = "reboot"
	OPERATION_SHUTDOWN	= "shutdown"
	OPERATION_SUSPEND   = "suspend"
)

const (
	CMD_LS              = "dir"
	CMD_REBOOT          = "systemctl reboot"
	CMD_SHUTDOWN				= "systemctl poweroff"
)

const (
	HTTP_CONTENT_TYPE		= "Content-Type"
)

const (
	CONTENT_TYPE_JSON		= "application/json"
)

const (
	CH_SPACE						= " "
)