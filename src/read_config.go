package src

import (
	"encoding/json"
	"os"
)

func LoadConfig() {
	Log.Info("Loading config...")

	// Check for config file exist
	_, err := os.Stat(ConfigPath)
	if os.IsNotExist(err) {
		Log.Panic("Config file doesn't exist: " + ConfigPath)
	}

	Log.Info("Config path: " + ConfigPath)
	configFile, err := os.Open(ConfigPath)
	if err != nil {
		Log.Panic("Cannot read config: " + err.Error())
	}
	defer func() {
		if err = configFile.Close(); err != nil {
			Log.Panic("Cannot read config: " + err.Error())
		}
	}()

	// Marshal the JSON data
	jsonParser := json.NewDecoder(configFile)
	err = jsonParser.Decode(&Config)
	if err != nil {
		Log.Panic("Cannot read config: " + err.Error())
	}

	// Print the config for debugging
	jsonConfig, _ := json.MarshalIndent(Config, "", "  ")
	Log.Info("Config JSON:", "\n", string(jsonConfig))
	Log.Info("Config loaded successfully")
}
