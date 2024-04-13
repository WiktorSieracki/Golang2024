
<div align="center">

## Silne i Słabe Liczby
### Autor  
#### [Wiktor Sieracki](https://github.com/WiktorSieracki)

</div>

### Wstęp

---
Program do zabawy większymi liczbami w GO.  
Program wylicza 2 liczby na podstawie 3 liter imienia i 3 liter nazwiska.
Mój "nick" to *WikSie*

## Początek
Żeby zacząć operować na "nicku" trzeba zamienić wszystkie litery na ich małe odpowiedniki. Następnie usunąć polskie znaki jeżeli takie występują.  
Służy do tego funkcja `PrepareNick()`

```go
func PrepareNick(nick string) (LoweredAndRemovedPolishChars string) {
        nick = replacePolishChars(nick)
        nick = strings.ToLower(nick)
        return nick
    }
```
> Korzystamy z funkcji `replacePolishChars()` która zamienia polskie znaki na ich odpowiedniki.
```go
func replacePolishChars(input string) string {
    polishChars := []string{"ą", "ć", "ę", "ł", "ń", "ó", "ś", "ź", "ż", "Ą", "Ć", "Ę", "Ł", "Ń", "Ó", "Ś", "Ź", "Ż"}
    englishChars := []string{"a", "c", "e", "l", "n", "o", "s", "z", "z", "A", "C", "E", "L", "N", "O", "S", "Z", "Z"}
    for i, char := range polishChars {
        input = strings.ReplaceAll(input, char, englishChars[i])
    }
    return input
}

```


---
### Liczenie wartości 297 liczby ciągu Fibonacciego

Funkcja Fibonacci(n int) zwraca n-tą liczbę ciągu Fibonacciego.


```go
func Fibonacci(n int) int {
    if n <= 1 {
        return n
    }
    return Fibonacci(n-1) + Fibonacci(n-2)
}
```

Jej złożoność obliczeniowa wynosi O(2^n) co sprawia teoretycznie, że każde kolejne wywołanie funkcji jest dwa razy wolniejsze od poprzedniego.

Czas który przebiega na obliczenie operacji w GO można zmierzyć za pomocą funkcji `time.Now()` oraz `time.Since()`. Importujemy do tego bibliotekę **time**.

```go
package main

import (
	"fmt"
	"time"
)

start := time.Now()
fmt.Println(Fibonacci(40)) // 102334155
end := time.Since(start)
fmt.Println(end) // 959.5387ms
```

> Zakładając, że funkcja Fibonacci(40) trwa 1s na naszym urządzeniu  
> Chcąc obliczyć wartość Fibbonacci(297) czyli dla mojej silnej liczby  
> musielibyśmy czekać $2^{257}$ razy dłużej czyli około $2^{257}$ sekund  
> zakładając, że rok ma 31,556,952 sekund to $7*10^{69}$ lat.  
> Porównując to do wieku wszechświata który wynosi około 13,8  miliarda czyli $13.8*10^9$ lat

---

### Rzeczy których się nauczyłem podczas pisania programu
 - importowanie funkcji z innych plików (go run . a nie go run main.go)
 - strings to bardzo przydatna biblioteka
 - pomiar czasu w GO