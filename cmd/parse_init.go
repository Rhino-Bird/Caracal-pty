package main

import (
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
	app.Name = tool.ProcessName
	app.Usage = tool.Usage
	app.Version = Version + "+" + CommitID
	app.HideHelp = true
	app.Authors = getAuthors()
	cli.AppHelpTemplate = helpTemplate

	ops := &tool.Options{}
	if err := tool.ApplyDefaultValues(ops); err != nil {
		tool.Exit(err, tool.ENOEXEC)
	}

	cmdOps := &ptycommand.Options{}
	if err := tool.ApplyDefaultValues(cmdOps); err != nil {
		tool.Exit(err, tool.ENOEXEC)
	}

	flg, fMap := tool.GenerateFlags(ops, cmdOps)
	app.Flags = append(
		flg,
		&cli.StringFlag{
			Name:    tool.ConfName,
			Value:   tool.ConfPath,
			Usage:   tool.ConfUsage,
			EnvVars: tool.ConfEnvVars,
		},
	)

	app.Action = func(c *cli.Context) error {
		if c.NArg() == 0 {
			err := errors.Errorf(globals.NoCommand)
			cli.ShowAppHelp(c)
			tool.Exit(err, tool.EINVAL)
		}

		conf := c.String(tool.ConfName)
		if _, err := os.Stat(conf); !os.IsNotExist(err) {
			if err := tool.ApplyConfigFile(conf, ops, cmdOps); err != nil {
				tool.Exit(err, tool.EINVAL)
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
		tool.Exit(err, tool.ENOEXEC)
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
