package Config_Reader

import (
	"Go-Watcher/src"
	"bufio"
	"fmt"
	"log"
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
		src.Log.Fatal("Config file doesn't exist")
	}
	file, err := os.Open("file.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = file.Close(); err != nil {
			src.Log.Fatal(err)
		}
	}()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func ReadConfig() {

}
