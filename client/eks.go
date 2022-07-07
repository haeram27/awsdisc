package client

import (
	apps "awsdisc/apps"
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/eks"
	"github.com/aws/aws-sdk-go-v2/service/eks/types"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/aws-iam-authenticator/pkg/token"
)

func EKSDescribeClusterSTCmd(cfg *aws.Config, name string) (*eks.DescribeClusterOutput, error) {
	if cfg == nil || cfg.Credentials == nil {
		err := errors.New("invalid aws config... ")
		apps.Logs.Error(err)
		return nil, err
	}

	if name == "" {
		err := errors.New("invalid arguments: empty name")
		apps.Logs.Error(err)
		return nil, err
	}

	client := eks.NewFromConfig(*cfg)
	if client == nil {
		err := errors.New("failed to initialize aws client... ")
		apps.Logs.Error(err)
		return nil, err
	}

	awsctx := context.TODO()
	input := &eks.DescribeClusterInput{}
	input.Name = &name
	return client.DescribeCluster(awsctx, input)
}

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

/*
  Resolve k8s config for k8s sdk from eks cluster information
*/
func EKSK8sClientset(cluster *types.Cluster) (*kubernetes.Clientset, error) {
	gen, err := token.NewGenerator(true, false)
	if err != nil {
		apps.Logs.Error(err)
		return nil, err
	}

	opts := &token.GetTokenOptions{
		ClusterID: *cluster.Name,
	}

	tok, err := gen.GetWithOptions(opts)
	if err != nil {
		apps.Logs.Error(err)
		return nil, err
	}

	ca, err := base64.StdEncoding.DecodeString(*cluster.CertificateAuthority.Data)
	if err != nil {
		apps.Logs.Error(err)
		return nil, err
	}

	k8sConfig := rest.Config{
		Host:        *cluster.Endpoint,
		BearerToken: tok.Token,
		TLSClientConfig: rest.TLSClientConfig{
			CAData: ca,
		},
	}

	k8sclientset, err := kubernetes.NewForConfig(&k8sConfig)
	if err != nil {
		apps.Logs.Error(err)
		return nil, err
	}

	return k8sclientset, nil
}
