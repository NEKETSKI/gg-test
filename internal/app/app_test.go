package app

import (
	"errors"
	"fmt"
	"net/http"
	"testing"

	"github.com/NEKETSKY/gg-test/internal/repository"
	"github.com/NEKETSKY/gg-test/pkg/logger"
	"github.com/NEKETSKY/gg-test/pkg/logger/testlogger"
	"github.com/stretchr/testify/assert"
)

type roundTripFunc func(*http.Request) (*http.Response, error)

func (r roundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return r(req)
}

type repoMock struct {
	storeStatusCodeMock func(requestUrl string, code int) error
	closeMock           func() error
}

func (m repoMock) StoreStatusCode(requestUrl string, code int) error {
	return m.storeStatusCodeMock(requestUrl, code)
}

func (m repoMock) Close() error {
	return m.closeMock()
}

func TestService_checkStatus(t *testing.T) {
	type fields struct {
		log    logger.Logger
		client *http.Client
		repo   repository.Repository
	}
	type args struct {
		endpoint string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr assert.ErrorAssertionFunc
	}{
		{
			"No error",
			fields{
				log: testlogger.New(),
				client: &http.Client{Transport: roundTripFunc(func(request *http.Request) (*http.Response, error) {
					return &http.Response{StatusCode: 200}, nil
				})},
				repo: repoMock{storeStatusCodeMock: func(requestUrl string, code int) error {
					return nil
				}},
			},
			args{endpoint: "http://example.com"},
			assert.NoError,
		},
		{
			"HTTP call error",
			fields{
				log: testlogger.New(),
				client: &http.Client{Transport: roundTripFunc(func(request *http.Request) (*http.Response, error) {
					return nil, errors.New("some http-call error")
				})},
				repo: nil,
			},
			args{endpoint: "http://example.com"},
			assert.Error,
		},
		{
			"DB call error",
			fields{
				log: testlogger.New(),
				client: &http.Client{Transport: roundTripFunc(func(request *http.Request) (*http.Response, error) {
					return &http.Response{StatusCode: 200}, nil
				})},
				repo: repoMock{storeStatusCodeMock: func(requestUrl string, code int) error {
					return errors.New("some error from db-call")
				}},
			},
			args{endpoint: "http://example.com"},
			assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				log:    tt.fields.log,
				client: tt.fields.client,
				repo:   tt.fields.repo,
			}
			tt.wantErr(t, s.checkStatus(tt.args.endpoint), fmt.Sprintf("checkStatus(%v)", tt.args.endpoint))
		})
	}
}
