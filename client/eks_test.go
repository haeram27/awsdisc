package client

import (
	awsutil "awsdisc/client/util"
	"testing"
)

func TestEksDiscoveryAll(t *testing.T) {
	j, err := EKSListClustersCmd(AwsConfig())
	if err != nil {
		t.Error(err)
	}
	t.Log(awsutil.PrettyJson(j).String())

	j, err = EKSDescribeClusterCmd(AwsConfig(), "eks-cicd-sec-test-ec2-ssh")
	if err != nil {
		t.Error(err)
	}
	t.Log(awsutil.PrettyJson(j).String())

	j, err = EKSDescribeClusterCmd(AwsConfig(), "private-subnet-cluster")
	if err != nil {
		t.Error(err)
	}
	t.Log(awsutil.PrettyJson(j).String())

	/*
		j, err = EKSDescribeClusterCmd(AwsConfig(), "")
		if err != nil {
			t.Error(err)
		}
		t.Log(awsutil.PrettyJson(j).String())
	*/
}
