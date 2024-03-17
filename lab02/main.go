package main

import "fmt"


func collatz(liczba int)int{
	counter:=0
	//copy:= liczba
	for liczba != 1 {
		if (liczba % 2 == 0){
			liczba /= 2
			}else{
				liczba = (liczba*3)+1
			}
			fmt.Printf("%v ",liczba)
			counter+=1
		}
		//fmt.Printf("Ilość operacji dla %9d: %d \n",copy,counter)
		return counter
}

func srednia(dolny int, gorny int){
	wszystkie:=0
	for i := dolny; i <= gorny; i++ {
		wszystkie+=collatz(i)
	}
	fmt.Printf("Średnia liczb od %d do %d: %d\n",dolny,gorny,wszystkie/(gorny-dolny))
}

func min_max(dolny int, gorny int){
	max:=0
	min:=collatz(dolny)
	var max_liczba int
	var min_liczba int
	for i := dolny; i <= gorny; i++ {
		wyniczek:= collatz(i)
		if wyniczek<=min{
			min=wyniczek
			min_liczba=i
		}
		if wyniczek>max{
			max=wyniczek
			max_liczba=i
		}
	}
	fmt.Printf("Zakres liczb od %d do %d\n",dolny,gorny)
	fmt.Printf("   --Minimalna liczba operacji dla liczby %d: %d\n   --Maksymalna liczba operacji %d: %d\n",min_liczba,min,max_liczba,max)
}

func main() {
	//collatz(123)
	//collatz(1234)
	//collatz(12345)
	//collatz(123456)
	//collatz(1234567)
	//collatz(12345678)
	collatz(1234567891011)
	// srednia(1,1000)
	// srednia(1000,2000)
	// min_max(1,1000)
	// min_max(1000,2000)



}