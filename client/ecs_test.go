package client

import (
	"testing"
)

func TestEcsDiscoverAll(t *testing.T) {
	j, err := ECSListClustersCmd(AwsConfig())
	if err != nil {
		t.Errorf("ECRDescribeRepositoriesCmd() failed ")
	}
	t.Log(PrettyJson(j).String())

	j, err = ECSDescribeClustersCmd(AwsConfig(), []string{"cicd-ecs-ec2-cluster", "swh-ecs-cluster-ssh", "cicd-ecs-cluster"})
	if err != nil {
		t.Errorf("ECRDescribeRepositoriesCmd() failed ")
	}
	t.Log(PrettyJson(j).String())

	j, err = ECSListContainerInstancesCmd(AwsConfig(), "cicd-ecs-ec2-cluster")
	if err != nil {
		t.Errorf("ECRDescribeRepositoriesCmd() failed ")
	}
	t.Log(PrettyJson(j).String())

	j, err = ECSListContainerInstancesCmd(AwsConfig(), "swh-ecs-cluster-ssh")
	if err != nil {
		t.Errorf("ECRDescribeRepositoriesCmd() failed ")
	}
	t.Log(PrettyJson(j).String())

	j, err = ECSListContainerInstancesCmd(AwsConfig(), "cicd-ecs-cluster")
	if err != nil {
		t.Errorf("ECRDescribeRepositoriesCmd() failed ")
	}
	t.Log(PrettyJson(j).String())

	j, err = ECSListTaskDefinitionsCmd(AwsConfig())
	if err != nil {
		t.Errorf("ECRDescribeRepositoriesCmd() failed ")
	}
	t.Log(PrettyJson(j).String())

	j, err = ECSDescribeTaskDefinitionCmd(AwsConfig(), "cicd-task-nginx:1")
	if err != nil {
		t.Errorf("ECRDescribeRepositoriesCmd() failed ")
	}
	t.Log(PrettyJson(j).String())

	j, err = ECSDescribeTaskDefinitionCmd(AwsConfig(), "cicd-task-ubuntu_nginx:2")
	if err != nil {
		t.Errorf("ECRDescribeRepositoriesCmd() failed ")
	}
	t.Log(PrettyJson(j).String())

	j, err = ECSDescribeTaskDefinitionCmd(AwsConfig(), "sw-task:4")
	if err != nil {
		t.Errorf("ECRDescribeRepositoriesCmd() failed ")
	}
	t.Log(PrettyJson(j).String())
}
