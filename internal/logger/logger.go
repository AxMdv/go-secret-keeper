package logger

import (
	"log"
	"log/slog"
	"os"
	"strings"
)

var LOG *slog.Logger

func Init(logLevel string) error {
	var lvl slog.Level
	switch strings.ToUpper(logLevel) {
	case "ERROR":
		lvl = slog.LevelError
	case "INFO":
		lvl = slog.LevelInfo
	case "DEBUG":
		lvl = slog.LevelDebug
	default:
		log.Fatal("Logging level can be only ERROR, INFO, DEBUG")
	}

	file, err := os.OpenFile("tmp/log/app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("Failed to open log file:", err)
		return err
	}
	handler := slog.NewTextHandler(file, &slog.HandlerOptions{
		AddSource: false,
		Level:     lvl,
	})
	LOG = slog.New(handler)
	// jsonHandler := slog.NewJSONHandler(file, &slog.HandlerOptions{
	// 	AddSource: true,
	// 	Level:     slog.LevelDebug,
	// })

	// Flog = slog.New(jsonHandler)

	slog.SetDefault(LOG)
	return nil
}
