package main

import (
	"fmt"
	"regexp"
)

func main() {
	var DeviceDevIDPattern = regexp.MustCompile("^[a-z0-9]([a-z0-9-.]{0,61}[a-z0-9])?$")

	fmt.Println(DeviceDevIDPattern.MatchString("aa"))

}
