package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

var maxWordsCount = 10

func Top10(incomingText string) []string {
	if len(incomingText) == 0 {
		return make([]string, 0)
	}

	explodedWords := getMapOfCountedWords(incomingText)
	sortedWords := getSliceOfSortedWords(explodedWords)

	return sortedWords[:getMaxWordsCount(sortedWords)]
}

func getMapOfCountedWords(incomingText string) map[string]int {
	words := strings.Split(incomingText, " ")
	countedWords := make(map[string]int)

	for _, currentWord := range words {
		if currentWord == "" {
			continue
		}

		for _, explodedWords := range strings.Fields(currentWord) {
			countedWords[explodedWords]++
		}
	}

	return countedWords
}

func getSliceOfSortedWords(countedWords map[string]int) []string {
	sortedWords := make([]string, 0, len(countedWords))
	for word := range countedWords {
		sortedWords = append(sortedWords, word)
	}

	sort.Slice(sortedWords, func(i, j int) bool {
		if countedWords[sortedWords[i]] == countedWords[sortedWords[j]] {
			return sortedWords[i] < sortedWords[j]
		}

		return countedWords[sortedWords[i]] > countedWords[sortedWords[j]]
	})

	return sortedWords
}

func getMaxWordsCount(words []string) int {
	wordsCount := len(words)
	if wordsCount < maxWordsCount {
		maxWordsCount = wordsCount
	}

	return maxWordsCount
}
