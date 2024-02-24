package hw02unpackstring

import (
	"errors"
	"fmt"
	"regexp"
)

var ErrInvalidString = errors.New("invalid string")

// Check format of string data.
//
// A successful check_data returns err == nil.
// Returns err == ErrInvalidString, if data is wrong format.
func check_data(data string) (err error) {
	res, err := regexp.MatchString(`^\d|[^\\]\d\d`, data)
	if err != nil {
		fmt.Println(err)
		return err
	} else if res {
		fmt.Println("invalid string")
		return ErrInvalidString
	}
	return
}

// Repeat repeatable string amount times.
func slice_repetion(repeatable []rune, amount int) []rune {
	var new_slice []rune
	for i := 0; i < amount; i++ {
		new_slice = append(new_slice, repeatable...)
	}
	return new_slice
}

// func get_char(runes []rune, start_idx int, end_idx int) (char []rune) {
// 	var re *regexp.Regexp = regexp.MustCompile(`\\(\d)`)
// 	char = runes[start_idx:end_idx]
// 	res := re.FindStringSubmatch(string(runes))
// 	if len(res) != 0 {
// 		char = []rune(res[1])
// 	}
// 	return
// }

// Format string accoding to given format.
//
// idxs - list of indexes received after parsing string by regexp:
//
//	[start matching, stop matching, start repeatable part, end repeatable part, start repetition amount,
//	end repetition amount]
//
// end_idx - index of symbol, where last formatting was finished
// runes_str - formattable string
//
// returning:
// []rune - formatted string
// int - new end index
func format_string(idxs []int, formatted_str []rune, end_idx int, runes_str []rune) ([]rune, int) {
	formatted_str = append(formatted_str, runes_str[end_idx:idxs[0]]...)
	var new_end_idx int = idxs[1]
	char := runes_str[idxs[2]:idxs[3]]
	repetition_amount := int(runes_str[idxs[4]] - '0')
	var repetition []rune = slice_repetion(char, int(repetition_amount))
	formatted_str = append(formatted_str, repetition...)
	return formatted_str, new_end_idx
}

// unpack string,  e.g. a4gh3 => aaaaghhh
//
// data - string for unpacking
//
// returning:
// string - unpacked string
// error - ErrInvalidString
func Unpack(data string) (string, error) {
	if err := check_data(data); err != nil {
		return "", err
	}
	var re *regexp.Regexp = regexp.MustCompile(`([a-zA-Z\s]|\\[^$])(\d)`)
	var group_idxs [][]int = re.FindAllStringSubmatchIndex(data, -1)
	if len(group_idxs) == 0 {
		return data, nil
	}
	runes := []rune(data)
	formated_str := []rune{}
	end_inx := 0
	for _, idxs := range group_idxs {
		formated_str, end_inx = format_string(idxs, formated_str, end_inx, runes)
	}
	formated_str = append(formated_str, runes[end_inx:len(data)+1]...)
	return string(formated_str), nil
}
