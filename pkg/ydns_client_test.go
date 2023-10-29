package pkg

import (
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
