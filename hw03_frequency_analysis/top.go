package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

func Top10(text string) []string {
	wordsFreq := make(map[string]int)

	fields := strings.Fields(text)
	for _, word := range fields {
		wordsFreq[word]++
	}

	words := make([]string, 0, len(wordsFreq))
	for k := range wordsFreq {
		words = append(words, k)
	}

	sort.Strings(words)
	sort.SliceStable(words, func(i, j int) bool {
		return wordsFreq[words[i]] > wordsFreq[words[j]]
	})

	resultLength := 10
	if len(words) < 10 {
		resultLength = len(words)
	}

	return words[:resultLength]
}
