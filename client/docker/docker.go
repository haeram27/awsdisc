package dkr

import (
	"context"
	"io"
	"os"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func PullImage(url string, uri string, auth string) {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}

	// body, err := cli.RegistryLogin(ctx, types.AuthConfig{ServerAddress: url, RegistryToken: auth})
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("================== %+v", body)

	reader, err := cli.ImagePull(ctx, uri, types.ImagePullOptions{RegistryAuth: auth})
	if err != nil {
		panic(err)
	}

	defer reader.Close()
	io.Copy(os.Stdout, reader)
}
