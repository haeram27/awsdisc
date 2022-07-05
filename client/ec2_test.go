package client

import (
	"testing"
)

func TestEc2DiscoverAll(t *testing.T) {
	j, err := EC2DescribeInstancesCmd(AwsConfig())
	if err != nil {
		t.Errorf("EC2DescribeInstancesCmd() failed ")
	}
	t.Log(PrettyJson(j).String())
}
