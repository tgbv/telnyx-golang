package messaging

import (
	"net/http"

	"github.com/tgbv/telnyx-golang/internal"
)

// aliases
type Config = internal.Config

var e func(int, []byte) error = internal.E

/*
*	handles messaging related operations
 */
type Messaging struct {
	Config     *Config
	HttpClient *http.Client
	WebHook    webHook
}

/*
*	initializes messages with config
 */
func InitMessaging(c *Config, hc *http.Client) *Messaging {
	m := Messaging{}

	m.Config = c
	m.HttpClient = hc

	return &m
}
