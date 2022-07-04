package client

import (
	"testing"
)

func TestEcrDiscoverAll(t *testing.T) {
	ECRDescribeRegistryCmd(AwsConfig())
	ECRDescribeRepositoriesCmd(AwsConfig())
}
