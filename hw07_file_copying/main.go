package main

import (
	"flag"
	"fmt"
)

var (
	from, to      string
	limit, offset int64
)

func init() {
	flag.StringVar(&from, "from", "C:\\Users\\pruso\\Downloads\\from_file.txt", "file to read from")
	flag.StringVar(&to, "to", "C:\\Users\\pruso\\Downloads\\to_file.txt", "file to write to")
	flag.Int64Var(&limit, "limit", 10, "limit of bytes to copy")
	flag.Int64Var(&offset, "offset", 0, "offset in input file")
}

func main() {
	flag.Parse()
	fmt.Println(from, to)
	err := Copy(from, to, offset, limit)
	if err != nil {
		fmt.Println(err)
	}
}
