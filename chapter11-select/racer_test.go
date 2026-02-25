package racer

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {

	t.Run("normal test between two servers", func(t *testing.T) {
		slowServer := makeDelayedTestServer(30)
		fastServer := makeDelayedTestServer(0 * time.Millisecond)

		defer slowServer.Close()
		defer fastServer.Close()

		want := fastServer.URL
		got, _ := Racer(slowServer.URL, fastServer.URL)

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("timeout test if no response in 10s", func(t *testing.T) {
		server := makeDelayedTestServer(50 * time.Millisecond)

		defer server.Close()
		defer server.Close()

		_, err := ConfigurableRacer(server.URL, server.URL, 25*time.Millisecond)

		if err == nil {
			t.Errorf("expected an error but didn't get one")
		}
	})

}

func makeDelayedTestServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}
