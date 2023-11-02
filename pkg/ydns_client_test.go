package pkg

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
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
		t.Run(tc.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if r.URL.Path != "/ip" {
					t.Fatalf("Expected a request to '/ip', got: %s", r.URL.Path)
				}
				w.WriteHeader(tc.statusCode)
				w.Write([]byte(tc.ip))
			}))
			defer server.Close()

			client := NewYdnsClient(server.URL, username, password)
			actualIp, err := client.GetIp()
			if tc.expectError {
				if err == nil {
					t.Errorf("Expected error, got nil")
				}
			} else {
				if err != nil {
					t.Errorf("Did not expect error, got: %v", err)
				}
				if *actualIp != tc.ip {
					t.Errorf("Actual ip %s was not equal to %s", *actualIp, tc.ip)
				}
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
		t.Run(tc.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if r.URL.Path != "/update/" {
					t.Fatalf("Expected a request to '/update/', got: %s", r.URL.Path)
				}
				if r.URL.RawQuery != fmt.Sprintf("host=%s&ip=%s", tc.host, tc.ip) {
					t.Fatalf("Unexpected query, got: %s", r.URL.RawQuery)
				}
				w.WriteHeader(tc.statusCode)
				w.Write([]byte(tc.response))
			}))
			defer server.Close()

			client := NewYdnsClient(server.URL, username, password)
			err := client.Update(tc.host, tc.ip)
			if tc.expectError && err == nil {
				t.Errorf("Expected error, got nil")
			}
			if !tc.expectError && err != nil {
				t.Errorf("Did not expected error %s", err)
			}
		})
	}
}
