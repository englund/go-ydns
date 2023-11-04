package pkg

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

const username = "username"
const password = "password"

func TestGetIp(t *testing.T) {
	testCases := []struct {
		name        string
		ip          string
		statusCode  int
		expectError bool
	}{
		{
			name:        "SuccessfulRequest",
			ip:          "13.37.73.31",
			statusCode:  http.StatusOK,
			expectError: false,
		},
		{
			name:        "ServerError",
			ip:          "",
			statusCode:  http.StatusInternalServerError,
			expectError: true,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				require.Equal(t, "/ip", r.URL.Path)
				w.WriteHeader(tc.statusCode)
				w.Write([]byte(tc.ip))
			}))
			defer server.Close()

			client := NewYdnsClient(server.URL, username, password)
			actualIp, err := client.GetIp()
			if tc.expectError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.ip, *actualIp)
			}
		})
	}
}

func TestUpdate(t *testing.T) {
	testCases := []struct {
		name        string
		host        string
		ip          string
		response    string
		expectError bool
		statusCode  int
	}{
		{
			name:        "ValidResponseBody",
			host:        "test",
			ip:          "13.37.73.31",
			response:    "ok",
			expectError: false,
			statusCode:  http.StatusOK,
		},
		{
			name:        "InvalidResponseBody",
			host:        "test",
			ip:          "13.37.73.31",
			response:    "notok",
			expectError: true,
			statusCode:  http.StatusOK,
		},
		{
			name:        "InvalidStatusCode",
			host:        "test",
			ip:          "13.37.73.31",
			response:    "",
			expectError: true,
			statusCode:  http.StatusUnauthorized,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				require.Equal(t, "/update/", r.URL.Path)
				require.Equal(t, fmt.Sprintf("host=%s&ip=%s", tc.host, tc.ip), r.URL.RawQuery)

				actualUsername, actualPassword, ok := r.BasicAuth()
				require.True(t, ok)

				require.Equal(t, username, actualUsername)
				require.Equal(t, password, actualPassword)

				w.WriteHeader(tc.statusCode)
				w.Write([]byte(tc.response))
			}))
			defer server.Close()

			client := NewYdnsClient(server.URL, username, password)
			err := client.Update(tc.host, tc.ip)
			if tc.expectError {
				require.Error(t, err, "Expected error, got nil")
			} else {
				require.NoError(t, err, "Did not expect error, got: %s", err)
			}
		})
	}
}
