package racer

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {

	t.Run("find faster", func(t *testing.T) {
		fastServer := makeDelayedServer(10 * time.Millisecond)
		slowServer := makeDelayedServer(50 * time.Millisecond)

		defer slowServer.Close()
		defer fastServer.Close()

		slowURL := slowServer.URL
		fastURL := fastServer.URL

		want := fastURL

		start := time.Now()
		got, err := Racer(slowURL, fastURL)

		if err != nil {
			t.Fatalf("Error not expected. got: %v", err)
		}

		fmt.Printf("run time %s \n", time.Since(start)) // The method returns quickly but the test takes time as the slow server needs to response before closing.

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("return error if takes longer than 3 sec", func(t *testing.T) {
		fastServer := makeDelayedServer(4 * time.Second)
		slowServer := makeDelayedServer(5 * time.Second)

		defer slowServer.Close()
		defer fastServer.Close()

		_, err := Racer(slowServer.URL, fastServer.URL)

		if err == nil {
			t.Errorf("expected an error")
		}
	})
}

func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}
