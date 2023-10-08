// @program: goPractice
// @author: zhangyongxiang(zyxfox@foxmail.com)
// @create: 2021-04-27 01:46
package main

import (
	"fmt"
)

func half(numberToHalf int) (int, error) {
	if numberToHalf%2 != 0 {
		return -1, fmt.Errorf("can not half %v", numberToHalf)
	}
	return numberToHalf / 2, nil
}
func main4() {
	n, err := half(19)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(n)
}
