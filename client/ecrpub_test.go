package client

import (
	"testing"
)

func TestEcrPubDiscoverAll(t *testing.T) {
	ECRPubDescribeRegistryCmd(AwsConfig())
}
