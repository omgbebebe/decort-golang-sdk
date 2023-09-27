package decortsdk

import (
	"bytes"
	"context"
	"crypto/tls"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"sync"
	"time"

	"repository.basistech.ru/BASIS/decort-golang-sdk/pkg/cloudapi"
	"repository.basistech.ru/BASIS/decort-golang-sdk/pkg/cloudbroker"

	"github.com/google/go-querystring/query"
	"repository.basistech.ru/BASIS/decort-golang-sdk/config"
)

// HTTP-client for platform
type DecortClient struct {
	decortURL  string
	client     *http.Client
	cfg        config.Config
	expiryTime time.Time
	mutex      *sync.Mutex
}

// Сlient builder
func New(cfg config.Config) *DecortClient {
	if cfg.Retries == 0 {
		cfg.Retries = 5
	}

	var expiryTime time.Time

	if cfg.Token != "" {
		expiryTime = time.Now().AddDate(0, 0, 1)
	}

	return &DecortClient{
		decortURL: cfg.DecortURL,
		client: &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					//nolint:gosec
					InsecureSkipVerify: cfg.SSLSkipVerify,
				},
			},
		},
		cfg:        cfg,
		expiryTime: expiryTime,
		mutex:      &sync.Mutex{},
	}
}

// CloudAPI builder
func (dc *DecortClient) CloudAPI() *cloudapi.CloudAPI {
	return cloudapi.New(dc)
}

// CloudBroker builder
func (dc *DecortClient) CloudBroker() *cloudbroker.CloudBroker {
	return cloudbroker.New(dc)
}

// DecortApiCall method for sending requests to the platform
func (dc *DecortClient) DecortApiCall(ctx context.Context, method, url string, params interface{}) ([]byte, error) {
	values, err := query.Values(params)
	if err != nil {
		return nil, err
	}

	body := strings.NewReader(values.Encode())
	req, err := http.NewRequestWithContext(ctx, method, dc.decortURL+"/restmachine"+url, body)
	if err != nil {
		return nil, err
	}

	if err = dc.getToken(ctx); err != nil {
		return nil, err
	}

	resp, err := dc.do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, errors.New(string(respBytes))
	}

	return respBytes, nil
}

func (dc *DecortClient) getToken(ctx context.Context) error {
	dc.mutex.Lock()
	defer dc.mutex.Unlock()

	if dc.cfg.Token == "" || time.Now().After(dc.expiryTime) {
		body := fmt.Sprintf("grant_type=client_credentials&client_id=%s&client_secret=%s&response_type=id_token", dc.cfg.AppID, dc.cfg.AppSecret)
		bodyReader := strings.NewReader(body)

		dc.cfg.SSOURL = strings.TrimSuffix(dc.cfg.SSOURL, "/")

		req, _ := http.NewRequestWithContext(ctx, "POST", dc.cfg.SSOURL+"/v1/oauth/access_token", bodyReader)
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

		resp, err := dc.client.Do(req)
		if err != nil {
			return fmt.Errorf("cannot get token: %w", err)
		}

		tokenBytes, _ := io.ReadAll(resp.Body)
		resp.Body.Close()

		if resp.StatusCode != 200 {
			return fmt.Errorf("cannot get token: %s", tokenBytes)
		}

		token := string(tokenBytes)

		dc.cfg.Token = token
		dc.expiryTime = time.Now().AddDate(0, 0, 1)
	}

	return nil
}

func (dc *DecortClient) do(req *http.Request) (*http.Response, error) {
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Authorization", "bearer "+dc.cfg.Token)
	req.Header.Set("Accept", "application/json")

	// var resp *http.Response
	// var err error
	buf, _ := io.ReadAll(req.Body)

	// for i := uint64(0); i < dc.cfg.Retries; i++ {
		// req = req.Clone(req.Context())
		req.Body = io.NopCloser(bytes.NewBuffer(buf))
		resp, err := dc.client.Do(req)

		// if err == nil {
			if resp.StatusCode == 200 {
				return resp, err
			}
			respBytes, _ := io.ReadAll(resp.Body)
			err = fmt.Errorf("%s", respBytes)
			resp.Body.Close()
		// }
	// }

	return nil, fmt.Errorf("could not execute request: %w", err)
}
