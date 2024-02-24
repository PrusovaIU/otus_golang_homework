package main

import (
	"fmt"

	"github.com/PrusovaIU/otus_golang_homework/hw02_unpack_string/hw02unpackstring"
)

func main() {
	data := [...]string{"a4bc2d5e", "abcd", "aaa0b", "d\n5abc", "qwe\\45", "qwe\\4\\5", "qwe\\\\5"}
	for i, el := range data {
		formattedStr, _ := hw02unpackstring.Unpack(el)
		fmt.Println(i, el, "=>", formattedStr)
	}
	fmt.Println("---\nErrors")
	errData := [...]string{"3abc", "45", "aaa10b"}
	for i, el := range errData {
		_, err := hw02unpackstring.Unpack(el)
		fmt.Println(i, el, "=>", err)
	}
}
