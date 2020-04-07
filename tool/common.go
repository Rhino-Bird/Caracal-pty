package tool

import (
	"fmt"
	"os"
	"runtime"
)

// Configure file
var (
	// ConfPath Configuration file path
	ConfPath string = "./conf/hcl/conf.hcl"

	// ConfName Configuration file name
	ConfName string = "config"

	// ConfUsage Configuration file usage
	ConfUsage string = "Config file path"

	// ConfEnvVars Configuration Environment variable
	ConfEnvVars []string = []string{"CARACAL_PTY_CONFIG"}

	// ConfDir configuration path
	ConfDir string = "./conf/yaml/"

	stackBuf int = 102428
)

// Exit exit process
func Exit(err error, code int) {
	if err != nil {
		fmt.Println(err)
	}
	os.Exit(code)
}

// Stack pid stack buf
func Stack() []byte {
	buf := make([]byte, stackBuf)
	n := runtime.Stack(buf, false)
	return buf[:n]
}
