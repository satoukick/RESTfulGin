package logs

import "go.uber.org/zap"

var Logger *zap.SugaredLogger

func init() {
	l, _ := zap.NewDevelopment()
	Logger = l.Sugar()
}
