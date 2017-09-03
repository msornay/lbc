package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

// fizzBuzz returns a slice of strings where all integers in [1, n+1] appear in order.
// In that slice :
//   * Multiple of a and b are replaced by sa+sb
//   * Multiple of a are replaced by sa
//   * Multiple of b are replaced by sb
func fizzBuzz(sa, sb string, a, b, n int) ([]string, error) {
	if a == 0 || b == 0 {
		return nil, errors.New("multiples cannot be 0")
	}
	if n < 0 {
		return nil, fmt.Errorf("invalid limit: %d", n)
	}
	res := make([]string, n)
	for i := 0; i < n; i++ {
		k := i + 1
		switch {
		case k%(a*b) == 0:
			res[i] = sa + sb
		case k%a == 0:
			res[i] = sa
		case k%b == 0:
			res[i] = sb
		default:
			res[i] = strconv.Itoa(k)
		}
	}
	return res, nil
}

// http request json body
type fizzBuzzParams struct {
	String1 string `json:"string1"`
	String2 string `json:"string2"`
	Int1    int    `json:"int1"`
	Int2    int    `json:"int2"`
	Limit   int    `json:"limit"`
}

// Limit the size of responses to limit*(len(string1)+len(string2)) < maxChar
const maxChar = 1048576

func FizzBuzzHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)

	if r.Body == nil {
		http.Error(w, "empty JSON body", 400)
		return
	}
	var p fizzBuzzParams
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, "cannot decode JSON body content", 400)
		return
	}

	if p.Limit*(len(p.String1)+len(p.String2)) > maxChar {
		log.Println(maxChar)
		http.Error(w, "FizzBuzz list too big", 400)
		return
	}

	fb, err := fizzBuzz(p.String1, p.String2, p.Int1, p.Int2, p.Limit)
	if err != nil {
		http.Error(w, "cannot generate fizzbuzz list", 400)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(fb)
	if err != nil {
		http.Error(w, "cannot encode response", 500)
		return
	}
}

func main() {
	var addr string
	flag.StringVar(&addr, "addr", ":8080", "TCP network address to listen to")

	flag.Parse()

	http.HandleFunc("/", FizzBuzzHandler)

	log.Printf("FizzBuzz server listening on %s\n", addr)
	http.ListenAndServe(addr, nil)
}
