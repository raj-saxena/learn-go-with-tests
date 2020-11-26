package dependencyInjection

import (
	"bytes"
	"testing"
)

func TestDependencyInjection(t *testing.T) {
	t.Run("Greet", func(t *testing.T) {
		b := bytes.Buffer{}

		Greet(&b, "Raj")

		got := b.String()
		want := "Hello, Raj"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}
