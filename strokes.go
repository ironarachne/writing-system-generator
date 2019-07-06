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
		//"circle",
		"horizontal line",
		"vertical line",
		"random line",
		"right angle",
		"horizontal wedge",
		"vertical wedge",
	}

	strokeType := strokeTypes[rand.Intn(len(strokeTypes))]

	if strokeType == "circle" {
		newPoint = glyph.getRandomCoordinatePoint()
		points = append(points, newPoint)

		newStroke = Stroke{
			Type:   "circle",
			Points: points,
			Render: func(image gg.Context) {
				image.DrawPoint(points[0].x, points[0].y, float64(glyph.Width/glyph.GridSizeX))
				image.Stroke()
			},
		}
	} else if strokeType == "horizontal line" {
		newPoint = glyph.getRandomCoordinatePoint()
		points = append(points, newPoint)
		newPoint = glyph.getRandomCoordinatePoint()
		newPoint.y = points[0].y
		points = append(points, newPoint)
		newStroke = Stroke{
			Type:   "line",
			Points: points,
			Render: func(image gg.Context) {
				image.DrawLine(points[0].x, points[0].y, points[1].x, points[1].y)
				image.Stroke()
			},
		}
	} else if strokeType == "vertical line" {
		newPoint = glyph.getRandomCoordinatePoint()
		points = append(points, newPoint)
		newPoint = glyph.getRandomCoordinatePoint()
		newPoint.x = points[0].x
		points = append(points, newPoint)
		newStroke = Stroke{
			Type:   "line",
			Points: points,
			Render: func(image gg.Context) {
				image.DrawLine(points[0].x, points[0].y, points[1].x, points[1].y)
				image.Stroke()
			},
		}
	} else if strokeType == "random line" {
		newPoint = glyph.getRandomCoordinatePoint()
		points = append(points, newPoint)
		newPoint = glyph.getRandomDifferentCoordinatePoint(points[0])
		points = append(points, newPoint)
		newStroke = Stroke{
			Type:   "line",
			Points: points,
			Render: func(image gg.Context) {
				image.DrawLine(points[0].x, points[0].y, points[1].x, points[1].y)
				image.Stroke()
			},
		}
	} else if strokeType == "right angle" {
		newPoint = glyph.getRandomCoordinatePoint()
		points = append(points, newPoint)
		newPoint = glyph.getRandomCoordinatePoint()
		newPoint.y = points[0].y
		points = append(points, newPoint)
		newPoint = glyph.getRandomCoordinatePoint()
		newPoint.x = points[1].x
		points = append(points, newPoint)
		newStroke = Stroke{
			Type:   "line",
			Points: points,
			Render: func(image gg.Context) {
				image.DrawLine(points[0].x, points[0].y, points[1].x, points[1].y)
				image.DrawLine(points[1].x, points[1].y, points[2].x, points[2].y)
				image.Stroke()
			},
		}
	} else if strokeType == "horizontal wedge" {
		newPoint = glyph.getRandomCoordinatePoint()
		points = append(points, newPoint)
		newPoint = glyph.getRandomCoordinatePoint()
		newPoint.x = points[0].x
		points = append(points, newPoint)
		newPoint = glyph.getRandomCoordinatePoint()
		newPoint.x = points[1].x + float64(glyph.Width/glyph.GridSizeX)
		newPoint.y = (points[0].y + points[1].y) / 2
		points = append(points, newPoint)
		newStroke = Stroke{
			Type:   "line",
			Points: points,
			Render: func(image gg.Context) {
				image.DrawLine(points[0].x, points[0].y, points[2].x, points[2].y)
				image.DrawLine(points[2].x, points[2].y, points[1].x, points[1].y)
				image.Stroke()
			},
		}
	} else if strokeType == "vertical wedge" {
		newPoint = glyph.getRandomCoordinatePoint()
		points = append(points, newPoint)
		newPoint = glyph.getRandomCoordinatePoint()
		newPoint.y = points[0].y
		points = append(points, newPoint)
		newPoint = glyph.getRandomCoordinatePoint()
		newPoint.y = points[1].y + float64(glyph.Height/glyph.GridSizeY)
		newPoint.x = (points[0].x + points[1].x) / 2
		points = append(points, newPoint)
		newStroke = Stroke{
			Type:   "line",
			Points: points,
			Render: func(image gg.Context) {
				image.DrawLine(points[0].x, points[0].y, points[2].x, points[2].y)
				image.DrawLine(points[2].x, points[2].y, points[1].x, points[1].y)
				image.Stroke()
			},
		}
	}

	return newStroke
}
