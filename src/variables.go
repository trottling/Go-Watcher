package src

import (
	"database/sql"
	"github.com/gookit/slog"
	"github.com/panjf2000/ants/v2"
	"gopkg.in/elazarl/goproxy.v1"
	"path/filepath"
	"time"
)

//goland:noinspection GoCommentStart
var (
	// Misc
	RunFolder = GetRunFolder()

	// Logger
	Log              *slog.Logger
	JsonLogFilePath  = filepath.FromSlash(GetLogsFolder() + "/" + time.Now().Format("2006-01-02_15-04-05") + "_log.json")
	PlainLogFilePath = filepath.FromSlash(GetLogsFolder() + "/" + time.Now().Format("2006-01-02_15-04-05") + "_log.txt")

	// Config
	ConfigPath = GetConfigPath()
	Config     = AppConfig{}

	// Database
	DBPath = filepath.FromSlash(RunFolder + "/" + "database.db")
	DBConn *sql.DB

	// Proxy server
	ProxyServer    *goproxy.ProxyHttpServer
	ConnectionPool *ants.MultiPoolWithFunc
)

// AppConfig App config from config.json
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

type Request struct {
	IPAddress  string // IP address of client
	Protocol   string // Client protocol
	Port       int    // Host port
	Path       string // Host path e.g. /index.html
	Location   string // Host location e.g. example.com or 127.0.0.1
	StatusCode int    // Host response status code
	Timestamp  int64  // Request timestamp
}
