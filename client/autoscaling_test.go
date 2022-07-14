package client

import (
	awsutil "awsdisc/client/util"
	"encoding/json"
	"errors"
	"fmt"
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

	values, err := awsutil.JsonPath(jsonBlob, "$.AutoScalingGroups[*].Instances[0].InstanceId")
	//value, err := awsutil.JsonPath(jsonBlob, "$.AutoScalingGroups[*].Instances[1:].InstanceId")
	if err != nil {
		t.Error(err)
	}
	//t.Log(reflect.TypeOf(values).Elem())
	t.Logf("%v", values)

}

func TestJsonObjectPath(t *testing.T) {
	rawjson := `{ "key" : "value" }`

	values, err := awsutil.JsonPath([]byte(rawjson), "$.key")
	if err != nil {
		t.Error(err)
	}

	// JsonPath's return type is []interface or []interface{}
	switch x := values.(type) {
	case []interface{}:
		fmt.Println("values is slice interface...")
	case interface{}:
		fmt.Println("values is single interface...")
		var s []interface{}
		s = append(s, x)
		t.Log(s)
	default:
		err := errors.New("invalid type...")
		t.Error(err)
	}

	t.Logf("%+v", values)
}
