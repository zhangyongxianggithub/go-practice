//@program: goPractice
//@author: zhangyongxiang(zyxfox@foxmail.com)
//@create: 2021-04-26 01:00
package main

import "fmt"

func main10() {
	defer fmt.Println("I am run after the function completes")
	fmt.Println("Hello defer")
}
