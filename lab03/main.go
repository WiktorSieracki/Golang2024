package main

import (
	"flag"
	"fmt"
	"os"
)

func collatz(liczba int){
	fmt.Printf("%v ",liczba)
	for liczba != 1 {
		if (liczba % 2 == 0){
			liczba /= 2
			}else{
				liczba = (liczba*3)+1
			}
			fmt.Printf("%v ",liczba)
		}
		fmt.Printf("\n")
}

func count_numbers(text string){
	tab:= []int{0,0,0,0,0,0,0,0,0,0}
	for _,char := range text{
		if char >= '0' && char <= '9'{
			tab[char-'0']++
		}
	}
	liczby :=[]int{0,1,2,3,4,5,6,7,8,9}
	for i,_ :=range liczby{
		fmt.Printf("%v %v \n",liczby[i],tab[i])
	}
}

func main(){
	var collatz_count = flag.Int("N",0,"Ilość liczb collatza")
	flag.Parse()
	if *collatz_count != 0 {
		for i := 1; i <= *collatz_count; i++ {
			collatz(i)
		}
	}else{
		dane,err:=os.ReadFile("./dane.txt")
		if err!=nil{
			panic(err)
		}
		count_numbers(string(dane))
	}
}