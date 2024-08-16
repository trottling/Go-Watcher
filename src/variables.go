package src

import (
	"time"
)

//goland:noinspection GoCommentStart
var (
	// Misc
	RunFolder = GetRunFolder()

	// Config
	ConfigPath = "config.json"
	Config     = ""

	// Logger | TODO Add logs files rotating
	Log              = GetLogger()
	JsonLogFilePath  = GetLogsFolder() + "/" + time.Now().Format("2006-01-02_15-04-05") + "_json_log.txt"
	PlainLogFilePath = GetLogsFolder() + "/" + time.Now().Format("2006-01-02_15-04-05") + "_plain_log.txt"
)

type AppConfig struct {
}
