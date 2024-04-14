package main

import (
	"math/big"
)

func Factorial(n int) *big.Int {
    result := big.NewInt(1)
    for i := 2; i <= n; i++ {
        result.Mul(result, big.NewInt(int64(i)))
    }
    return result
}

func CountFibonacciExecutions(n int) (result int, amountOfExecutions []int){
	amount_of_executions := make([]int, n+1)
	var helper func(int) int
	helper = func(n int) int {
		amount_of_executions[n] += 1
		if n <= 1 {
			return n
		}
		return helper(n-1) + helper(n-2)
	}
	return helper(n), amount_of_executions
}

func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}