package main

import (
	"fmt"
	"flag"
	"path/filepath"
	"os"
	"github.com/notnil/chess"
)

func main() {
	glob := flag.String("glob", "*.pgn", "Glob for input files")
	output := flag.String("output", "out", "Directory to output images")
	flag.Parse()

	matching, _ := filepath.Glob(*glob)
	for _, v := range matching {
		fmt.Println(v)
		pgnReader, err := os.Open(v)
		if err != nil {
			panic(err)
		}
		pgn, err := chess.PGN(pgnReader)
		if err != nil {
			panic(err)
		}
		game := chess.NewGame(pgn)
		for _, pos := range game.Positions() {
			fmt.Println(pos.Board().Draw())
		}
	}

	os.Mkdir(*output, 0755)
}
