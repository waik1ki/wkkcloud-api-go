package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (hs *HandlerStruct) GetOverviewHandler(c *gin.Context) {
	overview := hs.pqHandler.GetOverview()
	c.JSON(http.StatusOK, gin.H{"overview": overview})
}

func (hs *HandlerStruct) InitOverviewHandler(c *gin.Context) {
	hs.pqHandler.InitOverview()
	c.JSON(http.StatusOK, gin.H{"status": "initalize"})
}
