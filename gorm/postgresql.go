package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"io"
	"net/http"
	"time"
)

type Result struct {
	Embedding []float32 `json:"data"`
	Msg       string    `json:"msg"`
	Success   bool      `json:"success"`
}

type CommonQueryVector struct {
	ID        string
	Vector    pq.Float32Array `gorm:"type:float4[]"`
	Content   string
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

func (CommonQueryVector) TableName() string {
	return "common_query_vector"
}

var url = "http://10.75.212.40:8301/jinaembedding"

func genVector(query string) *Result {

	// 设置请求体
	data := map[string]interface{}{
		"inputs":    query,
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
	return &result
}

func main() {
	db, err := gorm.Open(postgres.Open("postgres://postgres:Comate%23123@10.63.41.18:8432/autowork_external?sslmode=disable"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	var queryVectors []CommonQueryVector
	result := db.Find(&queryVectors)
	if result.Error != nil {
		panic(result.Error)
	}
	for _, queryVector := range queryVectors {
		result := genVector(queryVector.Content)
		if result != nil {
			err := db.Model(&CommonQueryVector{}).Where("id = ?", queryVector.ID).Updates(CommonQueryVector{Vector: result.Embedding}).Error
			if err != nil {
				fmt.Printf("fail to recreate vector for query: %s, content: %s\n", queryVector.ID, queryVector.Content)
			} else {
				fmt.Printf("recreate vector for query: %s, content: %s\n", queryVector.ID, queryVector.Content)
			}
		}
	}
}
