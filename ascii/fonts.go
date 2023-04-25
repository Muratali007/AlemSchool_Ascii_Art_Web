package ascii

import (
	"fmt"
	"strings"
)

func ReadFont(res string) (map[rune][]string, error) {
	res = strings.ReplaceAll(res, "\r", "")

	standardList := strings.Split(string(res), "\n")

	s := make(map[rune][]string)

	var startPoint rune = ' '
	var temp []string

	for i, v := range standardList {
		i++
		if i%9 == 0 {
			temp = append(temp, v)
			s[startPoint] = temp
			startPoint++
			temp = nil
		} else if v != "" {
			temp = append(temp, v)
		}
	}

	return s, nil
}

func ConvertedText(text string, format map[rune][]string) string {
	if len(text) < 1 {
		return ""
	}

	res := ""
	text = strings.ReplaceAll(text, "\r", "")
	line := strings.Split(strings.ReplaceAll(text, "\n", "\\n"), "\\n")

	for j, word := range line {
		if len(word) < 1 {
			if j != len(line)-1 {
				res += "\n"
			}
			continue
		} else {
			for i := 0; i < 8; i++ {
				for _, char := range word {
					res += fmt.Sprint(format[char][i])
				}
				res += "\n"
			}
		}
	}

	return res
}
