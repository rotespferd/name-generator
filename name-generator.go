package main

import (
	"flag"
	"fmt"
)

// define CLI flags
//TODO: add default path for default file
var filepath = flag.String("filepath", "./names.json", "The path to the file with the names")

func main() {
	flag.Parse()

	fmt.Println("Load files from ", *filepath)
}
