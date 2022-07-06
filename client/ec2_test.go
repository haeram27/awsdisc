package client

import (
	awsutil "awsdisc/client/util"
	"testing"
)

func TestEc2DiscoverAll(t *testing.T) {
	j, err := EC2DescribeInstancesCmd(AwsConfig())
	if err != nil {
		t.Error(err)
	}
	t.Log(awsutil.PrettyJson(j).String())
}
