package logger

import (
	"github.com/sirupsen/logrus"
)

func Info(msg interface{}) {
	logrus.Infoln(msg)
}

func Warn(msg interface{}) {
	logrus.Warnln(msg)
}

func Error(msg interface{}) {
	logrus.Error(msg)
}

func Panic(msg interface{}) {
	logrus.Panic(msg)
}
