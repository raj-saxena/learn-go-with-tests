package racer

import (
	"fmt"
	"net/http"
	"time"
)

func Racer(a, b string) (string, error) {
	select {
	case <-measureTime(a):
		return a, nil
	case <-measureTime(b):
		return b, nil
	case <-time.After(3 * time.Second):
		return "", fmt.Errorf("Too late, timed out")
	}
}

func measureTime(url string) chan struct{} {
	ch := make(chan struct{})
	go func() {
		http.Get(url)
		close(ch)
	}()
	return ch
}
