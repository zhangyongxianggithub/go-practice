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

	for i := 0; i < 1; i++ {
		r := rand.Intn(100)
		request("benchmark"+strconv.Itoa(i), r)
		// deleteRule("benchmark" + strconv.Itoa(i))
	}
	// request("benchmark"+strconv.Itoa(1001), 40)
}

func deleteRule(ruleName string) {

	url := "http://10.68.115.50:8381/rules/" + ruleName
	method := "DELETE"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

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
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(res.Body)

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
func request(ruleName string, temperature int) {
	url := "http://10.162.195.165:8381/rules"
	method := "POST"

	tpl, _ := template.New("rule").Parse(`{
    "id": "{{ .RuleName}}",
    "name": "{{ .RuleName }}",
    "sql": "select kind,       meta->device as meta_device,       meta->accessTemplate as meta_access_template,       meta->deviceProduct as meta_device_product,       meta->node as meta_node,       meta->nodeProduct as meta_node_product,       cast(content->blink->timestamp,\"string\") as ts,       content->blink->properties->A_AEROBICDO1_END as a_aerobicdo1_end,       content->blink->properties->A_AEROBICDO1_START as a_aerobicdo1_start,    content->blink->properties->A_ALTERNATE as a_alternate,    content->blink->properties->A_ANAEROBICORP as a_anaerobicorp,    content->blink->properties->A_HYPOXIAORP as a_hypoxiaorp,    content->blink->properties->A_N1NH4_N_NO3_N as a_n1nh4_n_no3_n,    content->blink->properties->A_N1NO3_N as a_n1no3_n,    content->blink->properties->A_N1_AEROBICO_MLSS as a_n1_aerobico_mlss,    content->blink->properties->A_N2_AEROBICO_MLSS as a_n2_aerobico_mlss,    content->blink->properties->A_Q2_IN_FLOW_METER as a_q2_in_flow_meter,    content->blink->properties->A_Q2_OUT_FLOW_METER as a_q2_out_flow_meter,    content->blink->properties->B_AEROBICDO1_END as b_aerobicdo1_end,    content->blink->properties->B_AEROBICDO1_START as b_aerobicdo1_start,    content->blink->properties->B_ALTERNATE as b_alternate,    content->blink->properties->B_ANAEROBICORP as b_anaerobicorp,    content->blink->properties->B_HYPOXIAORP as b_hypoxiaorp,    content->blink->properties->B_N1NH4_N_NO3_N as b_n1nh4_n_no3_n,    content->blink->properties->B_N1NO3_N as b_n1no3_n,    content->blink->properties->B_N3_AEROBICO_MLSS as b_n3_aerobico_mlss,    content->blink->properties->B_N4_AEROBICO_MLSS as b_n4_aerobico_mlss,    content->blink->properties->B_Q1_IN_FLOW_METER as b_q1_in_flow_meter,    content->blink->properties->B_Q1_OUT_FLOW_METER as b_q1_out_flow_meter from chengdujiayao where a_aerobicdo1_end>0",
    "actions": [
        {
            "mqtt": {
                "insecureSkipVerify": true,
                "server": "10.68.115.50:8883",
                "topic": "lakeinsight/benchmark/result1",
                "omitIfEmpty": false,
                "sendSingle": false,
                "format": "json",
                "bufferLength": 1024,
                "enableCache": false,
                "runAsync": false
            }
        }
    ],
    "options": {
        "isEventTime": false,
        "lateTolerance": 1000,
        "concurrency": 16,
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
	err := tpl.Execute(buffer, map[string]any{"RuleName": ruleName, "Temperature": temperature})
	if err != nil {
		return
	}

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
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(res.Body)

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
