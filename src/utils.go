package src

import (
	"github.com/gookit/slog"
	"github.com/gookit/slog/handler"
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

	// File handler for JSON
	jsonFileHandler, err := handler.JSONFileHandler(JsonLogFilePath)
	if err != nil {
		panic("Error creating JSON file handler: " + err.Error())
	}
	// File handler for Plain text
	plainFileHandler, err := handler.NewFileHandler(PlainLogFilePath)
	if err != nil {
		panic("Error creating plain text file handler: " + err.Error())
	}

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
