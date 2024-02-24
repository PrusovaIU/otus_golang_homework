package main

import (
	"fmt"

	"github.com/fixme_my_friend/hw02_unpack_string/hw02unpackstring"
)

func main() {
	var data = [...]string{"a4bc2d5e", "abcd", "aaa0b", "d\n5abc", "qwe\\45", "qwe\\4\\5"}
	// var data = [...]string{"a4bc2d5e"}
	for i, el := range data {
		formatted_str, _ := hw02unpackstring.Unpack(el)
		fmt.Println(i, el, "=>", formatted_str)
	}
}
