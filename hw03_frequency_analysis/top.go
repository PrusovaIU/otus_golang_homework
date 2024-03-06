package hw03frequencyanalysis

import (
	"regexp"
	"sort"
	"strings"
)

func keys(_map map[int][]string) []int {
	keys := make([]int, 0, len(_map))

	for key := range _map {
		keys = append(keys, key)
	}
	return keys
}

func frequencyRange(frequency map[string]int) map[int][]string {
	frequencyRanging := make(map[int][]string)
	for key, value := range frequency {
		words, ok := frequencyRanging[value]
		if ok {
			words = append(words, key)

		} else {
			words = []string{key}
		}
		frequencyRanging[value] = words
	}
	return frequencyRanging
}

func top(frequencyRange map[int][]string) []string {
	topLen := 10
	var top []string

	mapKeys := keys(frequencyRange)
	sort.Sort(sort.Reverse(sort.IntSlice(mapKeys)))

	for i := 0; len(top) < topLen && i < len(mapKeys); i++ {
		words := frequencyRange[mapKeys[i]]
		sort.Strings(words)
		if FreePlacesCount := topLen - len(top); FreePlacesCount < len(words) {
			top = append(top, words[0:FreePlacesCount]...)
		} else {
			top = append(top, words...)
		}
	}
	return top
}

func Top10(data string) []string {
	re := regexp.MustCompile(`[.,!?":;\()+=\s] ?`)
	words := re.Split(data, -1)
	frequency := make(map[string]int, len(words))
	for _, word := range words {
		if word != `-` && word != `` {
			word := strings.ToLower(word)
			frequency[word]++
		}
	}
	frequencyRange := frequencyRange(frequency)

	return top(frequencyRange)
}
