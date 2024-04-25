package main

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
)

func main() {
	fmt.Println(SHA1("aaaaaaaaa.java"))
}

func SHA1(str string) string {
	c := sha1.New()
	c.Write([]byte(str))
	return hex.EncodeToString(c.Sum(nil))
}
