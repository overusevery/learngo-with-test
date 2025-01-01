package playerhandler

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPlayerHandler(t *testing.T) {
	t.Run("returns Pepper's score", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/api/players/Pepper", nil)
		response := httptest.NewRecorder()

		NewPlayerHandler().ServeHTTP(response, request)

		got := response.Body.String()
		want := "20"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
	t.Run("returns Floyd's score", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/api/players/Floyd", nil)
		response := httptest.NewRecorder()

		NewPlayerHandler().ServeHTTP(response, request)

		got := response.Body.String()
		want := "10"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}
