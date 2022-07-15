package client

import (
	awsutil "awsdisc/client/util"
	"encoding/json"
	"testing"
)

func TestASCLDescribeAutoScalingGroupsAll(t *testing.T) {
	var jsonBlob []byte

	result, err := ASCLDescribeAutoScalingGroupsCmd(AwsConfig())
	if err != nil {
		t.Error(err)
	} else {
		jsonBlob, err = json.Marshal(result)
		if err != nil {
			t.Error(err)
		}
	}
	//t.Log(awsutil.PrettyJson(jsonBlob).String())

	values := awsutil.JsonPath(jsonBlob, "$.AutoScalingGroups[*].Instances[0].InstanceId")
	//value, err := awsutil.JsonPath(jsonBlob, "$.AutoScalingGroups[*].Instances[1:].InstanceId"))

	t.Logf("%v", values)

}

func TestJsonObjectPath(t *testing.T) {
	rawjson := `{ "key" : "value" }`

	values := awsutil.JsonPath([]byte(rawjson), "$.key")

	t.Logf("%+v", values)
}
