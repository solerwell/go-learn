package logger

import (
	"github.com/sirupsen/logrus"
	"os"
)

var stdLog = logrus.New()

func init() {
	stdLog.SetOutput(os.Stdout)
	stdLog.SetFormatter(&logrus.JSONFormatter{})
	stdLog.SetLevel(logrus.DebugLevel)
}

// 获取标准输出logger
func GetStdLogger() *logrus.Logger {
	return stdLog
}
