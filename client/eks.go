package client

import (
	apps "awsdisc/apps"
	"context"
	"encoding/json"
	"errors"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/eks"
)

func EKSDescribeClusterCmd(cfg *aws.Config, name string) (j []byte, err error) {
	if cfg == nil || cfg.Credentials == nil {
		err := errors.New("invalid aws config... ")
		apps.Logs.Error(err)
		return []byte{}, err
	}

	if name == "" {
		err := errors.New("invalid arguments: empty name")
		apps.Logs.Error(err)
		return []byte{}, err
	}

	client := eks.NewFromConfig(*cfg)
	if client == nil {
		err := errors.New("failed to initialize aws client... ")
		apps.Logs.Error(err)
		return []byte{}, err
	}

	awsctx := context.TODO()
	input := &eks.DescribeClusterInput{}
	input.Name = &name
	result, err := client.DescribeCluster(awsctx, input)
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

func EKSDescribeNodeGroupCmd(cfg *aws.Config, clusterName string, nodeGroupName string) (j []byte, err error) {
	if cfg == nil || cfg.Credentials == nil {
		err := errors.New("invalid aws config... ")
		apps.Logs.Error(err)
		return []byte{}, err
	}

	if clusterName == "" || nodeGroupName == "" {
		err := errors.New("invalid arguments: empty name")
		apps.Logs.Error(err)
		return []byte{}, err
	}

	client := eks.NewFromConfig(*cfg)
	if client == nil {
		err := errors.New("failed to initialize aws client... ")
		apps.Logs.Error(err)
		return []byte{}, err
	}

	awsctx := context.TODO()
	input := &eks.DescribeNodegroupInput{}
	input.ClusterName = &clusterName
	input.NodegroupName = &nodeGroupName
	result, err := client.DescribeNodegroup(awsctx, input)
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

/*
func EKSGetTokenCmd(cfg aws.Config, name string) (j []byte, err error) {
	if cfg == nil || cfg.Credentials == nil {
		err := errors.New("invalid aws config: ")
		apps.Logs.Error(err)
		return []byte{}, err
	}

	if name == "" {
		err := errors.New("invalid arguments: empty name")
		apps.Logs.Error(err)
		return []byte{}, err
	}

	client := eks.NewFromConfig(cfg)

	awsctx := context.TODO()
	input := &eks.DescribeClusterInput{}
	input.Name = &name
	result, err := client.(awsctx, input)
	if err != nil {
		apps.Logs.Error("got an error retrieving information about your Amazon EKS: ", err.Error())
		return []byte{}, err
	}

	resultJson, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		apps.Logs.Error("marshaling is failed:", err.Error())
		return []byte{}, err
	}

	apps.Logs.Error(string(resultJson))
	return mashalledJson, nil
}
*/

func EKSListClustersCmd(cfg *aws.Config) (j []byte, err error) {
	if cfg == nil || cfg.Credentials == nil {
		err := errors.New("invalid aws config... ")
		apps.Logs.Error(err)
		return []byte{}, err
	}

	client := eks.NewFromConfig(*cfg)
	if client == nil {
		err := errors.New("failed to initialize aws client... ")
		apps.Logs.Error(err)
		return []byte{}, err
	}

	awsctx := context.TODO()
	input := &eks.ListClustersInput{}
	result, err := client.ListClusters(awsctx, input)
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

func EKSListNodeGroupsCmd(cfg *aws.Config, name string) (j []byte, err error) {
	if cfg == nil || cfg.Credentials == nil {
		err := errors.New("invalid aws config... ")
		apps.Logs.Error(err)
		return []byte{}, err
	}

	if name == "" {
		err := errors.New("invalid arguments: empty name")
		apps.Logs.Error(err)
		return []byte{}, err
	}

	client := eks.NewFromConfig(*cfg)
	if client == nil {
		err := errors.New("failed to initialize aws client... ")
		apps.Logs.Error(err)
		return []byte{}, err
	}

	awsctx := context.TODO()
	input := &eks.ListNodegroupsInput{}
	input.ClusterName = &name
	result, err := client.ListNodegroups(awsctx, input)
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
