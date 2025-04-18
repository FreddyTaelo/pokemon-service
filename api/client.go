// api/client.go
package api

import (
	"net/http"
	"time"
)

var client = &http.Client{
	Timeout: 10 * time.Second,
}

func GetClient() *http.Client {
	return client
} /*GetClient()*/

func SetClient(c *http.Client) { // client for testing
	client = c
}
