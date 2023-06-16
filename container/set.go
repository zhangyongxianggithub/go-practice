package main

import (
	"encoding/json"
	"fmt"

	set "github.com/deckarep/golang-set/v2"
)

func main() {
	s := set.NewSet[string]()
	s.Add("123")
	s.Add("234")
	marshal, err := json.Marshal(s)
	if err != nil {
		return
	}
	fmt.Println(string(marshal))
}
