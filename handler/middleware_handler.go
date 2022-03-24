package handler

import (
	"fmt"
	"net/http"

	"waikiki/wkkcloud/postgresql"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func (hs *HandlerStruct) SessionTestHandler(c *gin.Context) {
	session := sessions.Default(c)
	s := session.Get("sessionID")

	if s == nil {
		c.JSON(http.StatusOK, gin.H{"status": "failed"})
		return
	}
	var user postgresql.UserInfo

	rst, user := hs.pqHandler.SessionPASS(s.(string))
	if rst {
		c.JSON(http.StatusOK, gin.H{"status": "pass", "user": user})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"status": "failed"})
		return
	}
}

func DummyMiddleware(c *gin.Context) {
	fmt.Println("Im a dummy!")

	c.Next()
}
