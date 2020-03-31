package main

import (
	"fmt"
	"os"

	"github.com/pkg/errors"
	"github.com/rhino-bird/caracal-pty/globals"
	ptycommand "github.com/rhino-bird/caracal-pty/server/pty_command"
	"github.com/rhino-bird/caracal-pty/tool"
	"github.com/urfave/cli/v2"
)

// CommandArgs command options
type CommandArgs struct {
	Ops    *tool.Options
	CmdOps *ptycommand.Options
}

func parseArgs(args []string, ch chan<- CommandArgs) {
	app := cli.NewApp()
	app.Name = AppName
	app.Usage = Usage
	app.Version = Version + "+" + CommitID
	app.HideHelp = true
	app.Authors = getAuthors()
	cli.AppHelpTemplate = helpTemplate

	ops := &tool.Options{}
	if err := tool.ApplyDefaultValues(ops); err != nil {
		exit(err, globals.ENOEXEC)
	}

	cmdOps := &ptycommand.Options{}
	if err := tool.ApplyDefaultValues(cmdOps); err != nil {
		exit(err, globals.ENOEXEC)
	}

	flg, fMap := tool.GenerateFlags(ops, cmdOps)
	app.Flags = append(
		flg,
		&cli.StringFlag{
			Name:    globals.ConfName,
			Value:   globals.ConfPath,
			Usage:   globals.ConfUsage,
			EnvVars: globals.ConfEnvVars,
		},
	)

	app.Action = func(c *cli.Context) error {
		if c.NArg() == 0 {
			err := errors.Errorf(globals.NoCommand)
			cli.ShowAppHelp(c)
			exit(err, globals.EINVAL)
		}

		conf := c.String(globals.ConfName)
		if _, err := os.Stat(conf); !os.IsNotExist(err) {
			if err := tool.ApplyConfigFile(conf, ops, cmdOps); err != nil {
				exit(err, globals.EINVAL)
			}
		}

		tool.ApplyFlags(flg, fMap, c, ops, cmdOps)

		host, _ := os.Hostname()
		ops.TitleVariable = map[string]interface{}{
			"command":  args[0],
			"argv":     args[1:],
			"hostname": host,
		}

		ch <- CommandArgs{
			Ops:    ops,
			CmdOps: cmdOps,
		}
		return nil
	}

	err := app.Run(args)
	if err != nil {
		exit(err, globals.ENOEXEC)
	}
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
