package util

import (
	"awsdisc/apps"
	"bytes"
	"encoding/json"
)

func PrintPrettyJson(rawJson []byte) {
	var prettyJSON bytes.Buffer
	err := json.Indent(&prettyJSON, rawJson, "", "  ")
	if err != nil {
		apps.Logs.Error("JSON parse error: ", err.Error())
		return
	} else {
		apps.Logs.Debug(prettyJSON.String())
	}
}
