package main

import (
	"os"

	"github.com/rhino-bird/caracal-pty/tool"
	dumpStack "github.com/rhino-bird/caracal-pty/utils/dump_stack"
	"github.com/rhino-bird/caracal-pty/utils/pidfile"
	workerManager "github.com/rhino-bird/caracal-pty/utils/worker_manager"
)

var (
	pidFile = tool.Config.GetString("log.logpath") + tool.ProcessName + ".pid"
)

func main() {
	// Initialize the configuration file
	tool.InitConfig()

	// Parse args
	ch := make(chan CommandArgs)
	go func() {
		args := os.Args
		parseArgs(args, ch)
	}()
	cmd := <-ch

	_ = cmd

	pid, err := pidfile.CheckExit(pidFile)
	if err != nil {
		tool.Log.Error("pidfile raise error, err: %s", err.Error())
		tool.Exit(err, tool.InvalidArgsCode)
	}
	tool.Log.Info("process start sucess, pid: %d", pid)

	go dumpStack.SetupStackTrap()

	wm, err := prepareAllWorker()
	if err != nil {
		tool.Log.Fatalf("prepareAllWorker failed, err: %s", err.Error())
		tool.Exit(err, tool.CannotExecCode)
	}

	wm.Start()
	wm.MakeRecvSignal()
	wm.Stop()

	code := wm.WaitTimeout(tool.ForceExitTime)
	switch code {
	case workerManager.ForceTimeoutExitCode:
		tool.Log.Errorf("timeout force exit, timeout: %v", tool.ForceExitTime)
	case workerManager.GracefulExitCode:
		tool.Log.Info("graceful exit")
	}
}
