package client

import (
	"awsdisc/apps"
	"awsdisc/client/util"
	"encoding/json"
)

func DiscoverAll() {
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
	var result interface{}
	json.Unmarshal(ec2, &result)
	data["EC2"] = result
	json.Unmarshal(ecr, &result)
	data["ECRRepositories"] = result
	json.Unmarshal(ecs, &result)
	data["ECSClusters"] = result
	json.Unmarshal(eks, &result)
	data["EKSClusters"] = result

	jsonData, err := json.Marshal(data)
	if err != nil {
		apps.Logs.Error("could not marshal json: %s\n", err)
		return
	}

	util.PrintPrettyJson(jsonData)
}
