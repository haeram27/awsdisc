package client

import (
	"testing"
)

func TestEc2DiscoverAll(t *testing.T) {
	EC2DescribeInstancesCmd(AwsConfig())
}
