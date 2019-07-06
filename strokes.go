package script

import (
	"math/rand"

	"github.com/fogleman/gg"
)

// Stroke is a single element of a glyph
type Stroke struct {
	Type   string
	Points []Point
	Render func(image gg.Context)
}

func (glyph Glyph) randomStroke() Stroke {
	var newPoint Point
	points := []Point{}

	newStroke := Stroke{}

	strokeTypes := []string{
		"point",
		"line",
	}

	strokeType := strokeTypes[rand.Intn(len(strokeTypes))]

	if strokeType == "point" {
		newPoint = randomPoint(glyph.Width, glyph.Height)
		points = append(points, newPoint)

		newStroke = Stroke{
			Type:   "point",
			Points: points,
			Render: func(image gg.Context) {
				image.DrawPoint(newPoint.x, newPoint.y, 8)
				image.Stroke()
			},
		}
	} else if strokeType == "line" {
		newPoint = randomPoint(glyph.Width, glyph.Height)
		points = append(points, newPoint)
		newStroke = Stroke{
			Type:   "point",
			Points: points,
			Render: func(image gg.Context) {
				image.DrawPoint(newPoint.x, newPoint.y, 8)
				image.Stroke()
			},
		}
	}

	return newStroke
}
