package printer

import (
	"encoding/json"
	"fmt"
)

func PrintAny(object any) string {
	// TODO check if zero value for all data type
	if err, ok := object.(error); ok {
		return fmt.Sprintf("error: %s", err)
	}
	bytes, err := json.Marshal(object)
	if err != nil {
		return fmt.Sprintf("json error: %s", err)
	} else {
		return string(bytes)
	}
}
