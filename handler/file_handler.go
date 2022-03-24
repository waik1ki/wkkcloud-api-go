package handler

import (
	"fmt"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"
	"strings"

	"waikiki/wkkcloud/azstorage"
	"waikiki/wkkcloud/postgresql"

	"github.com/gin-gonic/gin"
)

func ExtractionFiletype(file string) string {
	dotIndex := strings.LastIndex(file, ".")
	ftype := file[dotIndex+1:]

	return ftype
}

func (hs *HandlerStruct) DeleteFileHandler(c *gin.Context) {
	var data Files
	c.ShouldBindJSON(&data)

	// fmt.Println(data.File)

	for _, f := range data.File {
		if rst := azstorage.DeleteFile(data.ContainerId, f); rst {
			hs.pqHandler.DeleteFile(f)
		}
	}

	c.JSON(http.StatusOK, gin.H{"respText": "파일 삭제 완료!"})
}

func (hs *HandlerStruct) UploadFileHandler(c *gin.Context) {
	formData, _ := c.MultipartForm()
	c.ShouldBind(&formData)

	fileCnt := formData.Value["fileCnt"]
	count, _ := strconv.Atoi(fileCnt[0])

	fm = make(map[int][]*multipart.FileHeader)

	for i := 0; i < count; i++ {
		fn := strconv.Itoa(i)
		fm[i] = formData.File["file"+fn]
	}

	author := formData.Value["author"]
	container := formData.Value["container"]

	for _, v := range fm {
		c.SaveUploadedFile(v[0], v[0].Filename)
		downloadURL := azstorage.UploadFile(container[0], v[0].Filename)
		if downloadURL != "" {
			if rst := hs.pqHandler.InsertFile(v[0].Filename, v[0].Size, ExtractionFiletype(v[0].Filename), author[0], container[0], downloadURL); rst {
				hs.pqHandler.UpdateOverview(ExtractionFiletype(v[0].Filename), v[0].Size, "INSERT")
				// fmt.Println(v[0].Filename, v[0].Size, ExtractionFiletype(v[0].Filename), author[0], container[0])
			}
			os.Remove("./" + v[0].Filename)
		}
	}

	c.JSON(http.StatusOK, gin.H{"respText": "파일 업로드 완료!"})
}

func (hs *HandlerStruct) GetFilesHandler(c *gin.Context) {
	files := hs.pqHandler.GetFiles()
	c.JSON(http.StatusOK, gin.H{"files": files})
}

func (hs *HandlerStruct) GetRecentFilesHandler(c *gin.Context) {
	files := hs.pqHandler.GetRecentUploadFiles()
	c.JSON(http.StatusOK, gin.H{"files": files})
}

func (hs *HandlerStruct) GetFileByContainerHandler(c *gin.Context) {
	var container postgresql.ContainerInfo
	c.ShouldBindJSON(&container)

	files := hs.pqHandler.GetFileByContainer(container.Id)
	c.JSON(http.StatusOK, gin.H{"files": files})
}

// 사용 안함 [ 다운로드 방식 수정 (프론트에서 직접 처리) ]

func (hs *HandlerStruct) DownloadBlobHandler(c *gin.Context) {
	var data Files
	var downloadUrlList []string
	c.ShouldBindJSON(&data)

	files := hs.pqHandler.GetRecentUploadFiles()
	fmt.Println(len(files))
	fmt.Println(data.File)
	for _, f := range data.File {
		downloadURL := azstorage.DownloadFile(data.ContainerId, f)
		downloadUrlList = append(downloadUrlList, downloadURL)
	}
	c.JSON(http.StatusOK, gin.H{"status": downloadUrlList})
}

func (hs *HandlerStruct) DownloadRecentBlobHandler(c *gin.Context) {
	var data Files
	var downloadUrlList []string
	c.ShouldBindJSON(&data)

	files := hs.pqHandler.GetRecentUploadFiles()
	fmt.Println(data.File)
	for _, f := range data.File {
		for _, v := range files {
			if f == v.Name {
				downloadURL := azstorage.DownloadFile(v.Container, f)
				downloadUrlList = append(downloadUrlList, downloadURL)
			}
		}
	}
	c.JSON(http.StatusOK, gin.H{"status": downloadUrlList})
}
