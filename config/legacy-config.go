package config

import (
	"encoding/json"
	"os"
	"time"

	"gopkg.in/yaml.v3"
	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// Legacy client configuration
type LegacyConfig struct {
	// ServiceAccount username
	// Required: true
	// Example : "osh_mikoev"
	Username string `json:"username" yaml:"username" validate:"required"`

	// ServiceAccount password
	// Required: true
	// Example: "[1o>hYkjnJr)HI78q7t&#%8Lm"
	Password string `json:"password" yaml:"password" validate:"required"`

	// Platform token
	// Required: false
	// Example: "158e76424b0d4810b6086hgbhj928fc4a6bc06e"
	Token string `json:"token" yaml:"token"`

	// Address of the platform on which the actions are planned
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
func (c *LegacyConfig) SetTimeout(dur time.Duration) {
	c.Timeout = Duration(dur)
}

// ParseLegacyConfigJSON parses LegacyConfig from specified JSON-formatted file.
func ParseLegacyConfigJSON(path string) (LegacyConfig, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return LegacyConfig{}, err
	}

	var config LegacyConfig

	err = json.Unmarshal(file, &config)
	if err != nil {
		return LegacyConfig{}, err
	}

	err = validators.ValidateConfig(config)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return LegacyConfig{}, validators.ValidationError(validationError)
		}
	}

	return config, nil
}

// ParseLegacyConfigYAML parses LegacyConfig from specified YAML-formatted file.
func ParseLegacyConfigYAML(path string) (LegacyConfig, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return LegacyConfig{}, err
	}

	var config LegacyConfig

	err = yaml.Unmarshal(file, &config)
	if err != nil {
		return LegacyConfig{}, err
	}

	err = validators.ValidateConfig(config)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return LegacyConfig{}, validators.ValidationError(validationError)
		}
	}

	return config, nil
}
