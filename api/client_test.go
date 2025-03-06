// api/client_test.go
package api

import (
	"testing"
	"time"
)

func TestGetClient(t *testing.T) {

	client := GetClient()

	// Check if the client is not nil
	if client == nil {
		t.Errorf("Expected client to be non-nil, got nil")
	}

	// Check if the timeout is correctly set to 10 seconds
	expectedTimeout := 10 * time.Second
	if client.Timeout != expectedTimeout {
		t.Errorf("Expected client timeout to be %v, got %v", expectedTimeout, client.Timeout)
	}
}
