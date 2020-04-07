package dumpstack

import (
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"
)

const (
	timeFormat = "2006-01-02 15:04:05"

	// bufSize 1MB,
	// 102400 << 4 ==> 1,638,400 Bytes
	bufSize = 102400 << 4
)

var (
	stdFile = "./stack.log"
)

// SetupStackTrap dump pid stack
func SetupStackTrap(args ...string) {
	if len(args) > 0 {
		stdFile = args[0]
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGUSR1)
	go func() {
		for range c {
			dumpStacks()
		}
	}()
}

func dumpStacks() {
	buf := make([]byte, bufSize)
	buf = buf[:runtime.Stack(buf, true)]
	writeStack(buf)
}

func writeStack(buf []byte) {
	fd, _ := os.OpenFile(stdFile, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)

	now := time.Now().Format(timeFormat)
	fd.WriteString("\n\n\n\n\n")
	fd.WriteString(now + " stdout:" + "\n\n")
	fd.Write(buf)
	fd.Close()
}