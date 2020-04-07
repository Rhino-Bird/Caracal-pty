package tool

const (
	// ENOEXEC Exec format error
	ENOEXEC = 8

	// EINVAL Invalid argument
	EINVAL = 22

	// CannotExecCode command invoked cannot execute
	CannotExecCode = 0x7E

	// InvalidArgsCode invalid argument to exit
	InvalidArgsCode = 0x80
)

// Process manager
const (
	// ProcessName process name
	ProcessName string = "caracal-pty"

	// Usage usage infomation.
	Usage = "Caracal pty."

	// ForceExitTime force exit time
	ForceExitTime = 15 // second
)
