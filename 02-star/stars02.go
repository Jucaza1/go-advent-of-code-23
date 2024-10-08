package main

import "strings"

func replaceParse(l []byte) []byte {
	var list = [20]string{
		"one", "1",
		"two", "2",
		"three", "3",
		"four", "4",
		"five", "5",
		"six", "6",
		"seven", "7",
		"eight", "8",
		"nine", "9",
		"zero", "0",
	}
	str := string(l)
	replacer := strings.NewReplacer(list[:]...)
	result := []byte(replacer.Replace(str))
	return result
}

func processLine(l []byte) uint8 {
	str := string(replaceParse(l))
	var first uint8
	var last uint8
	for i := 0; i < len(str); i++ {
		if str[i] >= '0' && str[i] <= '9' {
			first = uint8(str[i] - '0')
			break
		}
	}
	for i := len(str) - 1; i >= 0; i-- {
		if str[i] >= '0' && str[i] <= '9' {
			last = uint8(str[i] - '0')
			break
		}
	}
	return first*10 + last
}
