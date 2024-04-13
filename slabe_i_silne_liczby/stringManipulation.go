package main

import "strings"

func replacePolishChars(input string) string {
    polishChars := []string{"ą", "ć", "ę", "ł", "ń", "ó", "ś", "ź", "ż", "Ą", "Ć", "Ę", "Ł", "Ń", "Ó", "Ś", "Ź", "Ż"}
    englishChars := []string{"a", "c", "e", "l", "n", "o", "s", "z", "z", "A", "C", "E", "L", "N", "O", "S", "Z", "Z"}

    for i, char := range polishChars {
        input = strings.ReplaceAll(input, char, englishChars[i])
    }

    return input
}

func PrepareNick(nick string) (LoweredAndRemovedPolishChars string) {
	nick = replacePolishChars(nick)
	nick = strings.ToLower(nick)
	return nick
}