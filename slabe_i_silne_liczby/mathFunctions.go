package main

import (
	"math/big"
	"time"
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


func Counter() int {
	n := 0
	for {
		start := time.Now()
		_ = Fibonacci(n)
		end := time.Since(start)
		if end.Seconds() > 1 {
			return n
		}
		n += 1
	}
}

func CounterAsync() int {
    n := 0
    done := make(chan bool)
    go func() {
        for {
            select {
            case <-done:
                return
            default:
                _ = Fibonacci(n)
                n++
            }
        }
    }()
    time.Sleep(time.Second)
    done <- true
    return n - 1
}