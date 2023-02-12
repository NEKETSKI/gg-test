// Package status provides a primitive for calling the provided endpoint and handling the response.
package status

import (
	"fmt"
	"net/http"
)

// GetEndpointStatus makes a "GET" request to the provided URL and returns the status code from the response.
// Returns an error if the HTTP call failed.
func GetEndpointStatus(client *http.Client, url string) (int, error) {
	req, err := http.NewRequest(http.MethodGet, url, http.NoBody)
	if err != nil {
		return 0, fmt.Errorf("failed to prepare request: %w", err)
	}
	resp, err := client.Do(req)
	if err != nil {
		return 0, fmt.Errorf("failed to make a request: %w", err)
	}
	return resp.StatusCode, nil
}
