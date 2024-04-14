package main

import (
	"math"
	"strconv"
	"strings"
)

func CreateAsciiNickArray(nick string) (asciiNickArray []int) {
    for _, char := range nick {
        asciiNickArray = append(asciiNickArray, int(char))
    }
    return asciiNickArray
}

func FindStrongNumber(nick string)  (strongNumber int) {
	var ascii_nick_array = CreateAsciiNickArray(nick)
	for {
		number_of_contains := 0
		BigNumber := Factorial(strongNumber)
		for _, v := range ascii_nick_array {
			if strings.Contains(BigNumber.Text(10), strconv.Itoa(v)) {
				number_of_contains += 1
			}
		}
		if number_of_contains >= len(nick) {
			break
		}
		strongNumber += 1
	}
	return strongNumber
}

func FindWeakNumber(nick string) int {
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