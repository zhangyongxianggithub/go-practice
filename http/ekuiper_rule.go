package main

import (
	"bytes"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"strconv"
	"text/template"
)

func main() {

	for i := 0; i < 1000; i++ {
		r := rand.Intn(100)
		request("benchmark"+strconv.Itoa(i), r)
	}
	// request("benchmark"+strconv.Itoa(1001), 40)
}

func request(ruleName string, temperature int) {
	url := "http://10.68.115.50:8381/rules"
	method := "POST"

	tpl, _ := template.New("rule").Parse(`{
    "id": "{{ .RuleName }}",
    "sql": "SELECT temperature,humidity FROM benchmark WHERE temperature > {{ .Temperature }}",
    "actions": [
        {
            "mqtt": {
                "bufferLength": 1024,
                "enableCache": false,
                "format": "json",
                "insecureSkipVerify": true,
                "omitIfEmpty": false,
                "runAsync": false,
                "sendSingle": false,
                "server": "tcp://10.68.115.50:8883",
                "topic": "lakeinsight/benchmark/result"
            }
        }
    ],
    "options": {
        "isEventTime": false,
        "lateTolerance": 1000,
        "concurrency": 128,
        "bufferLength": 1024,
        "sendMetaToSink": false,
        "sendError": true,
        "qos": 0,
        "checkpointInterval": 300000,
        "restartStrategy": {
            "attempts": 0,
            "delay": 1000,
            "multiplier": 2,
            "maxDelay": 30000,
            "jitter": 0.1
        }
    }
}`)

	buffer := bytes.NewBuffer(make([]byte, 0))
	tpl.Execute(buffer, map[string]any{"RuleName": ruleName, "Temperature": temperature})

	client := &http.Client{}
	req, err := http.NewRequest(method, url, buffer)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res.Status)
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
