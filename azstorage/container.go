package azstorage

import (
	"context"
	"fmt"
	"log"
	"net/url"

	"github.com/Azure/azure-storage-blob-go/azblob"
)

// type AzureContainerPipeline struct {
// 	pl pipeline.Pipeline
// }

// func ConnectContainer() AzureStorageHandler {
// 	credential, err := azblob.NewSharedKeyCredential(accountName, accountKey)
// 	if err != nil {
// 		log.Fatal("Invalid credentials with error: " + err.Error())
// 	}
// 	p := azblob.NewPipeline(credential, azblob.PipelineOptions{})

// 	return &AzureContainerPipeline{pl: p}
// }

func CreateContainer(containerId string) bool {
	credential, err := azblob.NewSharedKeyCredential(accountName, accountKey)
	if err != nil {
		log.Fatal("Invalid credentials with error: " + err.Error())
	}
	pipel := azblob.NewPipeline(credential, azblob.PipelineOptions{})

	URL, _ := url.Parse(
		fmt.Sprintf("https://%s.blob.core.windows.net/%s", accountName, containerId))

	containerURL := azblob.NewContainerURL(*URL, pipel)
	fmt.Printf("Creating a container named %s\n", containerId)
	ctx := context.Background()
	_, err = containerURL.Create(ctx, azblob.Metadata{}, azblob.PublicAccessContainer)

	return err == nil
}

func DeleteContainer(containerId string) bool {
	credential, err := azblob.NewSharedKeyCredential(accountName, accountKey)
	if err != nil {
		log.Fatal("Invalid credentials with error: " + err.Error())
	}
	pipel := azblob.NewPipeline(credential, azblob.PipelineOptions{})

	URL, _ := url.Parse(
		fmt.Sprintf("https://%s.blob.core.windows.net/%s", accountName, containerId))

	containerURL := azblob.NewContainerURL(*URL, pipel)
	ctx := context.Background()
	_, err = containerURL.Delete(ctx, azblob.ContainerAccessConditions{})

	return err == nil
}
