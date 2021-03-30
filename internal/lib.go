package internal

import "fmt"

/*
*	sends standard formatted error
 */
func E(code int, body []byte) error {
	return e(code, body)
}
func e(code int, body []byte) error {
	return fmt.Errorf("API error. Code: %d\r\n\r\n%s", code, body)
}
