package src

import (
	"database/sql"
	"github.com/gookit/slog"
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
	ProxyServer  *goproxy.ProxyHttpServer
	NetDumpsPath = filepath.FromSlash(RunFolder + "/" + "network_dumps")
)

// AppConfig App config from config.json
type AppConfig struct {
	ProxyServer struct {
		Port                  int  `json:"Port"`
		Threads               int  `json:"Threads"`
		PreAllocateMemory     bool `json:"Pre_Allocate_Memory"`
		ConnectionTimeout     int  `json:"Connection_Timeout"`
		ShowConnectionsSTDOUT bool `json:"Show_Connections_STDOUT"`
	} `json:"Proxy_Server"`
	ActivityHandler struct {
		NonLegitPortsRPM        int      `json:"Non_legit_Ports_RPM"`
		LegitPortsRPM           int      `json:"Legit_Ports_RPM"`
		LegitPathsIgnoreRegex   []string `json:"Legit_Paths_Ignore_Regex"`
		LegitPorts              []int    `json:"Legit_Ports"`
		BlockIPs                bool     `json:"Block_IPs"`
		BlockIPsTime            int      `json:"Block_IPs_time"`
		DumpRequests            bool     `json:"Dump_Requests"`
		RequestsDumpIgnoreRegex []string `json:"Requests_Dump_Ignore_Regex"`
	} `json:"Activity_Handler"`
}

// Connection request and response summary data, will be added to database
type Connection struct {
	IPAddress   string            // IP address of client
	Type        string            // Get, Post, etc.
	Port        int               // Host port
	Path        string            // Host path e.g. /index.html
	Location    string            // Host location e.g. example.com or 127.0.0.1
	StatusCode  int               // Host response status code
	Timestamp   int64             // Connection timestamp in unix format
	ReqHeaders  map[string]string // Request headers
	ReqBody     string            // Request body
	RespHeaders map[string]string // Response headers
	RespBody    string            // Response body
	Allowed     bool              // Is connection allowed
	DumpPath    string            // Connection dump file path
}
