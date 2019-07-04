package logs

import "go.uber.org/zap"

var Logger *zap.SugaredLogger

// FIXME: log struct
func init() {
	l, _ := zap.NewDevelopment()
	Logger = l.Sugar()
}
