package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
)

var m = map[string]string{
	"test": "aaaa>>>&&",
}

func main() {
	// v, _ := json.Marshal(m)
	// fmt.Printf(string(v))
	ParseHtml()
}

type Html struct {
	Title  string
	Body   string
	Footer string
}

func ParseHtml() {
	htmlJson := Html{
		Title:  "<title>北京欢迎你</title>",
		Body:   "<body>北京是中国的首都，有600多年的建都历史</body>",
		Footer: "<script>js:pop('123')</script>",
	}

	strJson, err := json.Marshal(htmlJson)
	if err == nil {
		// 原始的json串
		fmt.Println("原始json 串", string(strJson))
	}
	var content = string(strJson)
	// 第一种方法，替换'<', '>', '&'
	content = strings.Replace(string(strJson), "\\u003c", "<", -1)
	content = strings.Replace(content, "\\u003e", ">", -1)
	content = strings.Replace(content, "\\u0026", "&", -1)
	fmt.Println("第一种解决办法：", content)

	// 第二种方法，SetEscapeHTML(False)
	bf := bytes.NewBuffer([]byte{})
	jsonEncoder := json.NewEncoder(bf)
	jsonEncoder.SetEscapeHTML(false)
	err = jsonEncoder.Encode(htmlJson)
	if err != nil {
		return
	}
	fmt.Println("第二种解决办法：", bf.String())
}
