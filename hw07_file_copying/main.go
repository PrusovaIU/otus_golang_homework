package main

import (
	"flag"
	"fmt"

	"github.com/PrusovaIU/otus_golang_homework/hw07_file_copying/dd_copy"
)

var (
	from, to      string
	limit, offset int64
)

func init() {
	flag.StringVar(&from, "from", "", "file to read from")
	flag.StringVar(&to, "to", "", "file to write to")
	flag.Int64Var(&limit, "limit", 0, "limit of bytes to copy")
	flag.Int64Var(&offset, "offset", 0, "offset in input file")
}

func main() {
	flag.Parse()
	fmt.Println(from, to)
	err := dd_copy.Copy(from, to, offset, limit)
	if err != nil {
		fmt.Println(err)
	}
}
