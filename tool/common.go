package tool

import (
	"fmt"
	"os"
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
)

// Exit exit process
func Exit(err error, code int) {
	if err != nil {
		fmt.Println(err)
	}
	os.Exit(code)
}
