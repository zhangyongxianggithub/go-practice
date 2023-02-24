package main

import (
	"fmt"
	"log"

	"gopl/ch5/links"
)

func crawl(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}
func main() {
	fmt.Println(crawl("https://www.baidu.com"))
}
