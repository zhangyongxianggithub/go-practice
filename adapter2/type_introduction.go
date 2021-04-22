package main

import "fmt"

func sayHello(s string) string {
	return "Hello " + s
}
func main1() {
	fmt.Println(sayHello("George"))
}
