package main

import (
	"192.168.12.16/Source/IM/Jello/jlog"
	"errors"
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
}

// put to 192.168.12.16/Source/IM/Jello/jlog
// Use submodule

// Both glog and logrus write to file
func main() {
	jlog.Init("logrus.log")
	defer jlog.Flush()

	f1()
}