package gohttp

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"errors"
	"net"
	"net/http"
	"time"
)

const (
	defaultMaxIdleConnections = 5
	defaultRequestTimeout     = 5 * time.Second
	defaultConnectionTimeout  = 1 * time.Second
)

func (c *httpClient) getRequestBody(contentType string, body interface{}) ([]byte, error) {
	if body == nil {
		return nil, nil
	}

	switch contentType {
	case "application/json":
		return json.Marshal(body)

	case "application/xml":
		return xml.Marshal(body)

	default:
		return json.Marshal(body)
	}
}

func (c *httpClient) do(method string, url string, headers http.Header, body interface{}) (*http.Response, error) {
	fullHeaders := c.getRequestHeaders(headers)

	requestBody, err := c.getRequestBody(fullHeaders.Get("Content-type"), body)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest(method, url, bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, errors.New("unable to create a new request")
	}

	request.Header = fullHeaders

	client := c.getHttpClient()

	return client.Do(request)
}

func (c *httpClient) getHttpClient() *http.Client {
	if c.client != nil {
		return c.client
	}

	//once := sync.Once{}
	c.clientOnce.Do(func() {
		c.client = &http.Client{
			Timeout: c.getConnectionTimeout() + c.getRequestTimeout(),
			Transport: &http.Transport{
				MaxIdleConnsPerHost:   c.getMaxIdleConnection(),
				ResponseHeaderTimeout: c.getRequestTimeout(),
				DialContext: net.Dialer{
					Timeout: c.getConnectionTimeout(),
				}.Resolver.Dial,
			},
		}
	})

	return c.client
}

func (c *httpClient) getMaxIdleConnection() int {
	if c.builder.maxIdleConnections > 0 {
		return c.builder.maxIdleConnections
	}

	return defaultMaxIdleConnections
}

func (c *httpClient) getRequestTimeout() time.Duration {
	if c.builder.requestTimeout > 0 {
		return c.builder.requestTimeout
	}

	if c.builder.disableTimetous {
		return 0
	}

	return defaultRequestTimeout
}

func (c *httpClient) getConnectionTimeout() time.Duration {
	if c.builder.connectionTimeout > 0 {
		return c.builder.connectionTimeout
	}

	return defaultConnectionTimeout
}

func (c *httpClient) getRequestHeaders(requestHeaders http.Header) http.Header {
	result := make(http.Header)
	// Add common headers to the request:
	for header, value := range c.builder.headers {
		if len(value) > 0 {
			result.Set(header, value[0])
		}
	}

	// Add custom headers to the request:
	for header, value := range requestHeaders {
		if len(value) > 0 {
			result.Set(header, value[0])
		}
	}
	return result
}
