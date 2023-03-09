package main

import (
	"io"
)

func writeString(w io.Writer, s string) (n int, err error) {
	type stringWriter interface {
		WriteString(string) (int, error)
	}
	if sw, ok := w.(stringWriter); ok { // 使用类型断言推断w的动态类型是否满足新的接口
		return sw.WriteString(s)
	}
	return w.Write([]byte(s))
}
