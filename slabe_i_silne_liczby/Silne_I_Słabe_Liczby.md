
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
> Przydatna jest też funkcja `strings.ReplaceAll()` która zamienia wszystkie wystąpienia danego znaku na inny znak.
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
Po tych operacjach nasz nick *WikSie* zamieni się na **wiksie** i nim będziemy operować na dalszych etapach.  
Gdy mamy już przygotowany nick możemy obliczyć silną liczbę z naszego nicku.


## Czym jest silna liczba?

> Silna liczba jest to liczba, której silnia zawiera w sobie wszystkie wartości numeryczne kodów ASCII z nicku.

Funkcja `CreateAsciiNickArray()` tworzy tablicę z wartościami numerycznymi kodów ASCII z nicku.

```go
func CreateAsciiNickArray(nick string) (asciiNick []int) {
    for _, char := range nick {
        asciiNick = append(asciiNick, int(char))
    }
    return asciiNick
}
```

Dla nicku *wiksie*  tablica wygląda następująco:  
`[119 105 107 115 105 101]`

Następną przeszkodą jest obliczenie takiej liczby, której silnia zawiera w sobie wszystkie wartości numeryczne kodów ASCII z nicku.

Dla przypomnienia wzór na silnię:
$$n! = n * (n-1) * (n-2) * ... * 1$$


## Obliczanie silnej liczby

Musimy być przygotowani na to, że taka liczba może być bardzo duża.

Do operacji na dużych liczbach w GO używamy biblioteki `math/big`.

```go
import "math/big"
```

Podstawowymi funkcjami z tej biblioteki są `NewInt()` oraz `Mul()`.  
`NewInt()` tworzy nową zmienną typu `big.Int` a `Mul()` mnoży dwie liczby typu `big.Int`.

Funkcja Factorial() przystosowana do operacji na dużych liczbach wygląda następująco:

```go
func Factorial(n int) *big.Int {
    result := big.NewInt(1)
    for i := 2; i <= n; i++ {
        result.Mul(result, big.NewInt(int64(i)))
    }
    return result
}
```
> warto zauważyć, że funkcja Factorial() zwraca **wskaźnik** na zmienną typu `big.Int`



Znając już wartości numeryczne kodów ASCII z nicku oraz funkcję obliczającą silnię możemy obliczyć silną liczbę.

Mając tablicę wyglądającą tak:  
`[119 105 107 115 105 101]`  
Potrzebujemy znaleźć taką liczbę, której silnia zawiera w sobie wszystkie wartości z tej tablicy, czyli dla przykładu liczbę:  
2`105`7042`119`9060333`101`91488`107`301914`115`

```go
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
```

1. Funkcja ta tworzy tablicę z wartościami numerycznymi kodów ASCII z nicku.
2. Następnie w nieskończonej pętli oblicza silnię kolejnych liczb aż do momentu, w którym silnia zawiera w sobie wszystkie wartości numeryczne kodów ASCII z nicku.
    - Tworzy zmienną `BigNumber` typu `big.Int` i przypisuje jej wartość silni.
    - Następnie sprawdza czy silnia zawiera w sobie wszystkie wartości numeryczne kodów ASCII z nicku. (`number_of_contains >= len(nick)`)
    - Jeżeli nie to zwiększa silną liczbę (strongNumber) o 1 i przechodzi do kolejnej iteracji.
    - Jeżeli tak to kończy działanie pętli.
3. Funkcja zwraca wartość silnej liczby.

> W funkcji tej używamy funkcji `strings.Contains()` która zwraca wartość logiczną czy dany string zawiera w sobie inny string.  
> Do zamiany wartości z `big.Int` na string używamy metody `Text(10)`. Gdzie 10 oznacza system dziesiętny w jakim chcemy otrzymać wartość.  
> Do zamiany wartości numerycznej na string używamy funkcji `strconv.Itoa()`. Alternatywnie można użyć `fmt.Sprint()`, ale jest ona wolniejsza, bo musi sprawdzać typ zmiennej.

Dla nicku *wiksie* silna liczba wynosi **297**.

Mając już silną liczbę możemy obliczyć słabą liczbę.

## Czym jest słaba liczba?

> Słaba liczba to argument dla którego liczba wywołań funkcji podczas obliczania 30-tego elementu ciągu Fibonacciego jest najbliższa wartości silnej liczby.

Zmodyfikowana funkcja zwraca wartość liczby ciągu Fibonacciego oraz ilość wywołań funkcji.

```go
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
```

1. Funkcja ta tworzy tablicę `amount_of_executions` w której zapisuje ilość wywołań funkcji dla danego elementu ciągu Fibonacciego.
2. Następnie rekurencyjnie oblicza wartość n-tego elementu ciągu Fibonacciego. (funkcja `helper()`)
3. Funkcja zwraca wartość n-tego elementu ciągu Fibonacciego oraz tablicę `amount_of_executions`.

Ta funkcja będzie nam potrzebna do obliczenia słabej liczby.

## Obliczanie słabej liczby

```go
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
```
1. Funkcja ta oblicza silną liczbę z nicku. [FindStrongNumber()](#obliczanie-silnej-liczby)
2. Następnie oblicza ilość wywołań funkcji dla 30-tego elementu ciągu Fibonacciego. [CountFibonacciExecutions()](#czym-jest-słaba-liczba)
3. Tworzy tablicę `absolute_executions_minus_strong` w której zapisuje różnicę między ilością wywołań funkcji a wartością silnej liczby.
4. Następnie znajduje wartość minimalną z tablicy `absolute_executions_minus_strong` i zwraca indeks tej wartości.

Dla nicku *wiksie* słaba liczba wynosi **17**.


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

Jej złożoność obliczeniowa wynosi O(2^n) co sprawia teoretycznie, że jeżeli n wzrośnie o 1 funkcja będzie wykonywać się dwa razy dłużej.

### Pomiar czasu w GO

Czas który przebiega na obliczenie operacji w GO można zmierzyć za pomocą funkcji `time.Now()` oraz `time.Since()`. Importujemy do tego bibliotekę ***time***.

```go
import (
	"fmt"
	"time"
)

start := time.Now()
fmt.Println(Fibonacci(42)) // 267914296
end := time.Since(start)
fmt.Println(end) // 959.5387ms
```

> Zakładając, że funkcja Fibonacci(40) trwa 1s na naszym urządzeniu  
> Chcąc obliczyć wartość Fibbonacci(297) czyli dla mojej silnej liczby  
> musielibyśmy czekać $2^{257}$ razy dłużej czyli około $2^{257}$ sekund  
> zakładając, że rok ma 31,556,952 sekund to $7*10^{69}$ lat.  
> Porównując to do wieku wszechświata który wynosi około 13,8  miliarda czyli $13.8*10^9$ lat
> to jest to około $5*10^{59}$ razy dłużej niż wiek wszechświata.

Widać więc, że obliczenie wartości 297 liczby ciągu Fibonacciego jest niemożliwe w sensownym czasie.
Bardzo ciężko jest sobie wyobrazić jak długo by to trwało.

---

### Rzeczy których się nauczyłem podczas pisania programu
 - operacje na dużych liczbach w GO ([math/big](https://pkg.go.dev/math/big))
 - importowanie funkcji z innych plików (go run . a nie go run main.go)
 - strings to bardzo przydatna biblioteka ([strings](https://pkg.go.dev/strings))
 - pomiar czasu w GO ([time](https://pkg.go.dev/time))
 - przypomnienie złożoności obliczeniowej (O(2^n) dla funkcji Fibonacci)