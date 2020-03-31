package main

import (
	"os"
)

func main() {
	ch := make(chan CommandArgs)
	args := os.Args
	parseArgs(args, ch)

	cmd := <-ch

	_ = cmd
}
