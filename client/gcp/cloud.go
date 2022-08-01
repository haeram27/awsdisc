package gcp

// [START compute_instances_list]
import (
	"context"
	"fmt"
	"io"

	compute "cloud.google.com/go/compute/apiv1"
	"google.golang.org/api/iterator"
	computepb "google.golang.org/genproto/googleapis/cloud/compute/v1"
)

// listInstances prints a list of instances created in given project in given zone.
func listInstances(w io.Writer, projectID, zone string) error {
	// projectID := "your_project_id"
	// zone := "europe-central2-b"
	ctx := context.Background()
	instancesClient, err := compute.NewInstancesRESTClient(ctx)
	if err != nil {
		return fmt.Errorf("NewInstancesRESTClient: %v", err)
	}
	defer instancesClient.Close()

	req := &computepb.ListInstancesRequest{
		Project: projectID,
		Zone:    zone,
	}

	it := instancesClient.List(ctx, req)
	fmt.Fprintf(w, "Instances found in zone %s:\n", zone)
	for {
		instance, err := it.Next()
		if err == iterator.Done {
			break
		}

		if err != nil {
			return err
		}
		fmt.Fprintf(w, "- %s %s\n", instance.GetName(), instance.GetMachineType())
	}

	return nil
}
