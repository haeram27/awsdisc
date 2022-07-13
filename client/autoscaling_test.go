package client

import (
	awsutil "awsdisc/client/util"
	"encoding/json"
	"errors"
	"fmt"
	"testing"
	"time"
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

func TestASCLDescribeAutoScalingGroupsOutputQuery(t *testing.T) {
	rawjson := `{
		"AutoScalingGroups": [
			{
				"Instances": [
					{
						"InstanceId": "i-11"
					},
					{
						"InstanceId": "i-12"
					},
					{
						"InstanceId": "i-13"
					}
				]
			},
			{
				"Instances": [
					{
						"InstanceId": "i-21"
					},
					{
						"InstanceId": "i-22"
					},
					{
						"InstanceId": "i-23"
					}
				]
			},
			{
				"Instances": [
					{
						"InstanceId": "i-31"
					},
					{
						"InstanceId": "i-32"
					},
					{
						"InstanceId": "i-33"
					}
				]
			}
		]
	}`

	rawjson2 := `{
		"AutoScalingGroups": [
			{
				"Instances": [
					{
						"InstanceId": "i-12"
					},
					{
						"InstanceId": "i-13"
					}
				]
			},
			{
				"Instances": [
					{
						"InstanceId": "i-22"
					},
					{
						"InstanceId": "i-23"
					}
				]
			},
			{
				"Instances": [
					{
						"InstanceId": "i-32"
					},
					{
						"InstanceId": "i-33"
					}
				]
			}
		]
	}`

	rawjson3 := `[]`

	values, err := awsutil.JsonPath([]byte(rawjson), "$.AutoScalingGroups[*].Instances[*].InstanceId")
	if err != nil {
		t.Error(err)
	}

	values2, err := awsutil.JsonPath([]byte(rawjson2), "$.AutoScalingGroups[*].Instances[*].InstanceId")
	if err != nil {
		t.Error(err)
	}

	values3, err := awsutil.JsonPath([]byte(rawjson3), "$.AutoScalingGroups[*].Instances[*].InstanceId")
	if err != nil {
		t.Error(err)
	}

	_ = rawjson
	_ = rawjson2
	_ = rawjson3
	_ = values3
	t.Logf("%+v", values)

	start := time.Now()
	s := make([]string, len(values.([]interface{})))
	for i, v := range values.([]interface{}) {
		s[i] = fmt.Sprint(v)
	}

	// JsonPath's return type is []interface or []interface{}
	for _, e := range values2.([]interface{}) {
		s = remove(s, e.(string))
	}

	elipsed := time.Since(start)
	t.Log(elipsed, "=============================================================")
	t.Log(s)

	start = time.Now()
	r1 := make([]string, len(values.([]interface{})))
	for i, v := range values.([]interface{}) {
		r1[i] = fmt.Sprint(v)
	}

	r2 := make([]string, len(values2.([]interface{})))
	for i, v := range values2.([]interface{}) {
		r2[i] = fmt.Sprint(v)
	}

	r3 := diffstrslice(r1, r2)

	elipsed = time.Since(start)
	t.Log(elipsed, "=============================================================")
	t.Log(r3)

}

// O(n)
func remove[T comparable](l []T, v T) []T {
	for i, e := range l {
		if e == v {
			l = append(l[:i], l[i+1:]...)
		}
	}
	return l
}

// O(1) return a - b
func diffstrslice(a, b []string) []string {
	mb := make(map[string]struct{}, len(b))
	for _, v := range b {
		mb[v] = struct{}{}
	}

	var diff []string
	for _, v := range a {
		if _, found := mb[v]; !found {
			diff = append(diff, v)
		}
	}
	return diff
}
