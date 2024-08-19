package main

import (
	"os"
	"slices"
	"testing"
)

type testCase struct {
	fileName string
	want     []Link
}

func TestMain(t *testing.T) {
	cases := []testCase{
		{
			"cases/ex0.html",
			[]Link{
				{
					"/dog",
					"Something in a span Text not in a span Bold text!",
				},
			},
		},
	}

	for _, c := range cases {
		file, err := os.Open(c.fileName)

		if err != nil {
			t.Error(err)
		}

		got, err := Parse(file)

		if err != nil {
			t.Error(err)
		}

		equal := slices.EqualFunc(c.want, got, func(lhs, rhs Link) bool {
			return lhs.Text == rhs.Text && lhs.Href == rhs.Href
		})

		if !equal {
			t.Errorf("Parse(%+v) == %+v, want %+v", c.fileName, c.want, got)
		}
	}

}
