package main

import (
	"os"

	"github.com/rhino-bird/caracal-pty/globals"
	ptycommand "github.com/rhino-bird/caracal-pty/server/pty_command"
	"github.com/rhino-bird/caracal-pty/tool"
	"github.com/rhino-bird/caracal-pty/utils"
)

func main() {
	args := os.Args
	parseArgs(args)

	ops := &tool.Options{}
	if err := utils.ApplyDefaultValues(ops); err != nil {
		exit(err, globals.ENOEXEC)
	}

	cmdOps := &ptycommand.Options{}
	if err := utils.ApplyDefaultValues(cmdOps); err != nil {
		exit(err, globals.ENOEXEC)
	}
}
