package jlog_test

import (
	"testing"

	"github.com/chiahsun/jlog"
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
