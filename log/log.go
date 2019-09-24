package logs

import "go.uber.org/zap"

var logger *zap.SugaredLogger

func init() {
	l, _ := zap.NewDevelopment()
	logger = l.Sugar()
}

// Info provides info logging
func Info(v ...interface{}) {
	logger.Info(v...)
}

// Debug provides debug logging
func Debug(v ...interface{}) {
	logger.Debug(v...)
}

// Fatal provides fatal logging
func Fatal(v ...interface{}) {
	logger.Fatal(v...)
}

// Error provides error logging
func Error(v ...interface{}) {
	logger.Error(v...)
}

// Sync flushes buffered entries.
// This should be called before exiting application.
func Sync() {
	logger.Sync()
}
