package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

const CONFIG_FILE = "config\\portConf.json"

type PortFile struct {
	Table    string `json:"table"`
	Schema   string `json:"schema"`
	Data     string `json:"data"`
	Selected bool   `json:"selected"`
}

type Configuration struct {
	Tables []PortFile `json:""`
}

var PortConfiguration Configuration = Configuration{}

func init() {
	fileConf, err := os.Open(CONFIG_FILE)
	if err != nil {
		log.Fatalf("File open\n%v\n", err)
	}
	defer fileConf.Close()

	dataConf, err := ioutil.ReadAll(fileConf)
	if err != nil {
		log.Fatalf("File read\n%v\n", err)
	}

	err = json.Unmarshal(dataConf, &PortConfiguration)
	if err != nil {
		log.Fatalf("File load\n%v\n", err)
	}
}
