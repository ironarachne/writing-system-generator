package wrisysgen

import (
	"fmt"
	"html/template"
	"os"
)

const templateHTML = `
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

func renderHTML(symbols []string) {
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

	err = t.Execute(writer, symbols)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer writer.Close()
}
