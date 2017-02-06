package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/upamune/go-esa/esa"
)

const (
	configFileName = "config.json"
)

// Workflow alfred workflow
type Workflow struct {
	AppName  string
	Version  string
	DataPath string
	Config   Config
	Client   *esa.Client
	Alfred   *Alfred
	// Service
	Setup  *SetupService
	Search *SearchService
	Today  *TodayService
}

func (w Workflow) buildConfigFilePath() string {
	return filepath.Join(w.DataPath, configFileName)
}

func (w *Workflow) loadConfig() error {
	data, err := ioutil.ReadFile(w.buildConfigFilePath())
	if err != nil {
		return err
	}
	return json.Unmarshal(data, &w.Config)
}

// SaveConfig save config with create file and mkdir
func (w *Workflow) SaveConfig() error {
	configFilePath := w.buildConfigFilePath()
	data, err := json.Marshal(&w.Config)
	if err != nil {
		return err
	}

	err = os.MkdirAll(filepath.Dir(configFilePath), 0700)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(configFilePath, data, 0666)
}

// NewWorkflow create workflow
func NewWorkflow() *Workflow {
	// https://www.alfredapp.com/help/workflows/script-environment-variables/
	w := &Workflow{
		AppName:  os.Getenv("alfred_workflow_bundleid"),
		Version:  os.Getenv("alfred_version"),
		DataPath: os.Getenv("alfred_workflow_data"),
	}
	if err := w.loadConfig(); err == nil {
		w.Client = esa.NewClient(w.Config.AccessToken)
	}

	w.Alfred = &Alfred{}
	w.Setup = &SetupService{Workflow: w}
	w.Search = &SearchService{Workflow: w}
	w.Today = &TodayService{Workflow: w}
	return w
}
