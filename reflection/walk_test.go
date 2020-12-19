package reflection

import (
	"reflect"
	"testing"
)

type Address struct {
	address string
}

func TestWalk(t *testing.T) { // Apply function on all strings
	cases := []struct {
		Name     string
		Input    interface{}
		Expected []string
	}{
		{
			"One string field",
			struct {
				Name string
			}{"Raj"},
			[]string{"Raj"},
		},
		{
			"Multiple strings",
			struct {
				Name     string
				Lastname string
				Age      int
			}{"Raj", "Saxena", 31},
			[]string{"Raj", "Saxena"},
		},
		{
			"Nested struct",
			struct {
				Name    string
				Address Address
			}{"Raj", Address{"Berlin"}},
			[]string{"Raj", "Berlin"},
		},
		{
			"Pointers",
			&struct {
				Name    string
				Address *Address
			}{"Raj", &Address{"Berlin"}},
			[]string{"Raj", "Berlin"},
		},
		{
			"Slices",
			[]Address{
				{"Berlin"},
				{"Delhi"},
			},
			[]string{"Berlin", "Delhi"},
		},
		{
			"Array",
			[2]Address{
				{"Berlin"},
				{"Delhi"},
			},
			[]string{"Berlin", "Delhi"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string

			Walk(test.Input, func(i string) {
				got = append(got, i)
			})

			if !reflect.DeepEqual(got, test.Expected) {
				t.Errorf("got %+v, expected %+v", got, test.Expected)
			}
		})
	}
}
