package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	file, err := os.ReadFile("http/file/plaintext.txt")
	if err != nil {
		return
	}
	text := string(file)
	requestBody := map[string]any{
		"text":          text,
		"text_type":     "markdown",
		"chunk_size":    500,
		"chunk_overlap": 0,
	}

	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		panic(err)
	}

	// 发起POST请求
	resp, err := http.Post("http://10.63.41.18:8865/rest/v1/text/split", "application/json", bytes.NewBuffer(jsonBody))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// 读取响应
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	// 输出响应
	fmt.Println("Response Status:", resp.Status)
	fmt.Printf("Response Body: %v\n", string(body))
	responseFile, err := os.OpenFile("http/file/response.json", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	_, _ = io.WriteString(responseFile, string(body))

}
