package numbers

import (
	"net/http"

	"github.com/tgbv/telnyx-golang/internal"
)

// aliases
type Config = internal.Config

var e func(int, []byte) error = internal.E

/*
*	handles numbers related operations
 */
type Numbers struct {
	Config     *Config
	HttpClient *http.Client
}

/*
*	initializes the numbers
 */
func InitNumbers(c *Config, hc *http.Client) Numbers {
	n := Numbers{}

	n.Config = c
	n.HttpClient = hc

	return n
}
