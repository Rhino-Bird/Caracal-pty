package main

import (
	"os"

	"github.com/urfave/cli/v2"
)

func parseArgs(args []string) {
	app := cli.NewApp()
	app.Name = AppName
	app.Usage = Usage
	app.Version = Version + "+" + CommitID
	app.HideHelp = true

	app.CustomAppHelpTemplate = helpTemplate
	app.Action = func(c *cli.Context) error {
		if c.Args().Len() == 0 {
			cli.ShowAppHelp(c)
			os.Exit(1)
		}
		return nil
	}

	app.Run(args)
}
