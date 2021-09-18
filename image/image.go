package image

import (
	"bytes"
	"fmt"
	"image/color"
	"io/ioutil"
	"log"
	"os"
	"os/exec"

	"github.com/notnil/chess"
	"github.com/notnil/chess/image"
)

func Create() {
	fmt.Println("Started creating image")
	// create board position
	fenStr := "rnbqkbnr/pppppppp/8/8/3P4/8/PPP1PPPP/RNBQKBNR b KQkq - 0 1"
	pos := &chess.Position{}
	if err := pos.UnmarshalText([]byte(fenStr)); err != nil {
		log.Fatal(err)
	}

	// write board SVG to file
	buf := bytes.Buffer{}

	yellow := color.RGBA{255, 255, 0, 1}
	mark := image.MarkSquares(yellow, chess.D2, chess.D4)
	if err := image.SVG(&buf, pos.Board(), mark); err != nil {
		fmt.Println(err)
	}

	Save(buf.Bytes(), "1")
}

// Use rsvg to render svg (convert bytes from svg to png)
func Save(svg []byte, name string) {
	path := "./image/data/"
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
