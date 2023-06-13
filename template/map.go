package main

import (
	"encoding/json"
	"fmt"
	"os"
	"text/template"
)

func main() {
	// example()

	json_str := `{"data_point":{"PLC4_COD_WATER_INFLOW_OF_CPU_CHANNEL_EQUIPMENT":140.76,"PLC4_CPU_CHANNEL_COMMUNICATION_STATUS":0,"PLC4_CPU_CHANNEL_EQUIPMENT_CUMULATIVE_FLOW_INFLOW":8048627.5,"PLC4_CPU_CHANNEL_EQUIPMENT_PH_WATER_INFLOW":7.7,"PLC4_TOTALPHOSPHORUS_INFLOW_OF_CPU_CHANNEL_EQUIPMENT":0.86,"PLC4_TOTAL_NITROGEN_INFLOW_OF_CPU_CHANNEL_EQUIPMENT":18.94,"PLC4_WATER_INFLOW_OF_SUSPENDED_SOLIDS_IN_CPU_CHANNEL_EQUIPMENT":0,"WATER_INFLOW":1228.95,"presetpoint1":5},"device_id":"dev-51ct4t9b","ts":1686040965000}`

	const templ = `{{ range $key, $value := .data_point }}metric,device_id={{$.device_id}},data_point_id={{$key}} value={{$value}} {{$.ts}}
{{end}}`
	//const templ = `metric,device_id={{.device_id}}`
	x := map[string]any{}
	err := json.Unmarshal([]byte(json_str), &x)
	if err != nil {
		return
	}

	t, err := template.New("index.html").Parse(templ)
	if err != nil {
		fmt.Println("Could not parse template:", err)
		return

	}
	err = t.Execute(os.Stdout, x)
	if err != nil {
		return
	}

}
func example() {
	fmt.Println("Hello, playground")

	const templ = `Here is what they said
    {{ range $key, $value := . }}
        {{ $key }}:{{ $value }}
    {{end}}
    `
	x := map[string]string{
		"Danny": "The guy really talked about my first time out with you",
		"Doug":  "Well he said I'm really amazing, I did not believe at first",
	}

	t, err := template.New("index.html").Parse(templ)
	if err != nil {
		fmt.Println("Could not parse template:", err)
		return

	}
	t.Execute(os.Stdout, x)
}
