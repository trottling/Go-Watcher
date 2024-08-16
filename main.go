package main

import (
	"Go-Watcher/src"
	"Go-Watcher/src/Config-Reader"
)

func main() {
	src.Log.Info("Starting Go-Watcher...")
	Config_Reader.LoadConfig()
}
