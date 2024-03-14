package api

import (
	"bytes"
	"encoding/json"
	"log"
)

func StructToJsonString(value any) string {
	jsonBytes, err := json.Marshal(value)
	if err != nil {
		log.Fatalf("%v", err)
	}
	var str bytes.Buffer
	err = json.Indent(&str, jsonBytes, "", "\t")
	if err != nil {
		log.Fatalf("%v", err)
	}
	result := str.String()
	return result
}
