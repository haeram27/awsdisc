package client

import (
	"testing"
)

func TestEcsDiscoverAll(t *testing.T) {
	ECSListClustersCmd(AwsConfig())
	ECSDescribeClustersCmd(AwsConfig(), []string{"cicd-ecs-ec2-cluster", "swh-ecs-cluster-ssh", "cicd-ecs-cluster"})
	ECSListContainerInstancesCmd(AwsConfig(), "cicd-ecs-ec2-cluster")
	ECSListContainerInstancesCmd(AwsConfig(), "swh-ecs-cluster-ssh")
	ECSListContainerInstancesCmd(AwsConfig(), "cicd-ecs-cluster")
	ECSListTaskDefinitionsCmd(AwsConfig())
	ECSDescribeTaskDefinitionCmd(AwsConfig(), "cicd-task-nginx:1")
	ECSDescribeTaskDefinitionCmd(AwsConfig(), "cicd-task-ubuntu_nginx:2")
	ECSDescribeTaskDefinitionCmd(AwsConfig(), "sw-task:4")
}
