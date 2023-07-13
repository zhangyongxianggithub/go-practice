package main

import (
	"fmt"
	"regexp"
)

var chinese = regexp.MustCompile("^[\\w\u4e00-\u9fa5]{1,30}$")

var chinese_en = regexp.MustCompile("^[\u4e00-\u9fa5[:punct:][:alpha:]]+$")

func main() {

	exp := "^((?P<schema>http(s)?)://)?(?P<server>[^:]+)(:(?P<port>\\d+)?)?$"
	reg := regexp.MustCompile(exp)
	r := reg.FindStringSubmatch("https://192.168.0.1:8086")

	schema := reg.SubexpIndex("schema")
	server := reg.SubexpIndex("server")
	port := reg.SubexpIndex("port")
	fmt.Println(schema, r[schema], server, r[server], port, r[port])
	fmt.Println(chinese.MatchString("张112abcVVVV2"))
	fmt.Println(chinese_en.MatchString("张das...,,////"))

}
