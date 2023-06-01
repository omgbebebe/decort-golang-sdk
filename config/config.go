package config

import (
	"encoding/json"
	"os"
	"time"

	"gopkg.in/yaml.v3"
	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// Configuration for creating request to platform
type Config struct {
	// JWT platform token
	// Required: false
	// Example: "qwqwdfwv68979we0q9bfv7e9sbvd89798qrwv97ff"
	Token string `json:"token" yaml:"token"`

	// Application (client) identifier for authorization
	// in the cloud platform controller in oauth2 mode.
	// Required: true
	// Example: "ewqfrvea7s890avw804389qwguf234h0otfi3w4eiu"
	AppID string `json:"appId" yaml:"appId" validate:"required"`

	// Application (client) secret code for authorization
	// in the cloud platform controller in oauth2 mode.
	// Example: "frvet09rvesfis0c9erv9fsov0vsdfi09ovds0f"
	AppSecret string `json:"appSecret" yaml:"appSecret" validate:"required"`

	// Platform authentication service address
	// Required: true
	// Example: "https://sso.digitalenergy.online"
	SSOURL string `json:"ssoUrl" yaml:"ssoUrl" validate:"url"`

	// The address of the platform on which the actions are planned
	// Required: true
	// Example: "https://mr4.digitalenergy.online"
	DecortURL string `json:"decortUrl" yaml:"decortUrl" validate:"url"`

	// Amount platform request attempts
	// Default value: 5
	// Required: false
	Retries uint64 `json:"retries" yaml:"retries"`

	// Skip verify, true by default
	// Required: false
	SSLSkipVerify bool `json:"sslSkipVerify" yaml:"sslSkipVerify"`

	// HTTP client timeout, unlimited if left empty
	// Required: false
	Timeout Duration `json:"timeout" yaml:"timeout"`
}

// SetTimeout is used to set HTTP client timeout.
func (c *Config) SetTimeout(dur time.Duration) {
	c.Timeout = Duration(dur)
}

// ParseConfigJSON parses Config from specified JSON-formatted file.
func ParseConfigJSON(path string) (Config, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return Config{}, err
	}

	var config Config

	err = json.Unmarshal(file, &config)
	if err != nil {
		return Config{}, err
	}

	err = validators.ValidateConfig(config)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return Config{}, validators.ValidationError(validationError)
		}
	}

	return config, nil
}

// ParseConfigYAML parses Config from specified YAML-formatted file.
func ParseConfigYAML(path string) (Config, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return Config{}, err
	}

	var config Config

	err = yaml.Unmarshal(file, &config)
	if err != nil {
		return Config{}, err
	}

	err = validators.ValidateConfig(config)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return Config{}, validators.ValidationError(validationError)
		}
	}

	return config, nil
}
