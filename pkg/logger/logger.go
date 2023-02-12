// Package logger provides tools for working with custom logger implementations.
package logger

// Logger is an interface that contains methods for working with a specific logger implementation.
type Logger interface {
	Infof(template string, args ...interface{})
	Errorf(template string, args ...interface{})
	Fatalf(template string, args ...interface{})
}
