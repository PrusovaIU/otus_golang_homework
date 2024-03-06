package hw03frequencyanalysis

import (
	"regexp"
	"sort"
	"strings"
)

// func mapSort(sorted_map map[string]int) map[string]int {
// 	var keys []string
// 	for key := range sorted_map {
// 		keys = append(keys, key)
// 	}
// 	sort.Strings(keys)
// 	var result = make(map[string]int, len(keys))
// 	for _, key := range keys {
// 		result[key] = sorted_map[key]
// 	}
// 	return result
// }

func keys(_map map[int][]string) []int {
	var keys []int = make([]int, 0, len(_map))

	for key := range _map {
		keys = append(keys, key)
	}
	return keys
}

// func mapSort(sortable_map map[string]int) []string {
// 	var map_keys []string = keys(sortable_map)

// 	sort.SliceStable(map_keys, func(i, j int) bool {
// 		return sortable_map[map_keys[i]] > sortable_map[map_keys[j]]
// 	})

// 	return map_keys
// }

func frequency_range(frequency map[string]int) map[int][]string {
	frequency_ranging := make(map[int][]string)
	for key, value := range frequency {
		words, ok := frequency_ranging[value]
		if ok {
			words = append(words, key)

		} else {
			words = []string{key}
		}
		frequency_ranging[value] = words
	}
	return frequency_ranging
}

func top(frequency_range map[int][]string) []string {
	top_len := 10
	var top []string

	var map_keys []int = keys(frequency_range)
	sort.Sort(sort.Reverse(sort.IntSlice(map_keys)))

	for i := 0; len(top) < top_len && i < len(map_keys); i++ {
		var words []string = frequency_range[map_keys[i]]
		sort.Strings(words)
		if free_places_count := top_len - len(top); free_places_count < len(words) {
			top = append(top, words[0:free_places_count]...)
		} else {
			top = append(top, words...)
		}
	}
	return top
}

func Top10(data string) []string {
	var re *regexp.Regexp = regexp.MustCompile(`[.,!?":;\()+=\s] ?`)
	var words []string = re.Split(data, -1)
	frequency := make(map[string]int, len(words))
	for _, word := range words {
		if word != `-` && word != `` {
			word := strings.ToLower(word)
			frequency[word] += 1
		}
	}
	var frequency_range map[int][]string = frequency_range(frequency)

	return top(frequency_range)
}
