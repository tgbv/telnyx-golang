package internal

import (
	"net/http"
	"time"
)

/*
*	this type holds a HTTP client and a transport setup for HTTPS connections
 */
type client struct {
	tr     http.Transport
	Client http.Client
}

/*
*	initializes a client
 */
func InitHttpClient() http.Client {
	c := client{
		tr: http.Transport{
			MaxIdleConns:       10,
			IdleConnTimeout:    30 * time.Second,
			DisableCompression: true,
		},
	}

	c.Client = http.Client{Transport: &c.tr}

	return c.Client
}
