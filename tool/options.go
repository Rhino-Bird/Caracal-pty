package tool

// Options command line options
type Options struct {
	Address       string `hcl:"address" flagName:"address" flagSName:"a" flagDescribe:"IP address to listen" default:"0.0.0.0"`
	Port          string `hcl:"port" flagName:"port" flagSName:"p" flagDescribe:"Port number to liten" default:"8080"`
	Term          string `hcl:"term" flagName:"term" flagSName:"t" flagDescribe:"Terminal name to use on the browser, one of xterm or hterm." default:"xterm"`
	TitleFormat   string `hcl:"title_format" flagName:"title-format" flagSName:"" flagDescribe:"Title format of browser window" default:"{{ .command }}@{{ .hostname }}"`
	Maxconnection int    `hcl:"max_connection" flagName:"max-connection" flagDescribe:"Maximum connection to gotty" default:"0"`
	Timeout       int    `hcl:"timeout" flagName:"timeout" flagSName:"" flagDescribe:"Timeout seconds for waiting a client(0 to disable)" default:"0"`
	Width         int    `hcl:"width" flagName:"width" flagSName:"" flagDescribe:"Static width of the screen, 0(default) means dynamically resize" default:"0"`
	Height        int    `hcl:"height" flagName:"height" flagSName:"" flagDescribe:"Static height of the screen, 0(default) means dynamically resize" default:"0"`
	PermitWrite   bool   `hcl:"permit_write" flagName:"permit-write" flagSName:"w" flagDescribe:"Permit clients to write to the TTY (BE CAREFUL)" default:"false"`

	TitleVariable map[string]interface{}
}

type CommandArgs struct {
	Ops Options
}
