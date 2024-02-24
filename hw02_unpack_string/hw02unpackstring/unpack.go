package hw02unpackstring

import (
	"bytes"
	"errors"
	"regexp"
	"strconv"
)

var ErrInvalidString = errors.New("invalid string")

const backslachRegexp = `\\(.)`

// Check format of string data.
//
// A successful check_data returns err == nil.
// Returns err == ErrInvalidString, if data is wrong format.
func check_data(data string) (err error) {
	res, err := regexp.MatchString(`^\d|[^\\]\d\d`, data)
	if err != nil {
		return err
	} else if res {
		return ErrInvalidString
	}
	return
}

// Repetition symbols by regexp
func repetition(s []byte) []byte {
	var re *regexp.Regexp = regexp.MustCompile(`(\s|.+)(\d)`)
	var groups []string = re.FindStringSubmatch(string(s))
	var char []byte = []byte(groups[1])
	amount, _ := strconv.Atoi(groups[2])
	var result []byte = bytes.Repeat(char, amount)
	return result
}

// Replace backslash to ""
func replace_backslash(s []byte) []byte {
	var re *regexp.Regexp = regexp.MustCompile(backslachRegexp)
	var groups []string = re.FindStringSubmatch(string(s))
	return []byte(groups[1])
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
	var formatted = re.ReplaceAllFunc([]byte(data), repetition)
	re = regexp.MustCompile(backslachRegexp)
	formatted = re.ReplaceAllFunc(formatted, replace_backslash)
	return string(formatted), nil
}
