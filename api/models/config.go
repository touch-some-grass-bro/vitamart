package models

import (
	"fmt"
)

var Config Configuration

type DatabaseConf struct {
	Host     string `mapstructure:"host"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Name     string `mapstructure:"name"`
	Port     string `mapstructure:"port"`
}

func (d *DatabaseConf) URI() string {
	uri := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", d.User, d.Password, d.Host, d.Port, d.Name)
	return uri
}

type GoogleConf struct {
	ClientID     string `mapstructure:"client_id"`
	ClientSecret string `mapstructure:"client_secret"`
	RedirectURI  string `mapstructure:"redirect_uri"`
}
// TODO: Change to google config

type APIConf struct {
	JWTSecretKey    string   `mapstructure:"jwt_secret_key"`
	Host            string   `mapstructure:"host"`
	Port            string   `mapstructure:"port"`
	OAuthState      string   `mapstructure:"oauth_state"`
  FrontendUrl     string   `mapstructure:"frontend_url"`
}

type Configuration struct {
	API      *APIConf      `mapstructure:"api"`
	Database *DatabaseConf `mapstructure:"database"`
	Google  *GoogleConf  `mapstructure:"google"`
}
