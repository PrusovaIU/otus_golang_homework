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

func keys(_map map[string]int) []string {
	var keys []string = make([]string, 0, len(_map))

	for key := range _map {
		keys = append(keys, key)
	}
	return keys
}

func mapSort(sortable_map map[string]int) []string {
	var map_keys []string = keys(sortable_map)

	sort.SliceStable(map_keys, func(i, j int) bool {
		return sortable_map[map_keys[i]] > sortable_map[map_keys[j]]
	})

	return map_keys
}

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

func Top10(data string) []string {
	var re *regexp.Regexp = regexp.MustCompile(`[.,!?":;\()+=\s] ?`)
	var words []string = re.Split(data, -1)
	top := make(map[string]int, len(words))
	for _, word := range words {
		if word != `-` {
			word := strings.ToLower(word)
			top[word] += 1
		}
	}
	_top := mapSort(top)
	print(_top)
	return words
}
