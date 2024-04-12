package main

import "strings"

func ReplacePolishChars(input string) string {
	polishToEnglish := map[rune]rune{
		'ą': 'a', 'ć': 'c', 'ę': 'e', 'ł': 'l', 'ń': 'n', 'ó': 'o', 'ś': 's', 'ź': 'z', 'ż': 'z',
		'Ą': 'A', 'Ć': 'C', 'Ę': 'E', 'Ł': 'L', 'Ń': 'N', 'Ó': 'O', 'Ś': 'S', 'Ź': 'Z', 'Ż': 'Z',
	}

	output := make([]rune, len(input))
	for i, char := range input {
		if replacement, ok := polishToEnglish[char]; ok {
			output[i] = replacement
		} else {
			output[i] = char
		}
	}

	return string(output)
}

func PrepareNick(nick string) (LoweredAndRemovedPolishChars string) {
	nick = ReplacePolishChars(nick)
	nick = strings.ToLower(nick)
	return nick
}