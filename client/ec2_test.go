package client

import (
	awsutil "awsdisc/client/util"
	"encoding/json"
	"testing"
)

func TestEc2DiscoverAll(t *testing.T) {
	var jsonBlob []byte
	var result interface{}

	result, err := EC2DescribeInstancesCmd(AwsConfig())
	if err != nil {
		t.Error(err)
	} else {
		jsonBlob, err = json.Marshal(result)
		if err != nil {
			t.Error(err)
		}
	}
	t.Log(awsutil.PrettyJson(jsonBlob).String())
}
