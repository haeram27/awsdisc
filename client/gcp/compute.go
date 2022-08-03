package gcp

// [START compute_instances_list]
import (
	"awsdisc/apps"
	"context"
	"encoding/json"
	"fmt"

	compute "cloud.google.com/go/compute/apiv1"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	computepb "google.golang.org/genproto/googleapis/cloud/compute/v1"
)

// listInstances prints a list of instances created in given project in given zone.
// $ gcloud compute instances list
func GCPAPIComputeListInstances(projectID, zone string) ([]byte, error) {
	if projectID == "" {
		projectID = "poetic-diorama-358105"
	}

	if zone == "" {
		zone = "asia-northeast3-a"
	}

	ctx := context.Background()
	client, err := compute.NewInstancesRESTClient(ctx, option.WithCredentialsFile("/home/swvm/.config/gcloud/credentials/poetic-diorama-358105-8d9d41114576.json"))
	if err != nil {
		return []byte{}, err
	}
	defer client.Close()

	req := &computepb.ListInstancesRequest{
		Project: projectID,
		Zone:    zone,
	}

	var ret []byte
	if it := client.List(ctx, req); it != nil {
		for {
			fmt.Println("max: ", it.PageInfo().MaxSize)
			fmt.Println("remain: ", it.PageInfo().Remaining())
			instance, err := it.Next()
			if err == iterator.Done {
				break
			}

			if err != nil {
				return []byte{}, err
			}

			if blob, err := json.Marshal(instance); err != nil {
				return []byte{}, err
			} else {
				ret = append(ret, blob...)
			}
		}
	}

	return ret, nil
}

func GCPAPIComputeAggregatedListInstances(projectID string) ([]byte, error) {

	ctx := context.Background()
	client, err := compute.NewInstancesRESTClient(ctx, option.WithCredentialsFile("/home/swvm/.config/gcloud/credentials/poetic-diorama-358105-8d9d41114576.json"))
	if err != nil {
		return []byte{}, err
	}
	defer client.Close()

	req := &computepb.AggregatedListInstancesRequest{
		Project: projectID,
	}

	var ret []byte
	if it := client.AggregatedList(ctx, req); it != nil {
		for {
			pair, err := it.Next()
			if err == iterator.Done {
				break
			}

			if err != nil {
				return []byte{}, err
			}

			if pair.Value.Instances != nil {
				for _, instance := range pair.Value.Instances {
					if blob, err := json.Marshal(instance); err != nil {
						return []byte{}, err
					} else {
						apps.Logs.Debug(string(blob))
						ret = append(ret, blob...)
					}
				}
			}
		}
	}

	return ret, nil
}
