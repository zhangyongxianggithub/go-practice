package main

import (
	"fmt"

	"github.com/goccy/go-json"
)

const j = `
{
	"aaaa": "aaaaa\
             dadsadsa"
}

`

func main() {
	m := make(map[string]string)
	err := json.UnmarshalNoEscape([]byte(j), &m)
	if err != nil {
		return 
	}
	fmt.Println(m)

}
