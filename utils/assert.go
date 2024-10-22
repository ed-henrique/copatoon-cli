package utils

import (
	"fmt"
	"log/slog"
	"os"
)

const (
	logFile = "copatoon.log"
)

var (
	loggerFile, _ = os.OpenFile(logFile, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	logger        = slog.New(
		slog.NewTextHandler(
			loggerFile,
			&slog.HandlerOptions{
				Level: slog.LevelWarn,
			},
		),
	)
)

func Assert(condition bool, scope, msg string, attr ...any) {
	if !condition {
		logger.Error(fmt.Sprintf("[%s] %s", scope, msg), attr...)
		os.Exit(1)
	}
}
