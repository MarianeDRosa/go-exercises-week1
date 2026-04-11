package main

import (
	"fmt"
	"sort"
	"strings"
)

// WordStat stores the word and its frequency
type WordStat struct {
	Word  string
	Count int
}

func CountWords(text string) []WordStat {
	counts := make(map[string]int)
	words := strings.Fields(strings.ToLower(text))

	for _, word := range words {
		counts[word]++
	}

	// 1. Convert Map to Slice of Structs
	var stats []WordStat
	for word, count := range counts {
		stats = append(stats, WordStat{Word: word, Count: count})
	}

	// 2. Sort the Slice by Count (Descending - Highest first)
	sort.Slice(stats, func(i, j int) bool {
		if stats[i].Count == stats[j].Count {
			return stats[i].Word < stats[j].Word // Alphabetical if count is same
		}
		return stats[i].Count > stats[j].Count
	})

	return stats
}

func main() {
	text := "Go is fast and Go is fun"
	results := CountWords(text)

	fmt.Println("Word Frequency (Sorted by Count):")
	for _, stat := range results {
		fmt.Printf("%s: %d\n", stat.Word, stat.Count)
	}
}
