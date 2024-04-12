package main

import (
	"fmt"
)



func main() {
	// nick := PrepareNick("WikSie")
	nick := PrepareNick("pioarl")
	StrongNumber := FindStrongNumber(nick)
	fmt.Println(StrongNumber)

	// -------------------------------------------------------

	WeakNumber := FindWeakNumber(nick)
	fmt.Println(WeakNumber)
}
