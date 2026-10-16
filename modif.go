package main

import (
	"bufio"
	"strconv"
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
			if words[i] == "(hex)" && i > 0 {
				val, err := strconv.ParseInt(words[i-1], 16, 64)
				if err == nil {
					words[i-1] = strconv.FormatInt(val, 10)
					words = append(words[:i], words[i+1:]...)
					i--
				}
				continue
			}
			if words[i] == "(bin)" && i > 0 {
				val, err := strconv.ParseInt(words[i-1], 2, 64)
				if err == nil {
					words[i-1] = strconv.FormatInt(val, 10)
					words = append(words[:i], words[i+1:]...)
					i--
				}
				continue
			}
			if strings.HasPrefix(words[i], "(up,") || strings.HasPrefix(words[i], "(low,") || strings.HasPrefix(words[i], "(cap,") {
				tag := strings.TrimSuffix(words[i], ")")
				parts := strings.Split(tag, ",")
				if len(parts) == 2 {
					op := strings.TrimPrefix(strings.TrimSpace(parts[0]), "(")
					n, err := strconv.Atoi(strings.TrimSpace(parts[1]))
					if err == nil && n > 0 {
						start := i - n
						if start < 0 {
							start = 0
						}
						for j := start; j < i; j++ {
							if op == "up" {
								words[j] = strings.ToUpper(words[j])
							} else if op == "low" {
								words[j] = strings.ToLower(words[j])
							} else if op == "cap" {
								words[j] = capitalize(words[j])
							}
						}
						words = append(words[:i], words[i+1:]...)
						i--
						continue
					}
				}
			}
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
			if (strings.ToLower(words[i]) == "a" || strings.ToLower(words[i]) == "an") && i+1 < len(words) {
				nextWord := words[i+1]
				if len(nextWord) > 0 {
					firstChar := strings.ToLower(string(nextWord[0]))
					shouldBeAn := firstChar == "a" || firstChar == "e" || firstChar == "i" || firstChar == "o" || firstChar == "u" || firstChar == "h"
					if shouldBeAn {
						if words[i] == "A" {
							words[i] = "An"
						} else {
							words[i] = "an"
						}
					}
				}
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
