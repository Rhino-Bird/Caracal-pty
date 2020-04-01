package main

import (
	"fmt"
	"os"

	"github.com/rhino-bird/caracal-pty/tool"
)

func main() {
	// Initialize the configuration file
	tool.InitConfig()

	// Parse args
	ch := make(chan CommandArgs)
	go func() {
		args := os.Args
		parseArgs(args, ch)
	}()
	cmd := <-ch

	fmt.Println(*cmd.Ops, *cmd.CmdOps)
}
