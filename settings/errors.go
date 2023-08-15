package settings

import "fmt"

var (
	ErrEnvVarNotFound = fmt.Errorf("settings.go: env var not found")
	ErrParsingPort    = fmt.Errorf("settings.go: the port number must be an positive integer greater than 0")
)
