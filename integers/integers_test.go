package integers

import (
	"fmt"
	"testing"
)

func TestAddition(t *testing.T) {
	t.Run("Should return sum", func(t *testing.T) {
		got := Add(1, 2)
		want := 3

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}
