package fun

import (
	"encoding/json"
	"fmt"
)

func Pretty(v interface{}) string {
	b, _ := json.MarshalIndent(v, "", "\t")
	return string(b)
}

func PrettyPrintln(v interface{}) {
	fmt.Println(Pretty(v))
}
