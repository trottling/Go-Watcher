package src

import (
	"github.com/gookit/slog"
	"github.com/gookit/slog/handler"
	"github.com/gookit/slog/rotatefile"
	"os"
)

func GetRunFolder() string {
	filename, err := os.Getwd()
	if err != nil {
		panic("Cannot get running directory: " + err.Error())
	}
	return filename
}

func GetLogsFolder() string {
	logsFolder := RunFolder + "\\" + "logs"
	// Check for exist and create if nor exist
	if _, err := os.Stat(logsFolder); os.IsNotExist(err) {
		err := os.Mkdir(logsFolder, 0755) // Create the directory with permissions 0755
		if err != nil {
			panic("Cannot create non existing directory for logs: " + err.Error())
		}
	}
	return logsFolder
}

func GetConfigPath() string {
	return RunFolder + "\\" + "config.json"
}

func GetLogger() *slog.Logger {
	// JSON File Handler
	jsonFileHandler := handler.NewBuilder().
		WithLogfile(JsonLogFilePath).
		WithLogLevels(slog.AllLevels).
		WithBuffSize(1024 * 8).
		WithBuffMode(handler.BuffModeBite).
		WithRotateTime(rotatefile.EveryHour).
		WithUseJSON(true).
		Build()

	// Plain text File Handler
	plainFileHandler := handler.NewBuilder().
		WithLogfile(PlainLogFilePath).
		WithLogLevels(slog.AllLevels).
		WithBuffSize(1024 * 8).
		WithBuffMode(handler.BuffModeBite).
		WithRotateTime(rotatefile.EveryHour).
		Build()

	// Console formatter
	consoleFormatter := slog.NewTextFormatter()
	consoleFormatter.EnableColor = true
	// Console handler
	consoleHandler := handler.NewConsoleHandler(slog.AllLevels)
	consoleHandler.SetFormatter(consoleFormatter)

	// Create logger
	logger := slog.New()

	// Add the handlers to the logger
	logger.AddHandler(jsonFileHandler)
	logger.AddHandler(plainFileHandler)
	logger.AddHandler(consoleHandler)

	return logger
}
