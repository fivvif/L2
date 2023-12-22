package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	var arr []string

	for {
		var input string
		fmt.Print("Введите элемент (или введите 'exit' для завершения): ")
		_, err := fmt.Scan(&input)
		if err != nil {
			fmt.Println("Ошибка ввода элемента:", err)
			return
		}

		if input == "exit" {
			break
		}

		arr = append(arr, input)
	}
	anagramSets := searchAnagramsInDict(arr)

	for key, value := range anagramSets {
		fmt.Printf("key : %s => %v \n", key, value)
	}

}
func searchAnagramsInDict(arr []string) map[string][]string {
	anagramDict := make(map[string][]string)
	anagramSets := make(map[string][]string)
	for _, word := range arr {
		word = strings.ToLower(word)
		sortedWord := sortString(word)
		anagramDict[sortedWord] = append(anagramDict[sortedWord], word)
	}

	for _, words := range anagramDict {
		if len(words) > 1 {
			first := words[0]
			sort.Strings(words)
			anagramSets[first] = words
		}

	}
	return anagramSets
}

func sortString(s string) string {
	slice := strings.Split(s, "")
	sort.Strings(slice)
	return strings.Join(slice, "")
}
