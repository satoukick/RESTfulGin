package logs

import "go.uber.org/zap"

var Logger *zap.SugaredLogger

// FIXME:
func init() {
	l, _ := zap.NewDevelopment()
	Logger = l.Sugar()
}
