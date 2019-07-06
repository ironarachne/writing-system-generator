package script

import (
	"math/rand"

	"github.com/fogleman/gg"
)

// Point is a coordinate on a 2D plane
type Point struct {
	x float64
	y float64
}

func connectPointSet(image gg.Context, pointSet []Point) {
	//nextType := rand.Intn(10)
	//arcPoint := Point{0.0, 0.0}
	for index, point := range pointSet {
		if index == 0 {
			image.MoveTo(point.x, point.y)
		} else {
			/*
				nextType = rand.Intn(10)
				if nextType > 2 {
					image.LineTo(point.x, point.y)
				} else {
					arcPoint = nudgePoint(point)
					image.QuadraticTo(arcPoint.x, arcPoint.y, point.x, point.y)
				}*/
			image.LineTo(point.x, point.y)

		}
	}
	image.SetLineWidth(6)
	image.SetHexColor("AA0000")
	image.Stroke()
	image.ClosePath()
}

func generatePointSet(width int, height int) []Point {
	var points []Point

	startingPoint := randomPoint(width, height)
	nextPoint := startingPoint
	previousPoint := startingPoint

	points = append(points, startingPoint)

	for i := 0; i < rand.Intn(4)+1; i++ {
		nextPoint = nudgePoint(previousPoint)
		points = append(points, nextPoint)
		previousPoint = nextPoint
	}

	return points
}

func generateRegularPointSet(width int, height int) []Point {
	var points []Point
	point := Point{0.0, 0.0}

	stepSize := width / 8

	for x := 0; x < width; x += stepSize {
		for y := 0; y < height; y += stepSize {
			point = Point{float64(x), float64(y)}
			points = append(points, point)
		}
	}

	return points
}

func generateTurtlePointSet(width int, height int) []Point {
	var points []Point
	point := Point{0.0, 0.0}
	startingPoint := randomPoint(width, height)
	previousPoint := startingPoint

	stepSize := float64(width / 8)
	travelSize := stepSize

	numberOfSteps := rand.Intn(4) + 2

	possibleDirections := []string{}
	lastDirection := "left"

	j := 0

	for i := 0; i < numberOfSteps; i++ {
		possibleDirections = nil
		travelSize = float64(stepSize * float64(rand.Intn(3)+1))

		if previousPoint.x+travelSize < float64(width) {
			possibleDirections = append(possibleDirections, "right")
		}
		if previousPoint.x-travelSize > 0.0 {
			possibleDirections = append(possibleDirections, "left")
		}
		if previousPoint.y+travelSize < float64(height) && lastDirection != "down" && lastDirection != "up" {
			possibleDirections = append(possibleDirections, "down")
		}
		if previousPoint.y-travelSize > 0.0 && lastDirection != "up" && lastDirection != "down" {
			possibleDirections = append(possibleDirections, "up")
		}

		j = rand.Intn(len(possibleDirections))

		if possibleDirections[j] == "right" {
			point = Point{previousPoint.x + travelSize, previousPoint.y}
			lastDirection = "right"
		} else if possibleDirections[j] == "left" {
			point = Point{previousPoint.x - travelSize, previousPoint.y}
			lastDirection = "left"
		} else if possibleDirections[j] == "down" {
			point = Point{previousPoint.x, previousPoint.y + travelSize}
			lastDirection = "down"
		} else if possibleDirections[j] == "up" {
			point = Point{previousPoint.x, previousPoint.y - travelSize}
			lastDirection = "up"
		}

		points = append(points, point)

	}

	return points
}

func nudgePoint(point Point) Point {
	newX := point.x + float64(rand.Intn(64)) - float64(rand.Intn(64))
	newY := point.y + float64(rand.Intn(64)) - float64(rand.Intn(64))

	if newX < 0 || newY < 0 || newX > 256 || newY > 256 {
		return nudgePoint(point)
	}

	return Point{newX, newY}
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

func reducePointSet(points []Point) []Point {
	var newPoints []Point
	i := 0

	for _, point := range points {
		i = rand.Intn(10)
		if i > 3 {
			newPoints = append(newPoints, point)
		}
	}

	return newPoints
}

func perturbPoints(points []Point) []Point {
	var newPoints []Point

	for _, point := range points {
		newPoints = append(newPoints, nudgePoint(point))
	}

	return newPoints
}
