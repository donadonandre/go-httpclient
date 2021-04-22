package gohttp

import (
	"net"
	"net/http"
	"time"
)

type httpClient struct {
	client  *http.Client
	Headers http.Header
}

func New() HttpClient {
	client := http.Client{
		Transport: &http.Transport{
			MaxIdleConnsPerHost:   5,
			ResponseHeaderTimeout: 5 * time.Second,
			DialContext: net.Dialer{
				Timeout: 1 * time.Second,
			}.Resolver.Dial,
		},
	}

	httpClient := &httpClient{
		client: &client,
	}
	return httpClient
}

type HttpClient interface {
	SetHeaders(headers http.Header)

	Get(url string, headers http.Header) (*http.Response, error)
	Post(url string, headers http.Header, body interface{}) (*http.Response, error)
	Put(url string, headers http.Header, body interface{}) (*http.Response, error)
	Patch(url string, headers http.Header, body interface{}) (*http.Response, error)
	Delete(url string, headers http.Header, body interface{}) (*http.Response, error)
}

func (c *httpClient) SetHeaders(headers http.Header) {
	c.Headers = headers
}

func (c *httpClient) Get(url string, headers http.Header) (*http.Response, error) {
	return c.do(http.MethodGet, url, headers, nil)
}

func (c *httpClient) Post(url string, headers http.Header, body interface{}) (*http.Response, error) {
	return c.do(http.MethodPost, url, headers, nil)
}

func (c *httpClient) Put(url string, headers http.Header, body interface{}) (*http.Response, error) {
	return c.do(http.MethodPut, url, headers, nil)
}

func (c *httpClient) Patch(url string, headers http.Header, body interface{}) (*http.Response, error) {
	return c.do(http.MethodPatch, url, headers, nil)
}

func (c *httpClient) Delete(url string, headers http.Header, body interface{}) (*http.Response, error) {
	return c.do(http.MethodDelete, url, headers, nil)
}
