package Config_Reader

import (
	"Go-Watcher/src"
	"encoding/json"
	"os"
)

func LoadConfig() {
	src.Log.Info("Loading config...")
	ValidateConfig()
	ReadConfig()
}
func ValidateConfig() {
	// Check for config file exist
	_, err := os.Stat(src.ConfigPath)
	if os.IsNotExist(err) {
		src.Log.Panic("Config file doesn't exist: " + src.ConfigPath)
	}
}

func ReadConfig() {
	src.Log.Info("Config path: " + src.ConfigPath)
	file, err := os.Open(src.ConfigPath)
	if err != nil {
		src.Log.Fatal("Cannot create config: " + err.Error())
	}
	defer func() {
		if err = file.Close(); err != nil {
			src.Log.Fatal("Cannot create config: " + err.Error())
		}
	}()

	configFile, err := os.Open(src.ConfigPath)
	if err != nil {
		src.Log.Panic("Cannot read config: " + err.Error())
	}

	// Marshal the JSON data
	jsonParser := json.NewDecoder(configFile)
	err = jsonParser.Decode(&src.Config)
	if err != nil {
		src.Log.Panic("Cannot read config: " + err.Error())
	}

	// Print the config for debugging
	jsonConfig, _ := json.MarshalIndent(src.Config, "", "  ")
	src.Log.Info("Config JSON:", string(jsonConfig))
	src.Log.Info("Config loaded successfully")
}
