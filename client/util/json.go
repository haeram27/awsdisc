package client

import (
	apps "awsdisc/apps"
	"bytes"
	"encoding/json"

	"github.com/PaesslerAG/jsonpath"
)

func PrintPrettyJson(jsonBlob []byte) {
	var prettyJSON bytes.Buffer
	err := json.Indent(&prettyJSON, jsonBlob, "", "  ")
	if err != nil {
		apps.Logs.Error("JSON parse error: ", err.Error())
	} else {
		apps.Logs.Debug(prettyJSON.String())
	}
}

func PrettyJson(jsonBlob []byte) *bytes.Buffer {
	var prettyJSON bytes.Buffer
	json.Indent(&prettyJSON, jsonBlob, "", "  ")
	return &prettyJSON
}

func PrettyMarshalJson(v interface{}) ([]byte, error) {
	return json.MarshalIndent(v, "", "  ")
}

func JsonPath(jsonBlob []byte, path string) (any, error) {
	jsonStrt := interface{}(nil)
	err := json.Unmarshal(jsonBlob, &jsonStrt)
	if err != nil {
		return jsonStrt, err
	}

	return jsonpath.Get(path, jsonStrt)
}
