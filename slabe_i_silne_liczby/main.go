package main

import (
	"fmt"
	"math"
	"time"
)



func main() {
	nick := PrepareNick("WikSie")
	// nick := PrepareNick("pioarl")

	ascii_array := CreateAsciiNickArray(nick)
	fmt.Println(ascii_array)


	// -------------------------------------------------------
	StrongNumber := FindStrongNumber(nick)
	fmt.Printf("Strong number: %d\n", StrongNumber)

	// -------------------------------------------------------

	WeakNumber := FindWeakNumber(nick)
	fmt.Printf("Weak number: %d\n", WeakNumber)

	// -------------------------------------------------------

	start := time.Now()
	_ = Fibonacci(42)

	end := time.Since(start)
	fmt.Printf("Execution time for Fibonacci(42): %v\n", end)

	// -------------------------------------------------------
	// 2^257 seconds in years
	fmt.Println(math.Pow(2,257)/31556952)

}
