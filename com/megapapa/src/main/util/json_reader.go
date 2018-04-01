package util

import (
	"os"
	"encoding/json"
	"log"
	. "../entity"
)

const _CONFIG_FILE_PATH = "com/megapapa/resources/parse_config.json"

/* Simple json-config-reader */
func ReadParseConfig() (ParseConfig) {
	var config ParseConfig
	configFile, err := os.Open(_CONFIG_FILE_PATH)
	defer configFile.Close()
	if err != nil {
		log.Println(err.Error())
	}
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)
	return config
}
