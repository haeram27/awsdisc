package client

import (
	"awsdisc/apps"
	"encoding/json"
)

func DiscoverAll() ([]byte, error) {
	// EC2
	ec2, err := EC2DescribeInstancesCmd(AwsConfig())
	if err != nil {
		ec2 = []byte{}
	}

	// ECR
	ecr, err := ECRDescribeRepositoriesCmd(AwsConfig())
	if err != nil {
		ecr = []byte{}
	}

	// ECS
	//ECSListClustersCmd(AwsConfig())
	ecs, err := ECSDescribeClustersCmd(AwsConfig(), []string{"cicd-ecs-ec2-cluster", "swh-ecs-cluster-ssh", "cicd-ecs-cluster"})
	if err != nil {
		ecs = []byte{}
	}

	// EKS
	eks, err := EKSListClustersCmd(AwsConfig())
	if err != nil {
		eks = []byte{}
	}

	// consolidate
	data := make(map[string]interface{})
	var inner interface{}
	json.Unmarshal(ec2, &inner)
	data["EC2"] = inner
	json.Unmarshal(ecr, &inner)
	data["ECRRepositories"] = inner
	json.Unmarshal(ecs, &inner)
	data["ECSClusters"] = inner
	json.Unmarshal(eks, &inner)
	data["EKSClusters"] = inner

	result, err := json.Marshal(data)
	if err != nil {
		apps.Logs.Error("could not marshal json: ", err)
		return []byte{}, err
	}

	return result, nil
}
