package src

import (
	"time"
)

//goland:noinspection GoCommentStart
var (
	// Misc
	RunFolder = GetRunFolder()

	// Config
	ConfigPath = GetConfigPath()
	Config     = AppConfig{}

	// Logger
	Log              = GetLogger()
	JsonLogFilePath  = GetLogsFolder() + "\\" + time.Now().Format("2006-01-02_15-04-05") + "_log.json"
	PlainLogFilePath = GetLogsFolder() + "\\" + time.Now().Format("2006-01-02_15-04-05") + "_log.txt"
)

type AppConfig struct {
	ProxyServer struct {
		Address           string `json:"Address"`
		Port              int    `json:"Port"`
		Threads           int    `json:"Threads"`
		PreAllocateMemory bool   `json:"Pre_Allocate_Memory"`
		ConnectionTimeout int    `json:"Connection_Timeout"`
	} `json:"Proxy_Server"`
	ActivityHandler struct {
		NonLegitPortsRPM int   `json:"Non_legit_Ports_RPM"`
		LegitPortsRPM    int   `json:"Legit_Ports_RPM"`
		LegitPorts       []int `json:"Legit_Ports"`
		BlockIPs         bool  `json:"Block_IPs"`
		BlockIPsTime     int   `json:"Block_IPs_time"`
	} `json:"Activity_Handler"`
}
