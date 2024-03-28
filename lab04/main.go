package main

import (
	"flag"
	"fmt"
)


func dobry_pesel(pesel string)bool  {
	pesel_chars := []rune(pesel)
	sprawdzacz :=[...]int{1,3,7,9,1,3,7,9,1,3}
	dodane:=0
	for i := range sprawdzacz{
		pesel_razy_sprawdzacz:=sprawdzacz[i]*int(pesel_chars[i]-'0')
		dodane+=(pesel_razy_sprawdzacz%10)
	}
	wynik:=10-(dodane%10)


	return wynik==int(pesel_chars[len(pesel_chars)-1]-'0')
}


func main(){
	var pesel string
    flag.StringVar(&pesel, "p", "", "pesel")
    flag.Parse()
    fmt.Println(dobry_pesel(pesel))
	// fmt.Println(dobry_pesel("11028537491"))
	// fmt.Println(dobry_pesel("98989743999"))
	// fmt.Println(dobry_pesel("10110101110"))
	// fmt.Println(dobry_pesel("90090515836"))
	// fmt.Println(dobry_pesel("02070803628"))

}