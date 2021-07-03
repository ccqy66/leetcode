package competion

import "fmt"

func isSumEqual(firstWord string, secondWord string, targetWord string) bool {
	if convert(firstWord)+convert(secondWord) == convert(targetWord) {
		return true
	}
	return false
}

func convert(str string) int {
	value := 0
	for i := 0; i < len(str); i++ {
		value = value*10 + int(str[i]-'a')
	}
	return value
}

func maxValue(n string, x int) string {
	inertIndex := len(n)
	if n[0] == '-' {
		for i := 1; i < len(n); i++ {
			if int(n[i]-'0') > x {
				inertIndex = i
				break
			}
		}
	} else {
		for i := 0; i < len(n); i++ {
			if int(n[i]-'0') < x {
				inertIndex = i
				break
			}
		}
	}
	return fmt.Sprintf("%s%d%s", n[:inertIndex], x, n[inertIndex:])
}
