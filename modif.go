package main

import (
	"bufio"
	"strings"
)

func Process(text string) string {
	var resultLines []string
	scanner := bufio.NewScanner(strings.NewReader(text))

	for scanner.Scan() {
		line := scanner.Text()
		if strings.TrimSpace(line) == "" {
			resultLines = append(resultLines, "")
			continue
		}

		words := strings.Fields(line)
		if len(words) == 0 {
			resultLines = append(resultLines, "")
			continue
		}

		for i := 0; i < len(words); i++ {
			if words[i] == "(up)" && i > 0 {
				words[i-1] = strings.ToUpper(words[i-1])
				words = append(words[:i], words[i+1:]...)
				i--
				continue
			}
			if words[i] == "(low)" && i > 0 {
				words[i-1] = strings.ToLower(words[i-1])
				words = append(words[:i], words[i+1:]...)
				i--
				continue
			}
			if words[i] == "(cap)" && i > 0 {
				words[i-1] = capitalize(words[i-1])
				words = append(words[:i], words[i+1:]...)
				i--
				continue
			}
		}

		resultLines = append(resultLines, strings.Join(words, " "))
	}

	return strings.Join(resultLines, "\n")
}

func capitalize(s string) string {
	if len(s) == 0 {
		return s
	}
	runes := []rune(s)
	runes[0] = []rune(strings.ToUpper(string(runes[0])))[0]
	return string(runes)
}
