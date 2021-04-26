//@program: goPractice
//@author: zhangyongxiang(zyxfox@foxmail.com)
//@create: 2021-04-27 00:52
package main

import (
	"errors"
	"fmt"
)

type Robot interface {
	PowerOn() error
}
type T850 struct {
	Name string
}

func (t *T850) PowerOn() error {
	return nil
}

type R2D2 struct {
	Broken bool
}

func (r *R2D2) PowerOn() error {
	if r.Broken {
		return errors.New("R2D2 is broken")
	} else {
		return nil
	}
}
func Boot(robot Robot) error {
	return robot.PowerOn()
}
func main() {
	t := T850{
		Name: "the Terminator",
	}
	r := R2D2{
		Broken: true,
	}
	err := Boot(&r)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Robot powered on")
	}
	err = Boot(&t)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Robot powered on")
	}
}
