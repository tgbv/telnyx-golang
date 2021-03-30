package verify

import (
	"net/http"

	"github.com/tgbv/telnyx-golang/internal"
)

// aliases
type Config = internal.Config

var e func(int, []byte) error = internal.E

/*
*	handles verification related operations
	see https://portal.telnyx.com/#/app/verify/profiles
*/
type Verify struct {
	Config     *Config
	HttpClient *http.Client
}

/*
*	initializes verify
 */
func InitVerify(c *Config, hc *http.Client) Verify {
	n := Verify{}

	n.Config = c
	n.HttpClient = hc

	return n
}
