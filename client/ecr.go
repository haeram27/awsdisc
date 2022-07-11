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

type EcrImage struct {
	repoUri *string
	tag     *string
	digest  *string
}

type EcrImageUri interface {
	TagUri() string
	DigestUri() string
}

func (img EcrImage) TagUri() string {
	return *img.repoUri + ":" + *img.tag
}

func (img EcrImage) DigestUri() string {
	return *img.repoUri + "@" + *img.digest
}

func ECRListImagesAll(cfg *aws.Config) []EcrImage {
	if cfg == nil || cfg.Credentials == nil {
		err := errors.New("invalid aws config: ")
		apps.Logs.Error(err)
		return nil
	}

	reposOut, err := ECRDescribeRepositoriesCmd(cfg)
	if err != nil {
		apps.Logs.Error(err)
		return nil
	} else {
		jsonBlob, err := json.Marshal(reposOut)
		if err != nil {
			apps.Logs.Error(err)
			return nil
		}

		awsutil.PrintPrettyJson(jsonBlob)
	}

	var images []EcrImage
	for _, repo := range reposOut.Repositories {
		imgOut, err := ECRListImagesCmd(cfg, *repo.RepositoryName)
		if err != nil {
			apps.Logs.Error(err)
			continue
		}

		for _, img := range imgOut.ImageIds {
			images = append(images, EcrImage{repo.RepositoryUri, img.ImageTag, img.ImageDigest})
		}
	}

	return images
}
