package main

import (
	"errors"

	"github.com/chiahsun/jlog"
)

func f1() {
	f2()
}

func f2() {
	f3()
}

func f3() {
	f4()
}

func f4() {
	jlog.Info("this is my message")
	jlog.Info(errors.New("my error"))
	jlog.Error("error message")
}

func main() {
	const (
		logLevel = jlog.DebugLevel
	)
	jlog.Init(jlog.NewLogConfig().SetLogStdout(true))
	// jlog.Init(jlog.NewLogConfig().SetLogFileOutput("log", "logrus.log").SetLogLevel(logLevel))  // Use this if you want to log to file

	f1()
}
