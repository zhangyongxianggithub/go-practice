//@program: goPractice
//@author: zhangyongxiang(zyxfox@foxmail.com)
//@create: 2021-04-26 01:34
package main

import "fmt"

func main() {
	players := make(map[string]int)
	players["cook"] = 32
	players["bairstow"] = 27
	delete(players, "cook")
	fmt.Println(players)
}
