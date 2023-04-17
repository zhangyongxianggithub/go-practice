package main

import (
	"bytes"
	"fmt"
	"text/template"
)

type Model struct {
	Name string `json:"name"`
}

func main() {
	parse, err := template.New("text").Parse(`{{ .Name }}`)
	if err != nil {
		return
	}
	bs := make([]byte, 0)
	buffer := bytes.NewBuffer(bs)
	config := Model{Name: "张永祥"}
	err = parse.Execute(buffer, config)
	if err != nil {
		return

	}
	fmt.Println(string(buffer.String()))

}
