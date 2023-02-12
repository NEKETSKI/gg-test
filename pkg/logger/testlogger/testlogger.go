//go:build testmode

package testlogger

import "github.com/NEKETSKY/gg-test/pkg/logger"

type fakelogger struct{}

// New returns a new instance of the test logger.
func New() logger.Logger {
	return new(fakelogger)
}

func (*fakelogger) Infof(_ string, _ ...interface{}) {
}

func (*fakelogger) Errorf(_ string, _ ...interface{}) {
}

func (*fakelogger) Fatalf(_ string, _ ...interface{}) {
}
