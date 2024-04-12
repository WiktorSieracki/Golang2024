package main

import (
	"math/big"
)

func Factorial(n uint64) *big.Int {
    result := big.NewInt(1)
    for i := uint64(2); i <= n; i++ {
        result.Mul(result, big.NewInt(0).SetUint64(i))
    }
    return result
}

func CountFibonacciExecutions(n int) (result int, amountOfExecutions []int){
	amount_of_executions := make([]int, n+1)
	var helper func(int) int
	helper = func(n int) int {
		amount_of_executions[n] += 1
		if n == 0 {
			return 0
		} else if n == 1 {
			return 1
		} else {
			return helper(n-1) + helper(n-2)
		}
	}
	return helper(n), amount_of_executions
}