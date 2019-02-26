package wrisysgen

import (
	"math/rand"
	"time"

	"github.com/fogleman/gg"
)

// Glyph represents a single glyph
type Glyph struct {
	Representation string
	Points         []Point
}

// Point is a coordinate on a 2D plane
type Point struct {
	x float64
	y float64
}

func connectPointSet(image gg.Context, pointSet []Point) {
	nextType := randomScalar(10)
	arcPoint := Point{0.0, 0.0}
	for index, point := range pointSet {
		if index == 0 {
			image.MoveTo(point.x, point.y)
		} else {
			nextType = randomScalar(10)
			if nextType > 2 {
				image.LineTo(point.x, point.y)
			} else {
				arcPoint = nudgePoint(point)
				image.QuadraticTo(arcPoint.x, arcPoint.y, point.x, point.y)
			}

		}
	}
	image.SetLineWidth(6)
	image.Stroke()
	image.ClosePath()
}

func drawCurvedLine(image gg.Context, startingPoint Point, maximum int) {
	image.MoveTo(startingPoint.x, startingPoint.y)

	arcPoint := randomPoint(maximum)
	nextPoint := randomPoint(maximum)

	image.QuadraticTo(arcPoint.x, arcPoint.y, nextPoint.x, nextPoint.y)

	image.SetLineWidth(6)
	image.Stroke()
	image.ClosePath()
}

func drawLine(image gg.Context, startingPoint Point, maximum int) {
	numberOfLines := randomScalar(4)
	image.MoveTo(startingPoint.x, startingPoint.y)

	nextPoint := Point{0.0, 0.0}

	for i := 0; i < numberOfLines; i++ {
		nextPoint = randomPoint(maximum)
		image.LineTo(nextPoint.x, nextPoint.y)
	}
	image.SetLineWidth(6)
	image.Stroke()
	image.ClosePath()
}

func drawPattern(image gg.Context, patternType string, point Point, maximum int) {
	if patternType == "C" {
		image.DrawCircle(point.x, point.y, 32)
	} else if patternType == "L" {
		drawLine(image, point, maximum)
	} else if patternType == "A" {
		drawCurvedLine(image, point, maximum)
	} else if patternType == "S" {
		pointSet := generatePointSet(maximum)
		connectPointSet(image, pointSet)
	} else if patternType == "R" {
		pointSet := reducePointSet(generateRegularPointSet(maximum))
		connectPointSet(image, pointSet)
	} else if patternType == "E" {
		drawRegularCharacter(image, maximum)
	} else if patternType == "T" {
		pointSet := generateTurtlePointSet(maximum)
		//connectPointSet(image, pointSet)
		drawPoints(image, reducePointSet(perturbPoints(pointSet)))
	}
}

func drawPoints(image gg.Context, points []Point) {
	for _, point := range points {
		image.DrawPoint(point.x, point.y, 8)
	}
	image.Stroke()
}

func drawRegularCharacter(image gg.Context, maximum int) {
	startingPoint := randomStartingPoint(maximum)
	point2 := Point{startingPoint.x + 128.0, startingPoint.y}
	point3 := Point{point2.x, point2.y + 64.0}
	point4 := Point{point3.x - 32.0, point2.y}
	i := 0
	q := Point{0.0, 0.0}

	var points []Point

	points = append(points, point2)
	points = append(points, point3)
	points = append(points, point4)

	image.MoveTo(startingPoint.x, startingPoint.y)

	for _, point := range points {
		i = randomScalar(10)
		if i > 3 {
			image.LineTo(point.x, point.y)
		} else {
			q = nudgePoint(point)
			image.QuadraticTo(q.x, q.y, point.x, point.y)
		}
	}

	image.SetLineWidth(6)
	image.Stroke()
	image.ClosePath()
}

func drawVerticalWedge(image gg.Context, maximum int) {
	startingPoint := randomStartingPoint(maximum)
	image.MoveTo(startingPoint.x, startingPoint.y)
	image.LineTo(startingPoint.x, startingPoint.y+64.0)
	image.QuadraticTo(startingPoint.x+16.0, startingPoint.y+16.0, startingPoint.x+32, startingPoint.y)
	image.SetLineWidth(6)
	image.Stroke()
	image.ClosePath()
}

