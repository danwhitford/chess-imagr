package main

import (
	"flag"
	"fmt"
	"github.com/notnil/chess"
	"github.com/notnil/chess/image"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func fileNameWithoutExtTrimSuffix(fileName string) string {
	return strings.TrimSuffix(path.Base(fileName), filepath.Ext(fileName))
}

func writeImage(fname string, position *chess.Position) {
	f, err := os.Create(fname)
	if err != nil {
		log.Panic(err)
	}
	defer f.Close()

	if err := image.SVG(f, position.Board()); err != nil {
		log.Panic(err)
	}
}

func main() {

	glob := flag.String("glob", "*.pgn", "Glob for input files")
	output := flag.String("output", "out", "Directory to output images")
	flag.Parse()

	os.Mkdir(*output, 0755)
	matching, _ := filepath.Glob(*glob)
	log.Printf("Found matching input files: %s", matching)
	for _, v := range matching {
		log.Printf("Parsing input file: %s", v)
		pgnReader, err := os.Open(v)
		if err != nil {
			log.Panic(err)
		}
		pgn, err := chess.PGN(pgnReader)
		if err != nil {
			log.Panic(err)
		}
		game := chess.NewGame(pgn)
		for i, pos := range game.Positions() {
			fname := fmt.Sprintf("%s/%s_%d.svg", *output, fileNameWithoutExtTrimSuffix(v), i)
			log.Printf("Writing move %d of %s to %s", i, v, fname)
			writeImage(fname, pos)
		}
	}
}
