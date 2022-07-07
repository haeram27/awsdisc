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

func ECRDescribeRegistryCmd(cfg *aws.Config) (*ecr.DescribeRegistryOutput, error) {
	if cfg == nil || cfg.Credentials == nil {
		err := errors.New("invalid aws config: ")
		apps.Logs.Error(err)
		return nil, err
	}

	client := ecr.NewFromConfig(*cfg)
	if client == nil {
		err := errors.New("failed to initialize aws client: ")
		apps.Logs.Error(err)
		return nil, err
	}

	awsctx := context.TODO()
	input := &ecr.DescribeRegistryInput{}
	return client.DescribeRegistry(awsctx, input)
}

func ECRDescribeRepositoriesCmd(cfg *aws.Config) (*ecr.DescribeRepositoriesOutput, error) {
	if cfg == nil || cfg.Credentials == nil {
		err := errors.New("invalid aws config: ")
		apps.Logs.Error(err)
		return nil, err
	}

	client := ecr.NewFromConfig(*cfg)
	if client == nil {
		err := errors.New("failed to initialize aws client: ")
		apps.Logs.Error(err)
		return nil, err
	}

	awsctx := context.TODO()
	input := &ecr.DescribeRepositoriesInput{}
	return client.DescribeRepositories(awsctx, input)
}

func ECRListImagesCmd(cfg *aws.Config, repoName string) (*ecr.ListImagesOutput, error) {
	if cfg == nil || cfg.Credentials == nil {
		err := errors.New("invalid aws config: ")
		apps.Logs.Error(err)
		return nil, err
	}

	if repoName == "" {
		err := errors.New("invalid repository name")
		apps.Logs.Error(err)
		return nil, err
	}

	client := ecr.NewFromConfig(*cfg)
	if client == nil {
		err := errors.New("failed to initialize aws client: ")
		apps.Logs.Error(err)
		return nil, err
	}

	awsctx := context.TODO()
	input := &ecr.ListImagesInput{
		RepositoryName: &repoName,
	}

	return client.ListImages(awsctx, input)
}

func ECRListImagesAll(cfg *aws.Config) []string {
	if cfg == nil || cfg.Credentials == nil {
		err := errors.New("invalid aws config: ")
		apps.Logs.Error(err)
		return nil
	}

	var jsonBlob []byte
	result, err := ECRDescribeRepositoriesCmd(cfg)
	if err != nil {
		apps.Logs.Error(err)
		return nil
	} else {
		jsonBlob, err = json.Marshal(result)
		if err != nil {
			apps.Logs.Error(err)
			return nil
		}
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
