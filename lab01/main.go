package main

import "fmt"

func reverse_array(s []int) []int {
    length := len(s)
    reversed := make([]int, length) // tworzenie array z int o określonym limicie
    for i := range s {
        reversed[i] = s[length-i-1]
    }
    return reversed
}

func reverse_string(s string) string {
    runes := []rune(s)
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    return string(runes)
}

func main() {
    array := []int{1, 2, 3, 4, 5}
	my_string := "Hello"
    fmt.Println(reverse_array(array))  // Output: [5 4 3 2 1]
	fmt.Println(reverse_string(my_string)) // Output: olleH
}