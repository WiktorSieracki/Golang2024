package main

import (
	"fmt"
	"math"
)

func main(){
	var a float64
	var b float64
	var c float64
	fmt.Println("Program do obliczania trójmianu kwadratowego ax^2+bx+c = 0")
	fmt.Println("Podaj a")
	fmt.Scan(&a)
	fmt.Println("Podaj b")
	fmt.Scan(&b)
	fmt.Println("Podaj c")
	fmt.Scan(&c)
	
	delta:= math.Pow(b,2)-4*a*c
	sqrt_delta:= math.Sqrt(delta)
	x1:= (-1*b-sqrt_delta)/(2*a)
	x2:= (-1*b+sqrt_delta)/(2*a)

	fmt.Printf("a: %f b: %f c:%f \n",a,b,c)
	fmt.Println("Delta:",delta)
	if delta < 0 {
		fmt.Println("Nie ma miejsc zerowych")
	}else if delta == 0 {
		fmt.Println("x:",x1)
	}else{
		fmt.Println("x1:",x1,"x2:",x2)
	}
	dwa:= "DWA"
	fmt.Printf("JEDEN, %s, TRZY",dwa)
}