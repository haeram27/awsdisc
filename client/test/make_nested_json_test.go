package test

import (
	"awsdisc/client"
	"encoding/json"
	"fmt"
	"testing"
)

func TestMakeNestedJson(t *testing.T) {
	inner1 := `{"intValue": 1234}`
	inner2 := `{"boolValue": true}`

	// consolidate
	data := make(map[string]interface{})
	var innerJson interface{}
	json.Unmarshal([]byte(inner1), &innerJson)
	data["inner1"] = innerJson
	json.Unmarshal([]byte(inner2), &innerJson)
	data["inner2"] = innerJson

	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("could not marshal json: ", err)
		return
	}

	fmt.Println(string(jsonData))
	client.PrintPrettyJson(jsonData)
}
