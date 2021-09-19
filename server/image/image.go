package image

import (
	"bytes"
	"fmt"
	"htn/server/game"
	"io/ioutil"
	"os"
	"os/exec"

	"github.com/notnil/chess/image"
)

// Generate image
func Generate(puzzle *game.Puzzle) error {
	board := puzzle.GetPosition().Board()

	// Write an SVG to a buffer
	buf := bytes.Buffer{}
	if err := image.SVG(&buf, board); err != nil {
		fmt.Println(err)
	}

	// Save the file
	svg := buf.Bytes()
	save(svg, puzzle.GetID())
	return nil
}

// Use rsvg to render svg (convert bytes from svg to png)
func save(svg []byte, id int) {
	path := "./server/image/data/"
	idString := fmt.Sprintf("%d", id)
	inputFilename := path + idString + ".svg"
	outputFilename := path + idString + ".png"

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
}
