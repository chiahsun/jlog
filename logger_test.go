package jlog_test

import (
	"sync"
	"testing"

	"192.168.12.16/Source/IM/Jello/jlog"
)

func TestMain(m *testing.M) {
	jlog.Init(jlog.NewLogConfig().SetLogFileOutput("logrus.log", "log"))
	m.Run()
}

func TestJLOG_Info(t *testing.T) {
	msg := "Info Test"
	jlog.Info(msg)
	jlog.Infof(msg)
	jlog.Infoln(msg)
}

func TestJLOG_Warning(t *testing.T) {
	msg := "Warning Test"
	jlog.Warning(msg)
	jlog.Warningf(msg)
	jlog.Warningln(msg)
}

func TestJLOG_Error(t *testing.T) {
	msg := "Error Test"
	jlog.Error(msg)
	jlog.Errorf(msg)
	jlog.Errorln(msg)
}

func TestJLOG_Fatal(t *testing.T) {
	// msg := "Fatal Test"
	// jlog.Fatal(msg)
	// jlog.Fatalf(msg)
	// jlog.Fatalln(msg)
}

func TestJLOG_Concurrency(t *testing.T) {
	msg := ""
	for i := 0; i < 10000; i++ {
		msg += "1"
	}

	wg := sync.WaitGroup{}
	for n := 0; n < 1000; n++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			jlog.Info(msg)
		}()
	}
	wg.Wait()
}
