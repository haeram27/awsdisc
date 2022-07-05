package client

import (
	"testing"
)

func TestEcrPubDiscoverAll(t *testing.T) {
	j, err := ECRPubDescribeRegistryCmd(AwsConfig())
	if err != nil {
		t.Errorf("ECRDescribeRepositoriesCmd() failed ")
	}
	t.Log(PrettyJson(j).String())
}
