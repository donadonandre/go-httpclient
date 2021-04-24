package gohttp

import (
	"net/http"
	"time"
)

type clientBuilder struct {
	headers            http.Header
	maxIdleConnections int
	connectionTimeout  time.Duration
	requestTimeout     time.Duration
	disableTimetous    bool
}

type ClientBuilder interface {
	SetHeaders(headers http.Header) ClientBuilder
	SetConnectionTimeout(timeout time.Duration) ClientBuilder
	SetRequestTimeout(timeout time.Duration) ClientBuilder
	SetMaxIdleConnections(connections int) ClientBuilder
	DisableTimeouts(disable bool) ClientBuilder
	Build() Client

	/*
		Essa definição acima de retornar ClientBuilder nos métodos é justamente para poder contatenar o uso dos mẽtodos
		somente usando o ponto (.)
		Ex:
		cliente := gohttp.NewBuilder().DisableTimeouts().Build()
	*/
}

func NewBuilder() ClientBuilder {
	// client := http.Client{
	// 	Transport: &http.Transport{
	// 		MaxIdleConnsPerHost:   5,
	// 		ResponseHeaderTimeout: 5 * time.Second,
	// 		DialContext: net.Dialer{
	// 			Timeout: 1 * time.Second,
	// 		}.Resolver.Dial,
	// 	},
	// }

	// httpClient := &httpClient{
	// 	client: &client,
	// }
	// return httpClient
	builder := &clientBuilder{}
	return builder
}

func (c *clientBuilder) Build() Client {
	client := httpClient{
		headers:            c.headers,
		maxIdleConnections: c.maxIdleConnections,
		connectionTimeout:  c.connectionTimeout,
		requestTimeout:     c.requestTimeout,
		disableTimetous:    c.disableTimetous,
	}
	return &client
}

func (c *clientBuilder) SetHeaders(headers http.Header) ClientBuilder {
	c.headers = headers
	return c
}

func (c *clientBuilder) SetConnectionTimeout(timeout time.Duration) ClientBuilder {
	c.connectionTimeout = timeout
	return c
}

func (c *clientBuilder) SetRequestTimeout(timeout time.Duration) ClientBuilder {
	c.requestTimeout = timeout
	return c
}

func (c *clientBuilder) SetMaxIdleConnections(connections int) ClientBuilder {
	c.maxIdleConnections = connections
	return c
}

func (c *clientBuilder) DisableTimeouts(disable bool) ClientBuilder {
	c.disableTimetous = disable
	return c
}
