package main

var helpTemplate = `NAME:
	{{.Name}} = {{.Usage}}

USAGE:
	{{.Name}} [options] <command> [<arguments...>]

VERSION:
	{{.Version}}

OPTIONS:
	{{range .Flags}}{{.}}
	{{end}}
`
