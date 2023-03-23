package main

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"
)

type Movie struct {
	Title  string
	Year   int  `json:"released"`
	Color  bool `json:"color,omitempty"`
	Actors []string
}

var movies = []Movie{
	{"Casablanca", 1942, false, []string{"Humphrey Bogart", "Ingrid Bergman"}},
}

func main() {
	data, err := json.MarshalIndent(movies, "", "  ")
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}
	fmt.Printf("%s\n", reflect.TypeOf(data).String())
	fmt.Printf("%s\n", data)
	var titles []struct{ Title string }
	if err := json.Unmarshal(data, &titles); err != nil {
		log.Fatalf("JSON unmarshaling failed: %s", err)
	}
	fmt.Println(titles)

}
