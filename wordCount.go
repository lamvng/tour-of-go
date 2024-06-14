package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	words := strings.Fields(s)
	counter := make(map[string]int)
	for _, word := range words {
		if counter[word] == 0 {
			counter[word] = 1
		} else {
			counter[word]++
		}
	}
	return counter
}

func main() {
	wc.Test(WordCount)
}
