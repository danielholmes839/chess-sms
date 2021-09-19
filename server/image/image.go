package image

import (
	"bytes"
	"fmt"
	"htn/server/game"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"

	"github.com/notnil/chess"
	"github.com/notnil/chess/image"
)

func Generate(puzzle *game.Puzzle) error {
	// Generate image
	games, err := chess.GamesFromPGN(strings.NewReader(pgn))
	if err != nil {
		return err
	}

	board := games[0].Position().Board()

	// Write an SVG to a buffer
	buf := bytes.Buffer{}
	if err := image.SVG(&buf, board); err != nil {
		fmt.Println(err)
	}

	// Save the file
	svg := buf.Bytes()
	fileName := fmt.Sprintf("%d.png", puzzleId)
	save(svg, fileName)
	return nil
}

// Use rsvg to render svg (convert bytes from svg to png)
func save(svg []byte, name string) {
	path := "./data/"
	inputFilename := path + name + ".svg"
	outputFilename := path + name + ".png"

	// write the .svg file
	if err := ioutil.WriteFile(inputFilename, svg, 0600); err != nil {
		panic(err)
	}

	// convert the .svg file to the output format (perhaps png)
	if err := exec.Command("rsvg-convert", inputFilename, "-b", "white", "-f", "png", "-o", outputFilename).Run(); err != nil {
		panic(err)
	}

	// Remove temporary svg
	if err := os.Remove(inputFilename); err != nil {
		panic(err)
	}

	fmt.Println("Finished creating image")
}
