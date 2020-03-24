package globals

// Process manager
const (
	// ProcessName process name
	ProcessName string = "CARACAL-PTY"
)

// Configure file
var (
	// ConfPath Configuration file path
	ConfPath string = "./conf/conf.yaml"

	// ConfName Configuration file name
	ConfName string = "config"

	// ConfUsage Configuration file usage
	ConfUsage string = "Config file path"

	// ConfEnvVars Configuration Environment variable
	ConfEnvVars []string = []string{"CARACAL_PTY_CONFIG"}
)
