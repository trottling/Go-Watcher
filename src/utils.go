package src

import (
	"fmt"
	"github.com/gookit/slog"
	"github.com/gookit/slog/handler"
	"github.com/gookit/slog/rotatefile"
	"github.com/matishsiao/goInfo"
	"github.com/panjf2000/ants/v2"
	"gopkg.in/elazarl/goproxy.v1"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

func LogMachineInfo() {
	HostInfo, err := goInfo.GetInfo()
	if err != nil {
		Log.Error("Cannot get host info: " + err.Error())
	}

	Log.Info("Running on:")
	Log.Info("\n" + strings.ReplaceAll(strings.ReplaceAll(HostInfo.String(), ",", "\n"), ":", ": "))
}

func GetRunFolder() string {
	filename, err := os.Getwd()
	if err != nil {
		panic("Cannot get running directory: " + err.Error())
	}
	return filename
}

func GetLogsFolder() string {
	logsFolder := filepath.FromSlash(RunFolder + "/" + "work_logs")
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
	return filepath.FromSlash(RunFolder + "\\" + "config.json")
}

func InitLogger() {
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
		WithUseJSON(false).
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

	Log = logger
}

func InitConnectionsPool() {
	Log.Infof("Initializing connections pool : %d CPUs : %d threads", runtime.NumCPU(), Config.ProxyServer.Threads)
	var err error
	ConnectionPool, err = ants.NewMultiPoolWithFunc(runtime.NumCPU(),
		Config.ProxyServer.Threads,
		func(i interface{}) {
			fmt.Println(i)
		},
		ants.RoundRobin,
		ants.WithLogger(Log),
		ants.WithPreAlloc(Config.ProxyServer.PreAllocateMemory))
	if err != nil {
		Log.Panic("Cannot create connections pool for Proxy-Server: " + err.Error())
	}
}

func InitProxyServer() {
	ProxyServer = goproxy.NewProxyHttpServer()
	ProxyServer.Verbose = true
	ProxyServer.
		log.Fatal(http.ListenAndServe(":8080", ProxyServer))
}
