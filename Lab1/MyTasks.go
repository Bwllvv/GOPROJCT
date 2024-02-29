package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(wordCount("Hello my name is lev"))
	fmt.Println(wordsCountMap("Hello<Bob!!"))
}

func wordCount(s string) int {

	words := 0
	splitString := strings.Fields(s)
	for i := 0; i < len(splitString); i++ {
		words++
	}
	return words
}

func wordsCountMap(s string) map[string]int {

	fields := strings.Fields(s)
	wordsMap := make(map[string]int)
	for _, value := range fields {

		word := strings.ToLower(
			strings.ReplaceAll(
				strings.ReplaceAll(
					strings.ReplaceAll(
						strings.ReplaceAll(
							value, ",", ""), "?", ""), "!", ""), ".", ""))

		wordsMap[word] += 1
	}
	return wordsMap
}
