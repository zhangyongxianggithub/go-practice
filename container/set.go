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

	s1 := set.NewSet[string]()
	s1.Add("234")
	s1.Add("456")
	s.Append(s1.ToSlice()...)
	marshal, err := json.Marshal(s)
	if err != nil {
		return
	}
	fmt.Println(string(marshal))
}
