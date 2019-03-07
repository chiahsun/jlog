package main

import (
	"errors"
	"jlog/jlog"
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
	// jlog.Init("logrus.log") // Use this if you want to log to file
	jlog.Init("")
	defer jlog.Flush()

	f1()
}
