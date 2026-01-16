// infrastructure/logging/logger.go
// application/port/logging/logger.go
package logging

type Arg struct {
	Key string
	Val any
}

type Logger interface {
	Info(msg string, args ...Arg)
	Error(msg string, args ...Arg)
	Warn(msg string, args ...Arg)
	Debug(msg string, args ...Arg)
	Trace(msg string, args ...Arg)
}
