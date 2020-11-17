package structsAndInterface

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f, want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {

	checkArea := func(t *testing.T, s Shape, want float64) {
		t.Helper()
		got := s.Area()
		if got != want {
			t.Errorf("got %g, want %g", got, want)
		}
	}

	t.Run("rectangle", func(t *testing.T) {
		rectangle := Rectangle{10.0, 10.0}
		want := 100.0

		checkArea(t, rectangle, want)
	})

	t.Run("circles", func(t *testing.T) {
		circle := Circle{10.0}
		want := 314.1592653589793
		checkArea(t, circle, want)
	})

	// table tests
	areaTests := []struct {
		s Shape
		w float64
	}{
		{Rectangle{Width: 12, Height: 6}, 72.0},
		{Circle{Radius: 10}, 314.1592653589793},
		{Triangle{Height: 12, Base: 6}, 36.0},
	}

	for _, tt := range areaTests {
		checkArea(t, tt.s, tt.w)
	}
}
