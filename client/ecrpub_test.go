package client

import (
	awsutil "awsdisc/client/util"
	"testing"
)

func TestEcrPubDiscoverAll(t *testing.T) {
	j, err := ECRPubDescribeRegistryCmd(AwsConfig())
	if err != nil {
		t.Error(err)
	}
	t.Log(awsutil.PrettyJson(j).String())
}
