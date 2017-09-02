package main

import (
	"fmt"
	"testing"
)

func ExampleFizzBuzz() {
	fmt.Println(FizzBuzz("fizz", "buzz", 3, 5, 16))
	// Output:  [1 2 fizz 4 buzz fizz 7 8 fizz buzz 11 fizz 13 14 fizzbuzz 16] <nil>
}

func TestFizzBuzz(t *testing.T) {
	// negative limit
	_, err := FizzBuzz("f", "b", 3, 5, -1)
	if err == nil {
		t.Errorf("expected non-nil error for negative limit")
	}

	cases := []struct {
		sa, sb  string
		a, b, n int
		out     []string
	}{
		{
			"f", "b", 1, 2, 3,
			[]string{"f", "fb", "f"},
		}, {
			"f", "b", 3, 5, 0,
			[]string{},
		},
	}
	for _, c := range cases {
		res, _ := FizzBuzz(c.sa, c.sb, c.a, c.b, c.n)
		if len(res) != len(c.out) {
			t.Errorf("fizzBuzz: unexpected output length")
		}
		for i, s := range res {
			if s != c.out[i] {
				t.Errorf("fizzBuzz: got %s want %s", s, c.out)
			}
		}
	}
}
