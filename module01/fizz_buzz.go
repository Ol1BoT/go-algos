package module01

import (
	"fmt"
	"strconv"
	"strings"
)

// FizzBuzz will print out all of the numbers
// from 1 to N replacing any divisible by 3
// with "Fizz", and divisible by 5 with "Buzz",
// and any divisible by both with "Fizz Buzz".
//
// Note: The test for this is a little
// complicated so that you can just use the
// `fmt` package and print to standard out.
// I wouldn't normally recommend this, but did
// it here to make life easier for beginners.
func FizzBuzz(n int) {
	if n == 0 {
		return
	}

	var res strings.Builder

	for i := 1; i <= n; i++ {
		if i%5 == 0 && i%3 == 0 {
			res.WriteString(" Fizz Buzz,")
			continue
		}
		if i%3 == 0 {
			res.WriteString(" Fizz,")
			continue
		}
		if i%5 == 0 {
			res.WriteString(" Buzz,")
			continue
		}

		res.WriteString(" " + strconv.Itoa(i) + ",")
	}

	final := res.String()[1 : res.Len()-1]

	fmt.Println(final)

}
