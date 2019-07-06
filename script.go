package script

import (
	"fmt"
	"html/template"
	"os"
)

// Script is a writing system script
type Script struct {
	Glyphs []Glyph
}

// Generate procedurally generates a set of glyphs
func Generate() Script {
	script := randomScript(128, 128)

	return script
}

func randomScript(glyphWidth int, glyphHeight int) Script {
	var newGlyph Glyph
	script := Script{}
	glyphs := []Glyph{}

	letters := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}

	for _, l := range letters {
		newGlyph = generateGlyph(l, glyphWidth, glyphHeight)
		glyphs = append(glyphs, newGlyph)
	}

	script.Glyphs = glyphs

	return script
}

// RenderHTML renders an HTML version of a script
func (script Script) RenderHTML() {
	symbols := []string{}
	templateHTML := `
<!DOCTYPE html>
<html>
    <head>
        <title>Writing System</title>
        <style type="text/css">
            body, html { font-size: 28px; }
            div.container { display: flex; flex-wrap: wrap; width: 1600px; margin: 1rem auto; }
            div.cell { width: 300px; height: 300px; margin: 1rem; text-align: center; font-weight: 700; }
            div.cell > img { display: block; }
        </style>
    </head>
    <body>
		<div class="container">
			{{range $index, $element := .}}
            <div class="cell">
                <img src="{{ $element }}.png">
                <p>{{ $element }}</p>
            </div>
            {{end}}
        </div>
    </body>
</html>
`

	writer, err := os.Create("./output/index.html")
	if err != nil {
		fmt.Println(err)
		return
	}

	t, err := template.New("htmlIndex").Parse(templateHTML)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, g := range script.Glyphs {
		symbols = append(symbols, g.Representation)
	}

	err = t.Execute(writer, symbols)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer writer.Close()
}

// RenderGlyphImages runs the render process for each glyph in a script
func (script Script) RenderGlyphImages() {
	for _, g := range script.Glyphs {
		g.Render()
	}
}
