package main

import (
	"math/rand"
	"time"

	script "github.com/ironarachne/writing-system-generator"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	newScript := script.Generate()

	newScript.RenderGlyphImages()
	newScript.RenderHTML()
}
