package main

var (
	// Version build flag
	Version string

	// Branch build branch name
	Branch string

	// CommitID build flag
	CommitID string

	// BuildTime build flag
	BuildTime string

	// Authors the authors information.
	Authors = map[string]string{
		"xuesongbj": "davidbjhd@gmail.com",
		"Youxun-Zh": "youxun.zh@gmail.com",
	}
)

const (
	// AppName application name.
	AppName = "caracal-pty"

	// Usage usage infomation.
	Usage = "Caracal pty."
)
