package jlog_test

import (
	"192.168.12.16/Source/IM/Jello/jlog"
	"testing"
)

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

