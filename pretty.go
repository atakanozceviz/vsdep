package vsdep

import (
	"encoding/json"
	"fmt"
)

// PrettyPrint json to standard output.
func PrettyPrint(v interface{}) (err error) {
	b, err := json.MarshalIndent(v, "", "  ")
	if err == nil {
		fmt.Println(string(b))
	}
	return
}
