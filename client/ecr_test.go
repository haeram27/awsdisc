package client

import (
	"testing"
)

func TestEcrDiscoverAll(t *testing.T) {
	j, err := ECRDescribeRegistryCmd(AwsConfig())
	if err != nil {
		t.Errorf("ECRDescribeRegistryCmd() failed ")
	}
	t.Log(PrettyJson(j).String())

	j, err = ECRDescribeRepositoriesCmd(AwsConfig())
	if err != nil {
		t.Errorf("ECRDescribeRepositoriesCmd() failed ")
	}
	t.Log(PrettyJson(j).String())

}
