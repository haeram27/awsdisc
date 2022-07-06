package client

import (
	apps "awsdisc/apps"
	awsutil "awsdisc/client/util"
	"context"
	"encoding/json"
	"errors"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ecr"
)

func ECRDescribeRegistryCmd(cfg *aws.Config) (j []byte, err error) {
	if cfg == nil || cfg.Credentials == nil {
		err := errors.New("invalid aws config: ")
		apps.Logs.Error(err)
		return []byte{}, err
	}

	client := ecr.NewFromConfig(*cfg)
	if client == nil {
		err := errors.New("failed to initialize aws client: ")
		apps.Logs.Error(err)
		return []byte{}, err
	}

	awsctx := context.TODO()
	input := &ecr.DescribeRegistryInput{}
	result, err := client.DescribeRegistry(awsctx, input)
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

func ECRDescribeRepositoriesCmd(cfg *aws.Config) (j []byte, err error) {
	if cfg == nil || cfg.Credentials == nil {
		err := errors.New("invalid aws config: ")
		apps.Logs.Error(err)
		return []byte{}, err
	}

	client := ecr.NewFromConfig(*cfg)
	if client == nil {
		err := errors.New("failed to initialize aws client: ")
		apps.Logs.Error(err)
		return []byte{}, err
	}

	awsctx := context.TODO()
	input := &ecr.DescribeRepositoriesInput{}
	result, err := client.DescribeRepositories(awsctx, input)
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

func ECRListImagesCmd(cfg *aws.Config, repoName string) ([]byte, error) {
	if cfg == nil || cfg.Credentials == nil {
		err := errors.New("invalid aws config: ")
		apps.Logs.Error(err)
		return []byte{}, err
	}

	if repoName == "" {
		err := errors.New("invalid repository name")
		apps.Logs.Error(err)
		return []byte{}, err
	}

	client := ecr.NewFromConfig(*cfg)
	if client == nil {
		err := errors.New("failed to initialize aws client: ")
		apps.Logs.Error(err)
		return []byte{}, err
	}

	awsctx := context.TODO()
	input := &ecr.ListImagesInput{
		RepositoryName: &repoName,
	}

	result, err := client.ListImages(awsctx, input)
	if err != nil {
		apps.Logs.Error(err)
		return []byte{}, err
	}

	return json.Marshal(result)
}

func ECRListImagesAll(cfg *aws.Config) []string {
	if cfg == nil || cfg.Credentials == nil {
		err := errors.New("invalid aws config: ")
		apps.Logs.Error(err)
		return nil
	}

	jsonBlob, err := ECRDescribeRepositoriesCmd(cfg)
	if err != nil {
		return nil
	}

	names, err := awsutil.JsonPath(jsonBlob, "$.Repositories[:].RepositoryName")
	if err != nil {
		apps.Logs.Error(err)
		return nil
	}

	repos := []interface{}{}
	switch names.(type) {
	case []interface{}:
		repos = names.([]interface{})
	case interface{}:
		repos = append(repos, names)
	default:
		return nil
	}

	for _, info := range repos {
		apps.Logs.Debug("============================== repository name: ", info.(string))
		ECRListImagesCmd(cfg, info.(string))
	}

	return nil
}
