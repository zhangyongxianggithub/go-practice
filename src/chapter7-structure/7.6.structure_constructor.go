//@program: goPractice
//@author: zhangyongxiang(zyxfox@foxmail.com)
//@create: 2021-04-26 01:59
package main

import "fmt"

type Alarm struct {
	Time  string
	Sound bool
}

func NewAlarm(time string) Alarm {
	a := Alarm{
		Time:  time,
		Sound: true,
	}
	return a
}
func main6() {
	fmt.Printf("%+v\n", NewAlarm("07:00"))
}
