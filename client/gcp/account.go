package gcp

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/storage"
	"golang.org/x/oauth2/google"
	cloudkms "google.golang.org/api/cloudkms/v1"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

// [START auth_cloud_implicit]

// implicit uses Application Default Credentials to authenticate.
func implicit() {
	ctx := context.Background()

	// For API packages whose import path is starting with "cloud.google.com/go",
	// such as cloud.google.com/go/storage in this case, if there are no credentials
	// provided, the client library will look for credentials in the environment.
	storageClient, err := storage.NewClient(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer storageClient.Close()

	it := storageClient.Buckets(ctx, "project-id")
	for {
		bucketAttrs, err := it.Next()
		if err == iterator.Done {
			break
		}

		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(bucketAttrs.Name)
	}

	// For packages whose import path is starting with "google.golang.org/api",
	// such as google.golang.org/api/cloudkms/v1, use NewService to create the client.
	kmsService, err := cloudkms.NewService(ctx)
	if err != nil {
		log.Fatal(err)
	}

	_ = kmsService
}

// [END auth_cloud_implicit]

// [START auth_cloud_explicit]

// explicit reads credentials from the specified path.
func explicit(jsonPath, projectID string) {
	ctx := context.Background()
	client, err := storage.NewClient(ctx, option.WithCredentialsFile(jsonPath))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	fmt.Println("Buckets:")
	it := client.Buckets(ctx, projectID)
	for {
		battrs, err := it.Next()
		if err == iterator.Done {
			break
		}

		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(battrs.Name)
	}
}

// [END auth_cloud_explicit]

// [START auth_cloud_explicit_compute_engine]
// [START auth_cloud_explicit_app_engine]

// OAuth2 scopes used by this API.
const (
	// View and manage your data across Google Cloud Platform services
	CloudPlatformScope = "https://www.googleapis.com/auth/cloud-platform"

	// View your data across Google Cloud Platform services
	CloudPlatformReadOnlyScope = "https://www.googleapis.com/auth/cloud-platform.read-only"
)

// explicitDefault finds the default credentials.
//
// It is very uncommon to need to explicitly get the default credentials in Go.
// Most of the time, client libraries can use Application Default Credentials
// without having to pass the credentials in directly. See implicit above.
func explicitDefault(projectID string) {
	ctx := context.Background()

	creds, err := google.FindDefaultCredentials(ctx, CloudPlatformReadOnlyScope)
	if err != nil {
		log.Fatal(err)
	}

	client, err := storage.NewClient(ctx, option.WithCredentials(creds))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	fmt.Println("Buckets:")
	it := client.Buckets(ctx, projectID)
	for {
		battrs, err := it.Next()
		if err == iterator.Done {
			break
		}

		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(battrs.Name)
	}
}

// [END auth_cloud_explicit_compute_engine]
// [END auth_cloud_explicit_app_engine]
