package gohttp

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"errors"
	"net/http"
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

	return c.client.Do(request)
}

func (c *httpClient) getRequestHeaders(requestHeaders http.Header) http.Header {
	result := make(http.Header)
	// Add common headers to the request:
	for header, value := range c.Headers {
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
