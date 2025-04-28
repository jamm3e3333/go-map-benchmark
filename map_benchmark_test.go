package main

import (
	"bufio"
	"os"
	"strings"
	"testing"
)

func loadKingLearWords() []string {
	file, err := os.Open("kinglear.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	var words []string
	var word strings.Builder

	for {
		r, _, err := reader.ReadRune()
		if err != nil {
			break
		}

		if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9') {
			word.WriteRune(r)
		} else if word.Len() > 0 {
			words = append(words, strings.ToLower(word.String()))
			word.Reset()
		}
	}

	return words
}

func BenchmarkMapInsert(b *testing.B) {
	words := loadKingLearWords()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		wordToCount := make(map[string]int)
		for _, word := range words {
			wordToCount[word]++
		}
	}
}

func BenchmarkMapLookup(b *testing.B) {
	wordToCount := make(map[string]int)
	words := loadKingLearWords()
	for _, word := range words {
		wordToCount[word]++
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, word := range words {
			_ = wordToCount[word]
		}
	}
}

func BenchmarkMapUpdate(b *testing.B) {
	words := loadKingLearWords()
	wordToCount := make(map[string]int)
	for _, word := range words {
		wordToCount[word]++
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, word := range words {
			wordToCount[word] = len(word)
		}
	}
}

func BenchmarkMapDelete(b *testing.B) {
	words := loadKingLearWords()
	wordToCount := make(map[string]int)
	for _, word := range words {
		wordToCount[word]++
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, k := range words {
			delete(wordToCount, k)
		}
	}
}
