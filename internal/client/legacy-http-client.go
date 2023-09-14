package client

import (
	"crypto/tls"
	"net/http"
	"net/url"
	"sync"
	"time"

	"repository.basistech.ru/BASIS/decort-golang-sdk/config"
)

// NewLegacyHttpClient creates legacy HTTP Client
func NewLegacyHttpClient(cfg config.LegacyConfig) *http.Client {
	transCfg := &http.Transport{
		TLSClientConfig: &tls.Config{
			//nolint:gosec
			InsecureSkipVerify: cfg.SSLSkipVerify,
		},
	}

	var expiredTime time.Time

	if cfg.Token != "" {
		expiredTime = time.Now().AddDate(0, 0, 1)
	}

	return &http.Client{
		Transport: &transportLegacy{
			base:       transCfg,
			username:   url.QueryEscape(cfg.Username),
			password:   url.QueryEscape(cfg.Password),
			retries:    cfg.Retries,
			token:      cfg.Token,
			decortURL:  cfg.DecortURL,
			expiryTime: expiredTime,
			mutex:      &sync.Mutex{},
		},

		Timeout: cfg.Timeout.Get(),
	}
}
