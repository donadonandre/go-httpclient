package gohttp

import (
	"fmt"
	"net/http"
	"testing"
)

func TestGetRequestHeaders(t *testing.T) {
	// Initialization
	client := httpClient{}
	commonHeaders := make(http.Header)
	commonHeaders.Set("Content-Type", "application/json")
	commonHeaders.Set("User-Agent", "cool-http-client")
	client.builder.headers = commonHeaders

	// Execution
	requestHeaders := make(http.Header)
	requestHeaders.Set("X-Request-Id", "ABC-123")

	finalHeaders := client.getRequestHeaders(requestHeaders)

	// Validation
	if len(finalHeaders) != 3 {
		t.Error("We expect 3 headers")
	}

	if finalHeaders.Get("Content-Type") != "application/json" {
		t.Error("Invalid content type id received")
	}

	if finalHeaders.Get("User-Agent") != "cool-http-client" {
		t.Error("Invalid user agent received")
	}

	if finalHeaders.Get("X-Request-Id") != "ABC-123" {
		t.Error("Invalid request id received")
	}
}

func TestGetRequestBodyNilBody(t *testing.T) {
	// Initialization
	client := httpClient{}

	t.Run("noBodyNilResponse", func(t *testing.T) {
		// Execution
		body, err := client.getRequestBody("", nil)

		// Validation
		if err != nil {
			t.Error("No error expecting when passing a nil body")
		}

		if body != nil {
			t.Error("No body expecting when passing a nil body")
		}
	})

	t.Run("BodyWithJson", func(t *testing.T) {
		// Execution
		requestBody := []string{"one", "two"}
		body, err := client.getRequestBody("application/json", requestBody)

		// Validation
		if err != nil {
			t.Error("No error expecting when marshaling slice as json")
		}

		if string(body) != `["one","two"]` {
			t.Error("Invalid json body obtained")
		}
	})

	t.Run("BodyWithXml", func(t *testing.T) {
		// Execution
		requestBody := []string{"one", "two"}
		body, err := client.getRequestBody("application/xml", requestBody)

		// Validation
		if err != nil {
			t.Error("No error expecting when marshaling slice as xml")
		}

		if string(body) != `<string>one</string><string>two</string>` {
			fmt.Println("body:", string(body))
			t.Error("Invalid xml body obtained")
		}
	})

	t.Run("BodyWithJsonAsDefault", func(t *testing.T) {
		// Execution
		requestBody := []string{"one", "two"}
		body, err := client.getRequestBody("", requestBody)

		// Validation
		if err != nil {
			t.Error("No error expecting when marshaling slice as json as default")
		}

		if string(body) != `["one","two"]` {
			t.Error("Invalid json body obtained")
		}
	})
}
