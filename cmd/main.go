package main

import (
	"os"
)

func main() {
	args := os.Args
	parseArgs(args)
}

/*
func parseArgs(args []string) {
	app := cli.NewApp()
	app.Name = AppName
	app.Version = Version + "+" + CommitID
	app.HideHelp = true
	// app.AppHelpTemplate = tool.helpTemplate
	app.Action = func(c *cli.Context) error {
		if c.Args().Len() == 0 {
			cli.ShowAppHelp(c)
			os.Exit(1)
		}
		return nil
	}

	app.Run(args)
}*/
