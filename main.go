package main

import (
	"Go-Watcher/src"
	"Go-Watcher/src/Config-Reader"
)

func main() {
	src.Log.Info("Starting Go-Watcher...")
	src.LogMachineInfo()
	Config_Reader.LoadConfig()
}
