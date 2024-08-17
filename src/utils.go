package src

import (
	"fmt"
	"github.com/gookit/slog"
	"github.com/gookit/slog/handler"
	"github.com/gookit/slog/rotatefile"
	"github.com/matishsiao/goInfo"
	"github.com/panjf2000/ants/v2"
	"os"
	"runtime"
	"strings"
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
		WithRotateTime(rotatefile.EveryHour).
		WithUseJSON(true).
		Build()

	// Plain text File Handler
	plainFileHandler := handler.NewBuilder().
		WithLogfile(PlainLogFilePath).
		WithLogLevels(slog.AllLevels).
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

func LogMachineInfo() {
	HostInfo, err := goInfo.GetInfo()
	if err != nil {
		Log.Error("Cannot get host info: " + err.Error())
	}

	Log.Info("Running on:")
	Log.Info("\n" + strings.ReplaceAll(HostInfo.String(), ",", "\n"))
}

func GetConnectionsPool() *ants.MultiPoolWithFunc {
	pool, err := ants.NewMultiPoolWithFunc(runtime.NumCPU(),
		Config.ProxyServer.Threads,
		func(i interface{}) {
			fmt.Println(i)
		},
		ants.RoundRobin,
		ants.WithLogger(Log),
		ants.WithPreAlloc(Config.ProxyServer.PreAllocateMemory))
	if err != nil {
		Log.Panic("Cannot create connections pool for proxy-server: " + err.Error())
	}
	return pool
}
