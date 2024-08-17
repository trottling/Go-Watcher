package src

import "Go-Watcher/src/Config-Reader"

func Run() {
	InitLogger()
	Log.Info("Starting Go-Watcher...")
	LogMachineInfo()
	Config_Reader.LoadConfig()
	ConnectDB()
	InitConnectionsPool()
}
