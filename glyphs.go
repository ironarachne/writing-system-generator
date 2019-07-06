package script

import (
	"math/rand"

	"github.com/fogleman/gg"
)

// Glyph represents a single glyph
type Glyph struct {
	Representation string
	Strokes        []Stroke
	Width          int
	Height         int
	GridSizeX      int
	GridSizeY      int
}

func generateGlyph(representation string, width int, height int) Glyph {
	var newStroke Stroke

	numberOfStrokes := rand.Intn(3) + 1

	strokes := []Stroke{}

	glyph := Glyph{
		Representation: representation,
		Width:          width,
		Height:         height,
		GridSizeX:      5,
		GridSizeY:      5,
	}

	for i := 0; i < numberOfStrokes; i++ {
		newStroke = glyph.randomStroke()
		strokes = append(strokes, newStroke)
	}

	glyph.Strokes = strokes

	return glyph
}

// Render draws and saves a glyph
func (glyph Glyph) Render() {
	image := gg.NewContext(glyph.Width, glyph.Height)

	for _, s := range glyph.Strokes {
		s.Render(*image)
	}

	fileName := "./output/" + glyph.Representation + ".png"

	image.SavePNG(fileName)
}
