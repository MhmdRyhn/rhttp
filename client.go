package rhttp

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

type HTTPClient struct {
	// Request options
	Options RequestOptions

	// GoLanf standard http client.
	stdClient http.Client
}

func (httpClient *HTTPClient) Get(url string, requestHeaders map[string]string) (*http.Response, error) {
	if baseUrl := httpClient.Options.BaseUrl; baseUrl == "" {
		url = fmt.Sprintf("%s/%s", baseUrl, strings.TrimLeft(url, "/"))
	}
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	for headerKey, headerValue := range requestHeaders {
		request.Header.Set(headerKey, headerValue)
	}
	return httpClient.stdClient.Do(request)
}

func (httpClient *HTTPClient) Post(
	url string, body io.Reader, requestHeaders map[string]string,
) (*http.Response, error) {
	request, err := http.NewRequest(http.MethodPost, url, body)
	if err != nil {
		return nil, err
	}
	for headerKey, headerValue := range requestHeaders {
		request.Header.Set(headerKey, headerValue)
	}
	return httpClient.stdClient.Do(request)
}
