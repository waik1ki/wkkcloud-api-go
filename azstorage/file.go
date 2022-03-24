package azstorage

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"os"

	"github.com/Azure/azure-storage-blob-go/azblob"
)

func UploadFile(containerId string, fileName string) string {
	credential, err := azblob.NewSharedKeyCredential(accountName, accountKey)
	if err != nil {
		log.Fatal("Invalid credentials with error: " + err.Error())
	}
	pipel := azblob.NewPipeline(credential, azblob.PipelineOptions{})

	URL, _ := url.Parse(
		fmt.Sprintf("https://%s.blob.core.windows.net/%s", accountName, containerId))

	containerURL := azblob.NewContainerURL(*URL, pipel)
	ctx := context.Background()

	// 더미파일 생성 [ 업로드 테스트 용 ]
	// fmt.Printf("Creating a dummy file to test the upload and download\n")
	// data := []byte("hello world this is a blob\n")
	// err = ioutil.WriteFile(fileName, data, 0700)
	// if err != nil {
	// 	panic(err)
	// }

	blobURL := containerURL.NewBlockBlobURL(fileName)
	file, err := os.Open("./" + fileName)
	fmt.Println(blobURL)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Uploading the file with blob name: %s\n", fileName)
	_, err = azblob.UploadFileToBlockBlob(ctx, file, blobURL, azblob.UploadToBlockBlobOptions{
		BlockSize:   4 * 1024 * 1024,
		Parallelism: 16})

	if err != nil {
		return ""
	}

	return blobURL.String()
}

func GetFiles(containerId string) []string {
	credential, err := azblob.NewSharedKeyCredential(accountName, accountKey)
	if err != nil {
		log.Fatal("Invalid credentials with error: " + err.Error())
	}
	pipel := azblob.NewPipeline(credential, azblob.PipelineOptions{})

	URL, _ := url.Parse(
		fmt.Sprintf("https://%s.blob.core.windows.net/%s", accountName, containerId))

	containerURL := azblob.NewContainerURL(*URL, pipel)
	ctx := context.Background()

	list := []string{}

	fmt.Println("Listing the blobs in the container:")
	for marker := (azblob.Marker{}); marker.NotDone(); {
		listBlob, err := containerURL.ListBlobsFlatSegment(ctx, marker, azblob.ListBlobsSegmentOptions{})
		if err != nil {
			panic(err)
		}

		marker = listBlob.NextMarker

		for _, blobInfo := range listBlob.Segment.BlobItems {
			list = append(list, blobInfo.Name)
			fmt.Println(blobInfo.Properties.CreationTime)
			fmt.Println(blobInfo.Properties)
			fmt.Print("    Blob name: " + blobInfo.Name + "\n")
		}
	}
	return list
}

func DownloadFile(containerId string, fileName string) string {
	credential, err := azblob.NewSharedKeyCredential(accountName, accountKey)
	if err != nil {
		log.Fatal("Invalid credentials with error: " + err.Error())
	}
	pipel := azblob.NewPipeline(credential, azblob.PipelineOptions{})

	URL, _ := url.Parse(
		fmt.Sprintf("https://%s.blob.core.windows.net/%s", accountName, containerId))

	containerURL := azblob.NewContainerURL(*URL, pipel)

	blobURL := containerURL.NewBlobURL(fileName).String()

	if err != nil {
		panic(err)
	}

	return blobURL
}
func DeleteFile(containerId string, fileName string) bool {
	credential, err := azblob.NewSharedKeyCredential(accountName, accountKey)
	if err != nil {
		log.Fatal("Invalid credentials with error: " + err.Error())
	}
	pipel := azblob.NewPipeline(credential, azblob.PipelineOptions{})

	URL, _ := url.Parse(
		fmt.Sprintf("https://%s.blob.core.windows.net/%s", accountName, containerId))

	containerURL := azblob.NewContainerURL(*URL, pipel)
	ctx := context.Background()
	blobURL := containerURL.NewBlobURL(fileName)

	_, err = blobURL.Delete(ctx, "", azblob.BlobAccessConditions{})

	return err == nil
}
