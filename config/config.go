package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type Settings struct {
	ApiHost    string `json:"api_host"`
	ApiPort    string `json:"api_port"`
	MongoDbUrl string `json:"mongodb_url"`
}

func GetSettings() Settings {
	configFile, err := os.Open("config.json")
	if err != nil {
		panic(err)
	}
	defer configFile.Close()
	bytes, _ := ioutil.ReadAll(configFile)
	var settings Settings
	err = json.Unmarshal(bytes, &settings)
	if err != nil {
		panic(err)
	}
	return settings
}
