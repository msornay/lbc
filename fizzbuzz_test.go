package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func ExampleFizzBuzz() {
	fmt.Println(fizzBuzz("fizz", "buzz", 3, 5, 16))
	// Output:  [1 2 fizz 4 buzz fizz 7 8 fizz buzz 11 fizz 13 14 fizzbuzz 16] <nil>
}

func TestFizzBuzz(t *testing.T) {
	// negative limit
	_, err := fizzBuzz("f", "b", 3, 5, -1)
	if err == nil {
		t.Errorf("expected non-nil error for negative limit")
	}

	// a = 0
	_, err = fizzBuzz("f", "b", 0, 5, 16)
	if err == nil {
		t.Errorf("expected non-nil error zero multiples")
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
		res, _ := fizzBuzz(c.sa, c.sb, c.a, c.b, c.n)
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

func TestFizzBuzzHandler(t *testing.T) {

	cases := []struct {
		ReqBody  io.Reader
		Status   int
		Expected []string
	}{
		{nil, 400, nil},
		{bytes.NewBufferString("{"), 400, nil},
		{
			bytes.NewBufferString(
				`{"string1":"f", "string2":"b", "int1": 3, "int2": 5, "limit": 16}`),
			200, []string{"1", "2", "f", "4", "b", "f", "7", "8", "f", "b",
				"11", "f", "13", "14", "fb", "16"},
		}, {
			bytes.NewBufferString(
				`{"string1":"f", "string2":"b", "int1": 3, "int2": 5, "limit": 1048576}`),
			400, nil,
		},
	}

	for _, c := range cases {
		req, err := http.NewRequest("GET", "/", c.ReqBody)
		if err != nil {
			t.Fatal(err)
		}
		rec := httptest.NewRecorder()
		handler := http.HandlerFunc(FizzBuzzHandler)
		handler.ServeHTTP(rec, req)
		if rec.Code != c.Status {
			t.Errorf("Expected HTTP status %d, got %d", c.Status, rec.Code)
			return
		}

		// only check expected body for HTTP 200
		if rec.Code == http.StatusOK {
			var res []string
			err := json.NewDecoder(rec.Body).Decode(&res)
			if err != nil {
				t.Error("cannot decode response")
				return
			}
			if len(res) != len(c.Expected) {
				t.Errorf("Unexpected result: %+v", res)
				return
			}
			for i, s := range res {
				if s != c.Expected[i] {
					t.Errorf("Unexpected result: %+v", res)
				}
			}
		}
	}
}
