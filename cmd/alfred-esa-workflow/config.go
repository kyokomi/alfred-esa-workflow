package main

// Config esa config
type Config struct {
	AccessToken string `json:"accessToken"`
	TeamName    string `json:"teamName"`
}

// IsValid return config valid
func (c Config) IsValid() bool {
	return len(c.AccessToken) > 0 && len(c.TeamName) > 0
}
