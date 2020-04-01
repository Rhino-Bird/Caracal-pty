package tool

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"time"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	rotatelogs "github.com/xuesongbj/go-file-rotatelogs"
	"github.com/xuesongbj/lfshook"
)

// Log export outside
var Log = &logrus.Entry{
	Logger: logger,
}

// Inside
var logger = logrus.New()

// InitLog initialize the log
func InitLog() {
	// Automatic log segmentation
	InitLoggerLogrotate(
		Config.GetString("log.logpath"),
		Config.GetString("log.logfile"),
		Config.GetString("log.level"),
		Config.GetString("log.format"),
		Config.GetDuration("log.loglifetime")*time.Hour,
		Config.GetDuration("log.logrotation")*time.Hour,
	)

	// ignore terminal stdout
	logger.SetOutput(ioutil.Discard)
}

// InitLoggerLogrotate initialize the log file system
func InitLoggerLogrotate(logPath, logFileName, level, format string, maxAge time.Duration, rotationTime time.Duration) {
	baseLogPath := path.Join(logPath, logFileName)
	writer, err := rotatelogs.New(
		baseLogPath+".%Y%m%d",
		rotatelogs.WithLinkName(logPath+logFileName), // Generate soft link, point to the latest log file
		rotatelogs.WithMaxAge(maxAge),                // Maximum file save time
		rotatelogs.WithRotationTime(rotationTime),    // Log cutting interval
	)

	if err != nil {
		Log.Errorf("config local file system logger error. %+v", errors.WithStack(err))
		Exit(err, CannotExecCode)
	}

	switch level {
	case "debug":
		logger.SetLevel(logrus.DebugLevel)
	case "info":
		logger.SetLevel(logrus.InfoLevel)
	case "warn":
		logger.SetLevel(logrus.WarnLevel)
	case "error":
		logger.SetLevel(logrus.ErrorLevel)
	default:
		logger.SetLevel(logrus.InfoLevel)
	}

	lfHook := lfshook.NewHook(lfshook.WriterMap{
		logrus.DebugLevel: writer,
		logrus.InfoLevel:  writer,
		logrus.WarnLevel:  writer,
		logrus.ErrorLevel: writer,
		logrus.FatalLevel: writer,
		logrus.PanicLevel: writer,
	}, &logrus.TextFormatter{DisableColors: true, TimestampFormat: "2006-01-02 15:04:05.000"})

	if format == "json" {
		lfHook.SetFormatter(&logrus.JSONFormatter{})
	} else {
		lfHook.SetFormatter(&logrus.TextFormatter{})
	}

	logger.AddHook(lfHook)
}

// NewLogger need linux logrotate process split log
func NewLogger(logPath, fileName, level, typeof string) *logrus.Logger {
	file := logPath + fileName
	cLog := logrus.New()
	fileFd, err := os.OpenFile(file, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)

	if err != nil {
		fmt.Printf("config local file system logger error. %+v \n", errors.WithStack(err))
		Exit(err, CannotExecCode)
	}
	cLog.Out = fileFd

	switch typeof {
	case "text":
		// Text
		cLog.Formatter = &logrus.TextFormatter{}
	default:
		// Json
		cLog.Formatter = &logrus.JSONFormatter{}
	}

	levelFlag, err := logrus.ParseLevel(level)
	if err != nil {
		fmt.Printf("parse log level faild, error: %+v \n", errors.WithStack(err))
		levelFlag = logrus.InfoLevel
	}
	cLog.Level = levelFlag
	return cLog
}