func generateGlyph(representation string, patternSet []string, maximum int) {
	numberOfPatterns := rand.Intn(2) + 1
	point := randomPoint(maximum)
	nextPoint := Point{0.0, 0.0}
	patternType := ""

	image := gg.NewContext(maximum, maximum)
	image.MoveTo(point.x, point.y)

	for i := 0; i < numberOfPatterns; i++ {
		patternType = randomPattern(patternSet)
		nextPoint = randomPoint(maximum)
		drawPattern(*image, patternType, nextPoint, maximum)
	}

	fileName := "./output/" + representation + ".png"

	image.SavePNG(fileName)
}

func generatePatternSet() []string {
	var set []string
	set = append(set, "T")
	return set
}

func generatePointSet(maximum int) []Point {
	var points []Point

	startingPoint := randomStartingPoint(maximum)
	nextPoint := startingPoint
	previousPoint := startingPoint

	points = append(points, startingPoint)

	for i := 0; i < randomScalar(4)+1; i++ {
		nextPoint = nudgePoint(previousPoint)
		points = append(points, nextPoint)
		previousPoint = nextPoint
	}

	return points
}

func generateRegularPointSet(maximum int) []Point {
	var points []Point
	point := Point{0.0, 0.0}

	stepSize := maximum / 8

	for x := 0; x < maximum; x += stepSize {
		for y := 0; y < maximum; y += stepSize {
			point = Point{float64(x), float64(y)}
			points = append(points, point)
		}
	}

	return points
}

func generateTurtlePointSet(maximum int) []Point {
	var points []Point
	point := Point{0.0, 0.0}
	startingPoint := randomStartingPoint(maximum)
	previousPoint := startingPoint

	stepSize := float64(maximum / 8)
	travelSize := stepSize

	numberOfSteps := randomScalar(4) + 2

	possibleDirections := []string{}
	lastDirection := "left"

	j := 0

	for i := 0; i < numberOfSteps; i++ {
		possibleDirections = nil
		travelSize = float64(stepSize * float64(randomScalar(3)+1))

		if previousPoint.x+travelSize < float64(maximum) {
			possibleDirections = append(possibleDirections, "right")
		}
		if previousPoint.x-travelSize > 0.0 {
			possibleDirections = append(possibleDirections, "left")
		}
		if previousPoint.y+travelSize < float64(maximum) && lastDirection != "down" && lastDirection != "up" {
			possibleDirections = append(possibleDirections, "down")
		}
		if previousPoint.y-travelSize > 0.0 && lastDirection != "up" && lastDirection != "down" {
			possibleDirections = append(possibleDirections, "up")
		}

		j = randomScalar(len(possibleDirections))

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
	newX := point.x + float64(randomScalar(64)) - float64(randomScalar(64))
	newY := point.y + float64(randomScalar(64)) - float64(randomScalar(64))

	if newX < 0 || newY < 0 || newX > 256 || newY > 256 {
		return nudgePoint(point)
	}

	return Point{newX, newY}
}

func randomScalar(maximum int) int {
	return rand.Intn(maximum)
}

func randomPattern(patternSet []string) string {
	randomItem := patternSet[rand.Intn(len(patternSet))]
	return randomItem
}

func randomPoint(maximum int) Point {
	x := randomScalar(maximum)
	if x < 32 {
		x = 32
	} else if x > 220 {
		x = 220
	}
	y := randomScalar(maximum)
	if y < 32 {
		y = 32
	} else if y > 220 {
		y = 220
	}
	return Point{float64(x), float64(y)}
}

func randomStartingPoint(maximum int) Point {
	return randomPoint(int(maximum / 4))
}

func reducePointSet(points []Point) []Point {
	var newPoints []Point
	i := 0

	for _, point := range points {
		i = randomScalar(10)
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

// Generate procedurally generates a set of glyphs
func Generate() {
	rand.Seed(time.Now().UnixNano())
	maximum := 256

	patterns := generatePatternSet()

	var symbols []string

	letters := [26]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}

	for _, letter := range letters {
		generateGlyph(letter, patterns, maximum)
		symbols = append(symbols, letter)
	}

	renderHTML(symbols)
}
