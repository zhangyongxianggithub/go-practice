package main

import (
	"fmt"
	"os"
	"text/template"
)

func main() {

	// const templ = `{{3.1415926 | printf "%.2f"}}`

	const templ = `{{ printf "%.2f" 3.1415926 }}`

	t, err := template.New("format").Parse(templ)
	if err != nil {
		fmt.Println("Could not parse template:", err)
		return

	}
	err = t.Execute(os.Stdout, map[string]string{})
	if err != nil {
		return
	}

}
