package pidfile

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"syscall"
	"unsafe"
)

var (
	// ErrPid process id failed
	ErrPid = errors.New("process id error")

	// Debug whether to enable debug mode
	Debug bool

	// ErrServerAlreadyRun Service is already running
	ErrServerAlreadyRun = errors.New("server already running")

	// ErrPidExists pid already exists
	ErrPidExists = errors.New("pid file exists")

	// ErrLockFaild pid lock failed
	ErrLockFaild = errors.New("pid file lock faild")
)

// CheckExit check if pid is running
func CheckExit(filename string) (int, error) {
	pid, err := GetPID(filename)
	if err != nil {
		if b, _ := IsActive(pid); b {
			return pid, ErrServerAlreadyRun
		}
	}

	pval, err := Create(filename)
	if err != nil {
		return pval, ErrLockFaild
	}
	return pval, nil
}

// GetPID get pid
func GetPID(pidfile string) (int, error) {
	var pid int

	pval, err := ioutil.ReadFile(pidfile)
	if err != nil {
		if Debug {
			log.Printf("read pid file: %s\n", err.Error())
		}
		return pid, err
	}

	v := (*string)(unsafe.Pointer(&pval))
	tmp, err := strconv.ParseInt(*v, 10, 32)
	pid = int(tmp)
	if err != nil {
		if Debug {
			log.Printf("trans pid to int: %s\n", err.Error())
		}
		return pid, err
	}
	return pid, nil
}

// IsActive check if the pid process is alive
func IsActive(pid int) (bool, error) {
	if pid <= 0 {
		return false, ErrPid
	}

	p, err := os.FindProcess(pid)
	if err != nil {
		if Debug {
			log.Printf("find process: %s\n", err.Error())
		}
		return false, err
	}

	if err := p.Signal(os.Signal(syscall.Signal(0))); err != nil {
		if Debug {
			log.Printf("send signal [0]: %s\n", err.Error())
		}
		return false, err
	}
	return true, nil
}

// Create create pid file
func Create(pidfile string) (int, error) {
	var pid int

	if _, err := os.Stat(pidfile); !os.IsNotExist(err) {
		if ok, _ := IsActive(pid); ok {
			return pid, ErrPidExists
		}
	}

	// Create pidfile
	pf, err := os.OpenFile(pidfile, os.O_RDWR|os.O_CREATE, 0600)
	if err != nil {
		err = fmt.Errorf("error when create pid file: %s", err.Error())
		return pid, err
	}

	// Get pid
	pid = os.Getpid()
	pf.Write([]byte(strconv.Itoa(pid)))
	if err := CheckLockFile(pf.Fd()); err != nil {
		return pid, err
	}
	return pid, nil
}

// CheckLockFile check if the fd is locked
func CheckLockFile(fd uintptr) error {
	err := syscall.Flock(int(fd), syscall.LOCK_EX|syscall.LOCK_NB)
	if err == syscall.EWOULDBLOCK {
		err = ErrLockFaild
	}
	return err
}
