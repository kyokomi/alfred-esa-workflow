package main

import "github.com/upamune/go-esa/esa"

type esaWorkflow struct {
	client   *esa.Client
	teamName string
}

func newEsaWorkflowFromConfig() (*esaWorkflow, error) {
	c, err := loadConfig()
	if err != nil {
		return nil, err
	}
	return newEsaWorkflow(c.AccessToken, c.TeamName), nil
}

func newEsaWorkflow(accessToken, teamName string) *esaWorkflow {
	return &esaWorkflow{
		client:   esa.NewClient(accessToken),
		teamName: teamName,
	}
}
