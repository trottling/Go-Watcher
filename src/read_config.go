package src

import (
	"encoding/json"
	"os"
	"regexp"
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

	// Compile all regex from config
	CompileLegitPathsIgnoreRegex()
	CompileRequestsDumpIgnoreRegex()
}

func CompileLegitPathsIgnoreRegex() (LegitPathsIgnoreRegexList []*regexp.Regexp) {
	for _, path := range Config.ActivityHandler.LegitPathsIgnoreRegex {
		reg, err := regexp.Compile(path)
		if err != nil {
			Log.Error("Cannot compile regex: " + err.Error())
		}
		LegitPathsIgnoreRegexList = append(LegitPathsIgnoreRegexList, reg)
	}
	return LegitPathsIgnoreRegexList
}

func CompileRequestsDumpIgnoreRegex() (RequestsDumpIgnoreRegexList []*regexp.Regexp) {
	for _, path := range Config.ActivityHandler.RequestsDumpIgnoreRegex {
		reg, err := regexp.Compile(path)
		if err != nil {
			Log.Error("Cannot compile regex: " + err.Error())
		}
		RequestsDumpIgnoreRegexList = append(RequestsDumpIgnoreRegexList, reg)
	}
	return LegitPathsIgnoreRegexList
}
