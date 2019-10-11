package config

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
)

var Config Configuration

// Load ...
func Load() error {
	_, err := loadConfigFile(FilePath)
	return err
}

func loadConfigFile(filePath string) (*Configuration, error) {
	configFileContents, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Printf("Reading configuration from JSON (%s) failed (err: %v). SKIPPED.\n", filePath, err)
		return nil, err
	}

	reader := bytes.NewBuffer(configFileContents)

	err = json.NewDecoder(reader).Decode(&Config)
	if err != nil {
		return nil, err
	}

	return &Config, nil
}
