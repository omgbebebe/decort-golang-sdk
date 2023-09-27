package decortsdk

import (
	"bytes"
	"context"
	"crypto/tls"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"

	"github.com/google/go-querystring/query"
	"repository.basistech.ru/BASIS/decort-golang-sdk/config"
	"repository.basistech.ru/BASIS/decort-golang-sdk/pkg/cloudapi"
	"repository.basistech.ru/BASIS/decort-golang-sdk/pkg/cloudbroker"
)

// Legacy HTTP-client for platform
type LegacyDecortClient struct {
	decortURL  string
	client     *http.Client
	cfg        config.LegacyConfig
	expiryTime time.Time
	mutex      *sync.Mutex
}

// Legacy client builder
func NewLegacy(cfg config.LegacyConfig) *LegacyDecortClient {
	if cfg.Retries == 0 {
		cfg.Retries = 5
	}

	var expiryTime time.Time

	if cfg.Token != "" {
		expiryTime = time.Now().AddDate(0, 0, 1)
	}

	return &LegacyDecortClient{
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
func (ldc *LegacyDecortClient) CloudAPI() *cloudapi.CloudAPI {
	return cloudapi.New(ldc)
}

// CloudBroker builder
func (ldc *LegacyDecortClient) CloudBroker() *cloudbroker.CloudBroker {
	return cloudbroker.New(ldc)
}

// DecortApiCall method for sending requests to the platform
func (ldc *LegacyDecortClient) DecortApiCall(ctx context.Context, method, url string, params interface{}) ([]byte, error) {
	values, err := query.Values(params)
	if err != nil {
		return nil, err
	}

	if err = ldc.getToken(ctx); err != nil {
		return nil, err
	}

	body := strings.NewReader(values.Encode() + fmt.Sprintf("&authkey=%s", ldc.cfg.Token))

	req, err := http.NewRequestWithContext(ctx, method, ldc.decortURL+"/restmachine"+url, body)
	if err != nil {
		return nil, err
	}

	resp, err := ldc.do(req)
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

func (ldc *LegacyDecortClient) getToken(ctx context.Context) error {
	ldc.mutex.Lock()
	defer ldc.mutex.Unlock()

	if ldc.cfg.Token == "" || time.Now().After(ldc.expiryTime) {
		body := fmt.Sprintf("username=%s&password=%s", url.QueryEscape(ldc.cfg.Username), url.QueryEscape(ldc.cfg.Password))
		bodyReader := strings.NewReader(body)

		req, _ := http.NewRequestWithContext(ctx, "POST", ldc.cfg.DecortURL+"/restmachine/cloudapi/user/authenticate", bodyReader)
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

		resp, err := ldc.client.Do(req)
		if err != nil {
			return fmt.Errorf("unable to get token: %w", err)
		}

		tokenBytes, _ := io.ReadAll(resp.Body)
		resp.Body.Close()

		if resp.StatusCode != 200 {
			return fmt.Errorf("unable to get token: %s", tokenBytes)
		}

		token := string(tokenBytes)
		ldc.cfg.Token = token
		ldc.expiryTime = time.Now().AddDate(0, 0, 1)
	}

	return nil
}

func (ldc *LegacyDecortClient) do(req *http.Request) (*http.Response, error) {
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Accept", "application/json")

	// var resp *http.Response
	// var err error
	buf, _ := io.ReadAll(req.Body)

	// for i := uint64(0); i < ldc.cfg.Retries; i++ {
		// req = req.Clone(req.Context())
		req.Body = io.NopCloser(bytes.NewBuffer(buf))
		resp, err := ldc.client.Do(req)

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
