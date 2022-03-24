package handler

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"waikiki/wkkcloud/azstorage"
	"waikiki/wkkcloud/postgresql"

	"github.com/gin-gonic/gin"
)

func randomString() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return strconv.Itoa(r.Int())
}

func (hs *HandlerStruct) GetContainersHandler(c *gin.Context) {
	containers := hs.pqHandler.GetContainers()
	c.JSON(http.StatusOK, gin.H{"containers": containers})
}

func (hs *HandlerStruct) CreateContainerHandler(c *gin.Context) {
	containerId := randomString()

	var container postgresql.ContainerInfo
	c.ShouldBindJSON(&container)

	fmt.Println(containerId, container.Name)
	rst := azstorage.CreateContainer(containerId)

	if rst {
		if hs.pqHandler.CreateContainer(container.Name, containerId); true {
			c.JSON(http.StatusOK, gin.H{"respText": "컨테이너 생성 완료!"})
			return
		}
	}
}

func (hs *HandlerStruct) DeleteContainerHandler(c *gin.Context) {

	var container postgresql.ContainerInfo
	c.ShouldBindJSON(&container)

	rst := azstorage.DeleteContainer(container.Id)

	if rst {
		if hs.pqHandler.DeleteContainer(container.Id); true {
			c.JSON(http.StatusOK, gin.H{"respText": "컨테이너 삭제 완료!"})
			return
		}
	}

}
