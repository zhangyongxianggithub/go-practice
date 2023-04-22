package main

import (
	"encoding/json"
	"fmt"
)

type S struct {
	B `json:",inline"`
	A int `json:"a"`
}

type B struct {
	C int `json:"c"`
}

func main() {

	s := S{B: B{C: 1}, A: 2}
	byt, _ := json.Marshal(s)

	fmt.Println(string(byt))
	var s1 S
	json.Unmarshal(byt, &s1)

	fmt.Println(s1)

}
