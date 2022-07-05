package client

import (
	"awsdisc/apps"
	"context"
	"encoding/json"
	"errors"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

func EC2DescribeInstancesCmd(cfg *aws.Config) (j []byte, err error) {
	if cfg == nil || cfg.Credentials == nil {
		err := errors.New("invalid aws config: ")
		apps.Logs.Error(err.Error())
		return []byte{}, err
	}

	client := ec2.NewFromConfig(*cfg)
	if client == nil {
		err := errors.New("failed to initialize aws client: ")
		apps.Logs.Error(err.Error())
		return []byte{}, err
	}

	awsctx := context.TODO()
	input := &ec2.DescribeInstancesInput{}
	result, err := client.DescribeInstances(awsctx, input)
	if err != nil {
		apps.Logs.Error("got an error retrieving information about your Amazon EC2: ", err.Error())
		return []byte{}, err
	}

	mashalledJson, err := json.Marshal(result)
	if err != nil {
		apps.Logs.Error("marshaling is failed:", err.Error())
		return []byte{}, err
	}

	// DEBUG
	PrintPrettyJson(mashalledJson)

	return mashalledJson, nil
}
