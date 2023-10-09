package main

import (
	"fmt"
	"os"
	"text/template"
)

func main() {

	// 需要记住管道与函数处理的关系，管道是干什么的，怎么连接多个函数
	// 格式化数字的方式等，基本与C是类似的
	// const templ = `{{3.1415926 | printf "%.2f"}}`

	const templ = `{{ printf "%.2f" 3.1415926 }}`

	t, err := template.New("format").Parse(templ)
	if err != nil {
		fmt.Println("Could not parse template:", err)
		return

	}
	err = t.Execute(os.Stdout, map[string]string{})
	if err != nil {
		return
	}

}
