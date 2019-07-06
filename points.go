package script

import (
	"math/rand"
)

// Point is a coordinate on a 2D plane
type Point struct {
	x float64
	y float64
}

// Equals compares two points to see if they're equivalent
func (point Point) Equals(otherPoint Point) bool {
	if point.x == otherPoint.x && point.y == otherPoint.y {
		return true
	}

	return false
}

func (glyph Glyph) getPointForCoordinate(x, y int) Point {
	pointX := float64(x * (glyph.Width / glyph.GridSizeX))
	pointY := float64(y * (glyph.Height / glyph.GridSizeY))

	return Point{
		x: pointX,
		y: pointY,
	}
}

func (glyph Glyph) getRandomCoordinatePoint() Point {
	x := rand.Intn(glyph.GridSizeX)
	y := rand.Intn(glyph.GridSizeY)

	return glyph.getPointForCoordinate(x, y)
}

func (glyph Glyph) getRandomDifferentCoordinatePoint(oldPoint Point) Point {
	x := rand.Intn(glyph.GridSizeX)
	y := rand.Intn(glyph.GridSizeY)

	newPoint := glyph.getPointForCoordinate(x, y)

	for oldPoint.Equals(newPoint) {
		x = rand.Intn(glyph.GridSizeX)
		y = rand.Intn(glyph.GridSizeY)

		newPoint = glyph.getPointForCoordinate(x, y)
	}

	return newPoint
}

func (glyph Glyph) getCoordinateForPoint(point Point) Point {
	x := point.x / float64(glyph.GridSizeX)
	y := point.y / float64(glyph.GridSizeY)

	return Point{
		x: x,
		y: y,
	}
}

func randomPoint(width int, height int) Point {
	x := rand.Intn(width)
	y := rand.Intn(height)

	newPoint := Point{
		x: float64(x),
		y: float64(y),
	}

	return newPoint
}
