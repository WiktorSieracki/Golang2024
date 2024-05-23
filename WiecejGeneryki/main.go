package main

import (
	"fmt"
)

type FractionType[T types] struct{
	numerator, denominator T
}

type types interface {
	int | int16 | int32 | int64  | uint | uint16 | uint32 | uint64
}

func gcd[T types](a, b T) T {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func (f FractionType[T]) simplify() FractionType[T] {
    g := gcd(f.numerator, f.denominator)
    return FractionType[T]{f.numerator / g, f.denominator / g}
}

func (f1 FractionType[T]) multiply(f2 FractionType[T]) FractionType[T] {
	result := FractionType[T]{f1.numerator * f2.numerator, f1.denominator * f2.denominator}
	return result.simplify()
}

func (f1 FractionType[T]) add(f2 FractionType[T]) FractionType[T] {
	result := FractionType[T]{f1.numerator * f2.denominator + f2.numerator * f1.denominator, f1.denominator * f2.denominator}
	return result.simplify()
}

func (f1 FractionType[T]) subtract(f2 FractionType[T]) FractionType[T] {
	result := FractionType[T]{f1.numerator * f2.denominator - f2.numerator * f1.denominator, f1.denominator * f2.denominator}
	return result.simplify()
}

func main() {
	fmt.Println("Hello, World!")
	f1 := FractionType[int]{2, 3}
	f2 := FractionType[int]{3, 4}
	fmt.Println(f1.multiply(f2))
	fmt.Println(f1.add(f2))
	fmt.Println(f1.subtract(f2).add(f2))
}