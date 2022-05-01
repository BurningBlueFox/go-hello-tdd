package website_racer

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	slowServer := makeDelayedServer(10 * time.Millisecond)
	fastServer := makeDelayedServer(0 * time.Millisecond)
	defer slowServer.Close()
	defer fastServer.Close()

	slowUrl := slowServer.URL
	fastUrl := fastServer.URL

	want := fastUrl
	got := Racer(slowUrl, fastUrl)

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}

}

func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		time.Sleep(delay)
		writer.WriteHeader(http.StatusOK)
	}))
}
