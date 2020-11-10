package main

import (
	"github.com/sirupsen/logrus"
)

// NewLogger - Создаёт новый логгер
func NewLogger() (logger *logrus.Logger, err error) {
	lg := logrus.New()

	lg.ReportCaller = false
	lg.Level = logrus.DebugLevel
	lg.Formatter = &logrus.TextFormatter{}

	return lg, nil
}
