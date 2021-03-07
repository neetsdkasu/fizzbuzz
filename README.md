# FizzBuzz

Library of FizzBuzz  

### example
```go
import (
	"fmt"
	. "github.com/neetsdkasu/fizzbuzz"
)

func main() {
	var i uint8
	for i = 10; i <= 17; i++ {
		fmt.Println(FizzBuzz(i))
	}
	// Output:
	// Buzz
	// 11
	// Fizz
	// 13
	// 14
	// FizzBuzz
	// 16
	// 17
}
```
