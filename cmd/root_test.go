package cmd

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRootCMD(t *testing.T) {
	t.Run("Test ping", func(t *testing.T) {
		req := httptest.NewRequest("GET", "http://localhost:9090/ping", nil)
		rr := httptest.NewRecorder()
		HandlePing(rr, req)
		if status := rr.Code; status != http.StatusOK {
			t.Errorf("Wanted %v but got %v", http.StatusOK, status)
		}
		expected := "pong"
		got := rr.Body.String()
		if got != expected {
			t.Errorf("Wanted %q in body but got %q", expected, got)
		}
	})

	t.Run("Test healtcheck", func(t *testing.T) {
		req := httptest.NewRequest("GET", "http://localhost:9090/healtcheck", nil)
		rr := httptest.NewRecorder()
		HandleHealthCheck(rr, req)
		if status := rr.Code; status != http.StatusOK {
			t.Errorf("Wanted %v but got %v", http.StatusOK, status)
		}
	})

	t.Run("Test run action", func(t *testing.T) {})
}
