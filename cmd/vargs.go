package main

import (
	workerManager "github.com/rhino-bird/caracal-pty/utils/worker_manager"
)

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

	// WorkerPool for debug
	WorkerPool = map[string][]workerManager.Worker{}
)
