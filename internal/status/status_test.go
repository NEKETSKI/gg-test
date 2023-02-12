package status

import (
	"errors"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

type roundTripFunc func(*http.Request) (*http.Response, error)

func (r roundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return r(req)
}

func TestGetEndpointStatus(t *testing.T) {
	tests := []struct {
		name    string
		client  *http.Client
		url     string
		want    int
		wantErr assert.ErrorAssertionFunc
	}{
		{
			"Status code 200",
			&http.Client{Transport: roundTripFunc(func(req *http.Request) (*http.Response, error) {
				assert.Equal(t, "http://epample.com", req.URL.String())
				return &http.Response{StatusCode: 200}, nil
			})},
			"http://epample.com",
			200,
			assert.NoError,
		},
		{
			"Status code 404",
			&http.Client{Transport: roundTripFunc(func(req *http.Request) (*http.Response, error) {
				return &http.Response{StatusCode: 404}, nil
			})},
			"",
			404,
			assert.NoError,
		},
		{
			"Request error",
			&http.Client{Transport: roundTripFunc(func(req *http.Request) (*http.Response, error) {
				return nil, errors.New("some request error happens")
			})},
			"",
			0,
			assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetEndpointStatus(tt.client, tt.url)
			if !tt.wantErr(t, err, fmt.Sprintf("GetEndpointStatus(%v, %v)", tt.client, tt.url)) {
				return
			}
			assert.Equalf(t, tt.want, got, "GetEndpointStatus(%v, %v)", tt.client, tt.url)
		})
	}
}
