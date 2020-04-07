package workermanager

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/rhino-bird/caracal-pty/tool"
)

const (
	// GracefulExitCode exit normally code
	GracefulExitCode = 0

	// ForceTimeoutExitCode timeout exit code
	ForceTimeoutExitCode = 1
)

// Worker working external interface
type Worker interface {
	Start()
	Stop()
	GetProcessName() string
}

// WorkerManager working concrete implementation class
type WorkerManager struct {
	sync.WaitGroup
	WorkerSlice []Worker
	Running     bool
	Q           chan os.Signal
	Ctx         context.Context
	CtxCancel   context.CancelFunc
}

// NewWorkerManager initialization WorkerManager class
func NewWorkerManager() *WorkerManager {
	wm := WorkerManager{
		Running: true,
	}

	wm.WorkerSlice = make([]Worker, 0, 10)
	wm.Q = make(chan os.Signal)
	ctx, cancel := context.WithCancel(context.Background())
	wm.Ctx = ctx
	wm.CtxCancel = cancel
	return &wm
}

// MakeRecvSignal receive user signals
func (wm *WorkerManager) MakeRecvSignal() os.Signal {
	wm.MakeSignal()
	return wm.RecvSignal()
}

// MakeSignal semaphores are currently supported
func (wm *WorkerManager) MakeSignal() {
	signal.Notify(wm.Q,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
		os.Interrupt,
		os.Kill,
	)
}

// RecvSignal receive signal
func (wm *WorkerManager) RecvSignal() os.Signal {
	select {
	case s := <-wm.Q:
		fmt.Println("custom recv signale: ", s)
		return s
	}
}

// Start start all service
func (wm *WorkerManager) Start() {
	for _, worker := range wm.WorkerSlice {
		go func(w Worker) {
			TryCatch(w)
		}(worker)
	}
}

// Stop stop all service
func (wm *WorkerManager) Stop() {
	wm.CtxCancel()
	wm.Running = false

	for _, worker := range wm.WorkerSlice {
		go func(w Worker) {
			defer func() {
				err := recover()
				if err != nil {
					tool.Log.Errorf("WorkerManager error, error:%v, stack: %v\n",
						err, string(tool.Stack()))
				}
			}()

			w.Stop()
		}(worker)
	}
}

// AddWorkerList add workers and start service
func (wm *WorkerManager) AddWorkerList(w []Worker) {
	wm.WorkerSlice = append(wm.WorkerSlice, w...)
}

// AddWorker add worker and start service
func (wm *WorkerManager) AddWorker(w Worker) {
	wm.WorkerSlice = append(wm.WorkerSlice, w)
}

// WaitTimeout close channel and exit
func (wm *WorkerManager) WaitTimeout(timeout int) int {
	endQ := make(chan bool, 0)
	go func() {
		defer close(endQ)
		wm.Wait()
	}()

	select {
	case <-endQ:
		return GracefulExitCode
	case <-time.After(time.Duration(timeout) * time.Second):
		return ForceTimeoutExitCode
	}
}

// TryCatch start service
func TryCatch(f Worker) {
	running := true
	for running {
		func(w Worker) {
			defer func() {
				if e := recover(); e != nil {
					tool.Log.Errorf("worker_name: %s trycatch panicing %v ===> stask: %v \n",
						w.GetProcessName(),
						e,
						string(tool.Stack()),
					)
					time.Sleep(3 * time.Second)
				}
			}()

			w.Start()
			running = false
			return
		}(f)
	}
}
