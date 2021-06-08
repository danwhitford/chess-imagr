package main

import (
	"fmt"
	"flag"
	"path/filepath"
	"os"
)

func main() {
	glob := flag.String("glob", "*.pgn", "Glob for input files")
	output := flag.String("output", "out", "Directory to output images")
	flag.Parse()

	matching, _ := filepath.Glob(*glob)
	for _, v := range matching {
		fmt.Println(v)
	}

	os.Mkdir(*output, 0755)
}
