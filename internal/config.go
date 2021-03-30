package internal

import "github.com/tgbv/telnyx-golang/config"

/*
*	holds the telnyx configuration
 */
type Config struct {
	Api       config.ApiKeys
	Messaging config.Messaging
}
