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
		}

		resultLines = append(resultLines, strings.Join(words, " "))
	}

	return strings.Join(resultLines, "\n")
}
