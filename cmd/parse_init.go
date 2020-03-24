package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

func parseArgs(args []string) {

	fmt.Println(Version)
	app := cli.NewApp()
	app.Name = AppName
	app.Usage = Usage
	app.Version = Version + "+" + CommitID
	app.HideHelp = true
	app.Authors = getAuthors()
	cli.AppHelpTemplate = helpTemplate

	// app.CustomAppHelpTemplate = helpTemplate
	app.Action = func(c *cli.Context) error {
		if c.Args().Len() == 0 {
			cli.ShowAppHelp(c)
			os.Exit(1)
		}
		return nil
	}

	app.Run(args)
}

func getAuthors() []*cli.Author {
	auts := make([]*cli.Author, 0, 2)
	for k, v := range Authors {
		var aut cli.Author
		aut.Name = k
		aut.Email = v

		auts = append(auts, &aut)
	}

	return auts
}

func exit(err error, code int) {
	if err != nil {
		fmt.Println(err)
	}
	os.Exit(code)
}
