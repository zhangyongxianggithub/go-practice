package main

import (
	"encoding/json"
	"fmt"
)

const j = `{
	"field": "dsadsa \\
dsadsads"
}`

func main() {
	m := make(map[string]any)
	err := json.Unmarshal([]byte(j), &m)
	if err != nil {
		return
	}
	fmt.Println(m)
}
