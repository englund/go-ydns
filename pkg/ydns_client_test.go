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
	const expectedIp = "13.37.73.31"
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/ip" {
			t.Errorf("Expected a request to '/ip', got: %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(expectedIp))
	}))
	defer server.Close()

	client := NewYdnsClient(server.URL, username, password)
	actualIp, err := client.GetIp()
	if err != nil {
		t.Errorf("Did not expected error %s", err)
	}

	if *actualIp != expectedIp {
		t.Errorf("Actual ip %s was not equal to %s", *actualIp, expectedIp)
	}
}

func TestUpdate(t *testing.T) {
	testCases := []struct {
		name               string
		host               string
		ip                 string
		response           string
		expectError        bool
		expectedStatusCode int
	}{
		{
			name:               "ValidResponseBody",
			host:               "test",
			ip:                 "13.37.73.31",
			response:           "ok",
			expectError:        false,
			expectedStatusCode: http.StatusOK,
		},
		{
			name:               "InvalidResponseBody",
			host:               "test",
			ip:                 "13.37.73.31",
			response:           "notok",
			expectError:        true,
			expectedStatusCode: http.StatusOK,
		},
		{
			name:               "InvalidStatusCode",
			host:               "test",
			ip:                 "13.37.73.31",
			response:           "",
			expectError:        true,
			expectedStatusCode: http.StatusUnauthorized,
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
				w.WriteHeader(tc.expectedStatusCode)
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
