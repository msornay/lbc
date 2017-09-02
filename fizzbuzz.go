package fizzbuzz

import (
	"fmt"
	"strconv"
)

// FizzBuzz returns a slice of strings where all integers in [1, n+1] appear in order.
// In that slice :
// * Multiple of a and b are replaced by sa+sb
// * Multiple of a are replaced by sa
// * Multiple of b are replaced by sb
func FizzBuzz(sa, sb string, a, b, n int) ([]string, error) {
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
