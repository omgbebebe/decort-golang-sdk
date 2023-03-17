package client

import (
	"crypto/tls"
	"net/http"
	"time"

	"repository.basistech.ru/BASIS/decort-golang-sdk/config"
)

func NewHttpClient(cfg config.Config) *http.Client {

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
		Transport: &transport{
			base:         transCfg,
			retries:      cfg.Retries,
			clientID:     cfg.AppID,
			clientSecret: cfg.AppSecret,
			ssoURL:       cfg.SSOURL,
			token:        cfg.Token,
			expiryTime:   expiredTime,
			//TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},

		Timeout: 5 * time.Minute,
	}
}
