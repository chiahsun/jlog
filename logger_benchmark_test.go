package jlog_test

import (
	"testing"

	"192.168.12.16/Source/IM/Jello/jlog"
)

func BenchmarkInfo(b *testing.B) {
	msg := ""
	for i := 0; i < 100; i++ {
		msg += "jello"
	}

	for n := 0; n < b.N; n++ {
		jlog.Info(msg)
	}
}
