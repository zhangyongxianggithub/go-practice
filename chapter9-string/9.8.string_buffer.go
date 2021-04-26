//@program: goPractice
//@author: zhangyongxiang(zyxfox@foxmail.com)
//@create: 2021-04-27 01:18
package main

import (
	"bytes"
	"fmt"
)

func main8() {
	var buffer bytes.Buffer
	for i := 0; i < 500; i++ {
		buffer.WriteString("Z")
	}
	fmt.Println(buffer.String())
}
