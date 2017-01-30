package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/mitchellh/go-homedir"
)

const (
	configFileName         = "config.json"
	alfredWorkflowDataPath = "Library/Application Support/Alfred 3/Workflow Data/"
)

type config struct {
	AccessToken string `json:"accessToken"`
	TeamName    string `json:"teamName"`
}

func getDefaultConfigPath() (string, error) {
	homeDir, err := homedir.Dir()
	if err != nil {
		return "", err
	}
	return filepath.Join(homeDir, alfredWorkflowDataPath, bundleID), nil
}

func getDefaultConfigFilePath() (string, error) {
	configPath, err := getDefaultConfigPath()
	if err != nil {
		return "", err
	}
	return filepath.Join(configPath, configFileName), nil
}

func loadConfig() (config, error) {
	configFilePath, err := getDefaultConfigFilePath()
	if err != nil {
		return config{}, err
	}

	c := config{}
	data, err := ioutil.ReadFile(configFilePath)
	if err != nil {
		return config{}, err
	}

	return c, json.Unmarshal(data, &c)
}

func saveConfig(c config) error {
	configFilePath, err := getDefaultConfigFilePath()
	if err != nil {
		return err
	}

	data, err := json.Marshal(&c)
	if err != nil {
		return err
	}

	err = os.MkdirAll(filepath.Dir(configFilePath), 0700)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(configFilePath, data, 0666)
}
