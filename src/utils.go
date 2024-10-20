package src

import (
	"github.com/gookit/slog"
	"github.com/gookit/slog/handler"
	"github.com/gookit/slog/rotatefile"
	"github.com/matishsiao/goInfo"
	"io"
	"net/http"
	"os"
	"path/filepath"
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

func GetConfigPath() string {
	return filepath.FromSlash(RunFolder + "\\" + "config.json")
}

func GetRunFolder() string {
	filename, err := os.Getwd()
	if err != nil {
		panic("Cannot get running directory: " + err.Error())
	}
	return filename
}

func GetLogsFolder() (path string) {
	path = filepath.FromSlash(RunFolder + "/" + "work_logs")
	// Check for exist and create if nor exist
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.Mkdir(path, 0755) // Create the directory with permissions 0755
		if err != nil {
			panic("Cannot create non existing directory for logs: " + err.Error())
		}
	}
	return path
}

func GetNetDumpsPath() (path string) {
	path = filepath.FromSlash(RunFolder + "/" + "network_dumps")
	// Check for exist and create if nor exist
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.Mkdir(path, 0755) // Create the directory with permissions 0755
		if err != nil {
			panic("Cannot create non existing directory for logs: " + err.Error())
		}
	}
	return path
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

func GetReqBody(r *http.Request) (reqBody string) {
	if r.Body != nil {
		// Read body to string
		bodyBytes, err := io.ReadAll(r.Body)
		if err != nil {
			Log.Errorf("Error reading request body: %s", err)
			reqBody = "Error reading request body"
		} else {
			reqBody = string(bodyBytes)
		}
	} else {
		// Set body to "No request body" if it's empty
		reqBody = "No request body"
	}
	return reqBody
}

func GetReqHeaders(r *http.Request) map[string]string {
	headers := r.Header

	if headers == nil {
		Log.Warn("No request headers")
		return map[string]string{}
	}

	headersMap := make(map[string]string)
	for key, values := range r.Header {
		headersMap[key] = strings.Join(values, ", ")
	}
	return headersMap
}

func GetRespBody(r *http.Response) (respBody string) {
	if r.Body != nil {
		// Read body to string
		bodyBytes, err := io.ReadAll(r.Body)
		if err != nil {
			Log.Errorf("Error reading request body: %s", err)
			respBody = "Error reading request body"
		} else {
			respBody = string(bodyBytes)
		}
	} else {
		// Set body to "No request body" if it's empty
		respBody = "No request body"
	}
	return respBody
}

func GetRespHeaders(r *http.Response) map[string]string {
	headers := r.Header

	if headers == nil {
		Log.Warn("No request headers")
		return map[string]string{}
	}

	headersMap := make(map[string]string)
	for key, values := range r.Header {
		headersMap[key] = strings.Join(values, ", ")
	}
	return headersMap
}

func CheckListContainsInt(l []int, r int) bool {
	for _, i := range l {
		if i == r {
			return true
		}
	}
	return false
}
