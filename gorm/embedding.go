package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Result1 struct {
	Embedding []float32 `json:"data"`
	Msg       string    `json:"msg"`
	Success   bool      `json:"success"`
}

func main() {
	url := "http://10.75.212.40:8301/jinaembedding"

	// 设置请求体
	data := map[string]interface{}{
		"inputs":    "帮我梳理代码库",
		"normalize": true,
		"truncate":  true,
	}

	jsonBody, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	// 发起POST请求
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonBody))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// 读取响应
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	var result Result
	_ = json.Unmarshal(body, &result)

	// 输出响应
	fmt.Println("Response Status:", resp.Status)
	fmt.Printf("Response Body: %v\n", result)
}
