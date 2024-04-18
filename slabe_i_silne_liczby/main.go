package main

import (
	"fmt"
	"math"
	"time"
)



func main() {
	nick:= GetNick()
	fmt.Println("Nick: ", nick)
	// -------------------------------------------------------

	nick = PrepareNick(nick)

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
	fmt.Println("Finding the highest Fibonacci number in 1 seccond: ", Counter())

	// -------------------------------------------------------
	// 2^257 seconds in years
	fmt.Println("2^257 seconds in years: ",math.Pow(2,257)/31556952)

	// -------------------------------------------------------
}
