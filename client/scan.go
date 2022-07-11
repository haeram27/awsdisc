package client

import "fmt"

func ScanEcrPriv() []byte {
	images := ECRListImagesAll(AwsConfig())

	for _, img := range images {
		fmt.Println(EcrImageUri(img).TagUri())
		fmt.Println(EcrImageUri(img).DigestUri())
	}

	return []byte{}
}
