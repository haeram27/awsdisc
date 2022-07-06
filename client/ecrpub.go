package client

import (
	apps "awsdisc/apps"
	"context"
	"encoding/json"
	"errors"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ecrpublic"
)

func ECRPubDescribeRegistryCmd(cfg *aws.Config) (j []byte, err error) {
	if cfg == nil || cfg.Credentials == nil {
		err := errors.New("invalid aws config: ")
		apps.Logs.Error(err)
		return []byte{}, err
	}

	client := ecrpublic.NewFromConfig(*cfg)
	if client == nil {
		err := errors.New("failed to initialize aws client: ")
		apps.Logs.Error(err)
		return []byte{}, err
	}

	awsctx := context.TODO()
	input := &ecrpublic.DescribeRegistriesInput{}
	result, err := client.DescribeRegistries(awsctx, input)
	if err != nil {
		apps.Logs.Error(err)
		return []byte{}, err
	}

	mashalledJson, err := json.Marshal(result)
	if err != nil {
		apps.Logs.Error(err)
		return []byte{}, err
	}

	return mashalledJson, nil
}
