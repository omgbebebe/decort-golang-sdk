package client

import (
	"crypto/tls"
	"net/http"
	"net/url"
	"time"

	"repos.digitalenergy.online/BASIS/decort-golang-sdk/config"
)

// NewLegacyHttpClient creates legacy HTTP Client
func NewLegacyHttpClient(cfg config.LegacyConfig) *http.Client {
	transCfg := &http.Transport{
		TLSClientConfig: &tls.Config{
			//nolint:gosec
			InsecureSkipVerify: cfg.SSLSkipVerify,
		},
	}

	return &http.Client{
		Transport: &transportLegacy{
			base:      transCfg,
			username:  url.QueryEscape(cfg.Username),
			password:  url.QueryEscape(cfg.Password),
			retries:   cfg.Retries,
			token:     cfg.Token,
			decortURL: cfg.DecortURL,
		},

		Timeout: 5 * time.Minute,
	}
}
