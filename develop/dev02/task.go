package main

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	var input string
	_, err := fmt.Scan(&input)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Choose unpack method : \n 1.Simple unpack \n 2.Escape unpack")
	var choose int
	_, err = fmt.Scan(&choose)
	if err != nil {
		log.Fatal(err)
	}

	var output string
	if choose == 1 {
		output = unpackString(input)
	} else if choose == 2 {
		output, err = unpackEscape(input)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		log.Fatal("Enter 1 or 2")
	}

	fmt.Println(input, "=>", output)

}
func unpackEscape(s string) (string, error) {
	var result strings.Builder
	runes := []rune(s)
	escape := false
	globalFlag := false
	for i := 0; i < len(runes); i++ {
		char := runes[i]
		if char == '\\' {
			globalFlag = true
			escape = true
			continue

		}
		result.WriteRune(char)
		if escape {
			if i+1 < len(runes) && unicode.IsNumber(runes[i+1]) {
				count, _ := strconv.Atoi(string(runes[i+1]))
				result.WriteString(strings.Repeat(string(char), count-1))
				i++
				escape = false
			}
		}

	}
	if !globalFlag {
		return "", errors.New("invalid string")
	}
	return result.String(), nil
}

func unpackString(s string) string {
	var result strings.Builder
	runes := []rune(s)

	for i := 0; i < len(runes); i++ {
		char := runes[i]

		if i+1 < len(runes) && unicode.IsNumber(runes[i+1]) {
			count, _ := strconv.Atoi(string(runes[i+1]))
			if count == 0 {
				result.WriteRune(char)
				continue
			}
			result.WriteString(strings.Repeat(string(char), count-1))
			i++
		}
		result.WriteRune(char)
	}
	if isNumeric(result.String()) {
		return ""
	}
	return result.String()
}

func isNumeric(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}

//todo УЗНАТЬ АЛГОРИТМ РАСШИФРОВКИ
