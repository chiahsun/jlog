package main

import (
	"192.168.12.16/Source/IM/Jello/jlog"
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
	jlog.Trace("trace")
	jlog.Debug("debug")
	jlog.Info("info")
	jlog.Warning("warning")
	jlog.Error("error")
}

// Put to 192.168.12.16/Source/IM/Jello/jlog
// Use submodule

func main() {
	jlog.Init(jlog.NewLogConfig().SetLogFileOutput("logrus.log", "log"))

	f1()
}
