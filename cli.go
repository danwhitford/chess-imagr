package main

import (
	"flag"
	"fmt"
	"github.com/notnil/chess"
	"github.com/notnil/chess/image"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func fileNameWithoutExtTrimSuffix(fileName string) string {
	return strings.TrimSuffix(path.Base(fileName), filepath.Ext(fileName))
}

func writeImage(fname string, position *chess.Position) {
	fmt.Println(fname)
	f, err := os.Create(fname)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// write board SVG to file
	if err := image.SVG(f, position.Board()); err != nil {
		panic(err)
	}
}

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
		for i, pos := range game.Positions() {
			fname := fmt.Sprintf("%s/%s_%d.svg", *output, fileNameWithoutExtTrimSuffix(v), i)
			writeImage(fname, pos)
			// fmt.Println(pos.Board().Draw())
		}
	}

	os.Mkdir(*output, 0755)
}
