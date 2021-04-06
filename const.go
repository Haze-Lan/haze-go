package main

// Command is a signal used to control a running nats-server process.
type Command string

// Valid Command values.
const (
	CommandStop   = Command("stop")
	CommandQuit   = Command("quit")
	CommandReopen = Command("reopen")
	CommandReload = Command("reload")

	// private for now
	commandLDMode = Command("ldm")
	commandTerm   = Command("term")
)

var (
	gitCommit   string
	trustedKeys string
)

const (
	SERVICE_VERSION = "1.0.0"
	SERVICE_NAME    = "haze-framework"
	SERVICE_TAG     = SERVICE_NAME + "-" + SERVICE_VERSION
	PROTO           = 1
	DEFAULT_PORT    = 4222
	RANDOM_PORT     = -1
	DEFAULT_HOST    = "0.0.0.0"
)
