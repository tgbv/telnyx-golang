package messaging

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"
)

/*
*	holds the webhook server data regarding telnyx webhook messaging events
	see https://developers.telnyx.com/docs/api/v2/overview#webhooks and https://developers.telnyx.com/docs/v2/messaging/receiving-webhooks
*/
type webHook struct {
	HttpServer http.Server
	callbacks  []func(*http.Request)
	mux        sync.Mutex
}

/*
*	root request handler which executes the callbacks, then writes a common response
 */
func (W *webHook) ReqHandler(rw http.ResponseWriter, r *http.Request) {
	W.mux.Lock()
	for _, f := range W.callbacks {
		f(r)
	}
	W.mux.Unlock()

	rw.Header().Add("Content-type", "application/json")
	rw.Header().Add("Connection", "close")
	rw.Header().Add("Server", "Go")
}

/*
*	appends a callback to webhook callbacks queue
 */
func (W *webHook) PushCb(f func(*http.Request)) *webHook {
	W.mux.Lock()
	W.callbacks = append(W.callbacks, f)
	W.mux.Unlock()

	return W
}

/*
*	deletes a callback from slice
 */
func (W *webHook) DelCb(index uint) (*webHook, error) {
	W.mux.Lock()

	// check index
	l := len(W.callbacks)
	if index < 0 || index > uint(l) {
		W.mux.Unlock()
		return W, fmt.Errorf("Index out of boundaries: %d out of [0, %d]", index, l-1)
	}

	// do it.
	old := W.callbacks
	W.callbacks = old[:index]
	W.callbacks = append(W.callbacks, old[index+1:]...)

	W.mux.Unlock()

	return W, nil
}

/*
*	starts the webhock server
 */
func (W *webHook) StartServer() error {
	err := W.HttpServer.ListenAndServe()
	if err != nil {
		return err
	}

	return nil
}

/*
*	stops the server
 */
func (W *webHook) StopServer() {
	W.HttpServer.Shutdown(context.Background())
}

/*
*	initializes webhook
 */
func (M *Messaging) InitWebhook(host string) (*webHook, error) {
	M.WebHook = webHook{}

	M.WebHook.HttpServer = http.Server{
		Addr:           host,
		Handler:        http.HandlerFunc(M.WebHook.ReqHandler),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	M.WebHook.callbacks = make([]func(*http.Request), 0)

	M.WebHook.mux = sync.Mutex{}

	return &M.WebHook, nil
}
