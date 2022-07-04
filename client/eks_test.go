package client

import (
	"testing"
)

func TestEksDiscoveryAll(t *testing.T) {
	EKSListClustersCmd(AwsConfig())
	EKSDescribeClusterCmd(AwsConfig(), "eks-cicd-sec-test-ec2-ssh")
	EKSDescribeClusterCmd(AwsConfig(), "private-subnet-cluster")
	EKSDescribeClusterCmd(AwsConfig(), "")
}
