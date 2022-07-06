package client

import (
	"encoding/json"
	"testing"
)

func TestPrintPrettyJson(t *testing.T) {
	inner1 := `{"Name": "iron man"}`
	inner2 := `{"Type": "developer"}`

	// consolidate
	jsonStrt := make(map[string]interface{})
	var innerJson interface{}
	json.Unmarshal([]byte(inner1), &innerJson)
	jsonStrt["inner1"] = innerJson
	json.Unmarshal([]byte(inner2), &innerJson)
	jsonStrt["inner2"] = innerJson

	jsonBlob, err := json.Marshal(jsonStrt)
	if err != nil {
		t.Error(err)
		return
	}

	PrintPrettyJson(jsonBlob)
}

func TestPrettyJson(t *testing.T) {
	jsonBlob := `{
		Name: "iron man",
		Type: "developer",
	}`

	t.Log(PrettyJson([]byte(jsonBlob)).String())
}

func TestPrettyMarshalJson(t *testing.T) {
	jsonStrt := struct {
		Name string
		Type string
	}{
		Name: "iron man",
		Type: "developer",
	}

	jsonBlob, err := PrettyMarshalJson(jsonStrt)
	if err != nil {
		t.Error(err)
	}
	t.Log(string(jsonBlob))
}

func TestJsonPath(t *testing.T) {
	jsonBlob := `{ 
		"Name": { 
			"First": "man",
			"Last": "iron"
		},
		"Type": "developer"
	}`

	path := "$.Name.Last"

	v, err := JsonPath([]byte(jsonBlob), path)
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v", v)
}
