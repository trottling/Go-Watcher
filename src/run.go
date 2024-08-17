package src

import (
	"Go-Watcher/src/Config-Reader"
	"Go-Watcher/src/Proxy-Server"
)

func Run() {
	InitLogger()
	Log.Info("Starting Go-Watcher...")
	LogMachineInfo()
	Config_Reader.LoadConfig()
	ConnectDB()
	Proxy_Server.InitConnectionsPool()
	Proxy_Server.RunProxyServer()
}
