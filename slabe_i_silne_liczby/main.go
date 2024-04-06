package main

import (
	"fmt"
	"math/big"
	"strings"
)


func factorial(n uint64) *big.Int {
    result := big.NewInt(1)
    for i := uint64(2); i <= n; i++ {
        result.Mul(result, big.NewInt(0).SetUint64(i))
    }
    return result
}

func fib(n int) int {
	if n == 0 {
		return 0
	}else if n == 1 {
		return 1
	}else {
		return fib(n-1)+fib(n-2)
	}
}

func findStrongNumber(nick string) int {
	var ascii_nick_array [6]int
	for i, v := range nick {
		ascii_nick_array[i] = int(v)
	}
	var count uint64 = 1
	var running = true
	for running {
		number_of_contains := 0
		BigNumber := factorial(count)
		for _, v := range ascii_nick_array {
			if strings.Contains(BigNumber.Text(10), fmt.Sprint(v)) {
				number_of_contains += 1
			}
		}
		if number_of_contains == len(nick) {
			running = false
		}
		count += 1
	}
	return int(count)

}

func main() {
    nick := "WikSie"
	StrongNumber := findStrongNumber(nick)
	fmt.Println(StrongNumber)

	// -------------------------------------------------------

	WeakNumber := fib(30)
	fmt.Println(WeakNumber)
}
