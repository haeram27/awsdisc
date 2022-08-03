package gcp

// [START compute_instances_list]
import (
	"awsdisc/apps/util"
	"testing"
)

func TestGCPComputeListInstances(t *testing.T) {
	if blob, err := GCPAPIComputeListInstances("", ""); err != nil {
		t.Error(err)
	} else {
		t.Log(util.PrettyJson(blob))
	}
}

func TestGCPAPIComputeAggregatedListInstances(t *testing.T) {
	if blob, err := GCPAPIComputeAggregatedListInstances("poetic-diorama-358105"); err != nil {
		t.Error(err)
	} else {
		t.Log(util.PrettyJson(blob))
	}
}
