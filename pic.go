package main

const (
	APP_ADDRESS     		= "127.0.0.1"
	APP_NAME						= "raspberry pi controller"
	APP_PORT        		= "12000"
	APP_VERSION					= "1.0"
)

const (
	PARAM_OPERATION			= "operation"
)

const (
	OPERATION_REBOOT    = "reboot"
	OPERATION_SHUTDOWN	= "shutdown"
	OPERATION_SUSPEND   = "suspend"
)

const (
	CMD_LS              = "dir"
	CMD_REBOOT          = "reboot"
	CMD_SHUTDOWN				= "shutdown"
)

const (
	FLAG_HALT						= "-h"
	FLAG_WHEN       		= "now"
)
