package mocking

import (
	"bytes"
	"testing"
)

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

func TestMocking(t *testing.T) {
	b := &bytes.Buffer{}
	s := &SpySleeper{}
	Countdown(b, s)

	got := b.String()
	want := `3
2
1
Go!`

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}

	if s.Calls != 3 {
		t.Errorf("Sleeper count insufficient. want %d, got %d", 3, s.Calls)
	}
}
