package client

import (
	apps "awsdisc/apps"
	"context"
	"errors"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

func EC2DescribeInstancesCmd(cfg *aws.Config) (*ec2.DescribeInstancesOutput, error) {
	if cfg == nil || cfg.Credentials == nil {
		err := errors.New("invalid aws config")
		apps.Logs.Error(err)
		return nil, err
	}

	client := ec2.NewFromConfig(*cfg)
	if client == nil {
		err := errors.New("failed to initialize aws client")
		apps.Logs.Error(err)
		return nil, err
	}

	awsctx := context.TODO()
	input := &ec2.DescribeInstancesInput{}

	return client.DescribeInstances(awsctx, input)
}
