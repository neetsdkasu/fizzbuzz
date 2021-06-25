// Library of FizzBuzz
package fizzbuzz

import "fmt"

// Get FizzBuzz value of n
func FizzBuzz(n uint32) (res string) {
	switch n % 15 {
	case 0:
		res = "FizzBuzz"
	case 3, 6, 9, 12:
		res = "Fizz"
	case 5, 10:
		res = "Buzz"
	default:
		res = fmt.Sprint(n)
	}
	return
}
