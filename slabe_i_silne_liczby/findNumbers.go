package main

import (
	"fmt"
	"math"
	"strings"
)

func FindStrongNumber(nick string) int {
	var ascii_nick_array = make([]int, len(nick))
	for i, v := range nick {
		ascii_nick_array[i] = int(v)
	}
	fmt.Println(ascii_nick_array)
	var count uint64 = 0
	for {
		number_of_contains := 0
		BigNumber := Factorial(count)
		for _, v := range ascii_nick_array {
			if strings.Contains(BigNumber.Text(10), fmt.Sprint(v)) {
				number_of_contains += 1
			}
		}
		if number_of_contains >= len(nick) {
			break
		}
		count += 1
	}
	return int(count)
}

func FindWeakNumber(nick string) int {
	nick = strings.ToLower(nick)
	StrongNumber := FindStrongNumber(nick)
	_, amount_of_executions := CountFibonacciExecutions(30)
	var absolute_executions_minus_strong [31]float64
	for i, v := range amount_of_executions {
		absolute_executions_minus_strong[i] = math.Abs(float64(v - StrongNumber))
	}
	min := math.MaxFloat64
	var index int
	for i, v := range absolute_executions_minus_strong {
		if v < min {
			min = v
			index = i
		}
	}
	return index
}