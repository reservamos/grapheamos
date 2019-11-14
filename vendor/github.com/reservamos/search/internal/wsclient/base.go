package wsclient

import (
	"crypto/tls"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/reservamos/search/internal/config"
)

// WebService structure with http client, base url and auth params
type WebService struct {
	client *http.Client
	URL    url.URL
	auth   string
}

func NewWebService() WebService {
	baseURL := config.App.IntegrationsURL
	auth := config.App.IntegrationsAuthHeader
	return NewGenericWebService(baseURL, auth)
}

// NewWebService builds a WebService with params read from config
func NewGenericWebService(baseURL, auth string) WebService {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	urlParts := strings.Split(baseURL, "://")
	return WebService{
		client: &http.Client{Timeout: 100 * time.Second, Transport: tr},
		auth:   auth,
		URL: url.URL{
			Scheme: urlParts[0],
			Host:   urlParts[1],
		},
	}
}

// NewWSRequest Builds a request with default params and given methods, url and body
func (ws WebService) NewWSRequest(method string, url string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return req, err
	}
	req.Header.Set("Authorization", ws.auth)
	return req, err
}

// DoWSRequest sends a request with given params to server
func (ws WebService) DoWSRequest(method string, url string, body io.Reader) (*http.Response, error) {
	var res *http.Response

	req, err := ws.NewWSRequest(method, url, body)
	if err != nil {
		return res, err
	}

	res, err = ws.client.Do(req)
	if err != nil {
		return res, err
	}

	return res, err
}
