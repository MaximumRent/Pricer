package util

import (
	"os"
	"fmt"
	"encoding/json"
)

func ReadParseConfig(filename string) (ParseConfig) {
	var config ParseConfig
	configFile, err := os.Open(filename)
	defer configFile.Close()
	if err != nil {
		fmt.Println(err.Error())
	}
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)
	return config
}
