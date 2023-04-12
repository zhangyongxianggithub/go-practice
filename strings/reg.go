package main

import (
	"fmt"
	"regexp"
)

func main() {

	exp := "^((?P<schema>http(s)?)://)?(?P<server>[^:]+)(:(?P<port>\\d+)?)?$"
	reg := regexp.MustCompile(exp)
	r := reg.FindStringSubmatch("https://192.168.0.1:8086")

	schema := reg.SubexpIndex("schema")
	server := reg.SubexpIndex("server")
	port := reg.SubexpIndex("port")
	fmt.Println(schema, r[schema], server, r[server], port, r[port])

}
