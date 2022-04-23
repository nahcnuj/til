package reflection

import (
	"reflect"
	"testing"
)

type Person struct {
	Name    string
	Profile Profile
}
type Profile struct {
	Age  int
	City string
}

func TestWalk(t *testing.T) {
	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"struct with one string field",
			struct {
				Name string
			}{"Chris"},
			[]string{"Chris"},
		},
		{
			"struct with two string fields",
			struct {
				Name string
				City string
			}{"Tom", "London"},
			[]string{"Tom", "London"},
		},
		{
			"struct with non string field",
			struct {
				Name string
				Age  int
			}{"Nao", 30},
			[]string{"Nao"},
		},
		{
			"nested field",
			Person{
				"Junichi",
				Profile{26, "Kashiwa"},
			},
			[]string{"Junichi", "Kashiwa"},
		},
		{
			"pointers to struct",
			&Person{
				"Junichi",
				Profile{26, "Kashiwa"},
			},
			[]string{"Junichi", "Kashiwa"},
		},
		{
			"slices",
			[]Profile{
				{26, "Kashiwa"},
				{33, "London"},
			},
			[]string{"Kashiwa", "London"},
		},
		{
			"arrays",
			[2]Profile{
				{26, "Kashiwa"},
				{33, "London"},
			},
			[]string{"Kashiwa", "London"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			walk(test.Input, func(input string) {
				got = append(got, input)
			})

			want := test.ExpectedCalls
			if !reflect.DeepEqual(got, want) {
				t.Errorf("got %v, want %v", got, want)
			}
		})
	}
}
