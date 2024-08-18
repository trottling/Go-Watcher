package src

func Run() {
	InitLogger()
	Log.Info("Starting Go-Watcher...")
	LogMachineInfo()
	LoadConfig()
	ConnectDB()
	RunProxyServer()
}
