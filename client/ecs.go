package client

import (
	"awsdisc/apps"
	"awsdisc/client/util"
	"context"
	"encoding/json"
	"errors"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ecs"
)

func ECSDescribeClustersCmd(cfg *aws.Config, name []string) (j []byte, err error) {
	if cfg == nil || cfg.Credentials == nil {
		err := errors.New("invalid aws config: ")
		apps.Logs.Error(err.Error())
		return []byte{}, err
	}

	if len(name) == 0 {
		err := errors.New("invalid arguments: empty name")
		apps.Logs.Error(err.Error())
		return []byte{}, err
	}

	client := ecs.NewFromConfig(*cfg)
	if client == nil {
		err := errors.New("failed to initialize aws client: ")
		apps.Logs.Error(err.Error())
		return []byte{}, err
	}

	awsctx := context.TODO()
	input := &ecs.DescribeClustersInput{}
	input.Clusters = name
	result, err := client.DescribeClusters(awsctx, input)
	if err != nil {
		apps.Logs.Error("got an error retrieving information about your Amazon ECS: ", err.Error())
		return []byte{}, err
	}

	mashalledJson, err := json.Marshal(result)
	if err != nil {
		apps.Logs.Error("marshaling is failed:", err.Error())
		return []byte{}, err
	}

	// DEBUG
	util.PrintPrettyJson(mashalledJson)

	return mashalledJson, nil
}

func ECSDescribeTaskDefinitionCmd(cfg *aws.Config, task string) (j []byte, err error) {
	if cfg == nil || cfg.Credentials == nil {
		err := errors.New("invalid aws config: ")
		apps.Logs.Error(err.Error())
		return []byte{}, err
	}

	if task == "" {
		err := errors.New("invalid arguments: empty name")
		apps.Logs.Error(err.Error())
		return []byte{}, err
	}

	client := ecs.NewFromConfig(*cfg)
	if client == nil {
		err := errors.New("failed to initialize aws client: ")
		apps.Logs.Error(err.Error())
		return []byte{}, err
	}

	awsctx := context.TODO()
	input := &ecs.DescribeTaskDefinitionInput{}
	input.TaskDefinition = &task
	result, err := client.DescribeTaskDefinition(awsctx, input)
	if err != nil {
		apps.Logs.Error("got an error retrieving information about your Amazon ECS: ", err.Error())
		return []byte{}, err
	}

	mashalledJson, err := json.Marshal(result)
	if err != nil {
		apps.Logs.Error("marshaling is failed:", err.Error())
		return []byte{}, err
	}

	// DEBUG
	util.PrintPrettyJson(mashalledJson)

	return mashalledJson, nil
}

func ECSListClustersCmd(cfg *aws.Config) (j []byte, err error) {
	if cfg == nil || cfg.Credentials == nil {
		err := errors.New("invalid aws config: ")
		apps.Logs.Error(err.Error())
		return []byte{}, err
	}

	client := ecs.NewFromConfig(*cfg)
	if client == nil {
		err := errors.New("failed to initialize aws client: ")
		apps.Logs.Error(err.Error())
		return []byte{}, err
	}

	awsctx := context.TODO()
	input := &ecs.ListClustersInput{}
	result, err := client.ListClusters(awsctx, input)
	if err != nil {
		apps.Logs.Error("got an error retrieving information about your Amazon ECS: ", err.Error())
		return []byte{}, err
	}

	mashalledJson, err := json.Marshal(result)
	if err != nil {
		apps.Logs.Error("marshaling is failed:", err.Error())
		return []byte{}, err
	}

	// DEBUG
	util.PrintPrettyJson(mashalledJson)

	return mashalledJson, nil
}

func ECSListContainerInstancesCmd(cfg *aws.Config, name string) (j []byte, err error) {
	if cfg == nil || cfg.Credentials == nil {
		err := errors.New("invalid aws config: ")
		apps.Logs.Error(err.Error())
		return []byte{}, err
	}

	if name == "" {
		err := errors.New("invalid arguments: empty name")
		apps.Logs.Error(err.Error())
		return []byte{}, err
	}

	client := ecs.NewFromConfig(*cfg)
	if client == nil {
		err := errors.New("failed to initialize aws client: ")
		apps.Logs.Error(err.Error())
		return []byte{}, err
	}

	awsctx := context.TODO()
	input := &ecs.ListContainerInstancesInput{}
	input.Cluster = &name
	result, err := client.ListContainerInstances(awsctx, input)
	if err != nil {
		apps.Logs.Error("got an error retrieving information about your Amazon ECS: ", err.Error())
		return []byte{}, err
	}

	mashalledJson, err := json.Marshal(result)
	if err != nil {
		apps.Logs.Error("marshaling is failed:", err.Error())
		return []byte{}, err
	}

	// DEBUG
	util.PrintPrettyJson(mashalledJson)

	return mashalledJson, nil
}

func ECSListTaskDefinitionsCmd(cfg *aws.Config) (j []byte, err error) {
	if cfg == nil || cfg.Credentials == nil {
		err := errors.New("invalid aws config: ")
		apps.Logs.Error(err.Error())
		return []byte{}, err
	}

	client := ecs.NewFromConfig(*cfg)
	if client == nil {
		err := errors.New("failed to initialize aws client: ")
		apps.Logs.Error(err.Error())
		return []byte{}, err
	}

	awsctx := context.TODO()
	input := &ecs.ListTaskDefinitionsInput{}
	result, err := client.ListTaskDefinitions(awsctx, input)
	if err != nil {
		apps.Logs.Error("got an error retrieving information about your Amazon ECS: ", err.Error())
		return []byte{}, err
	}

	mashalledJson, err := json.Marshal(result)
	if err != nil {
		apps.Logs.Error("marshaling is failed:", err.Error())
		return []byte{}, err
	}

	// DEBUG
	util.PrintPrettyJson(mashalledJson)

	return mashalledJson, nil
}
