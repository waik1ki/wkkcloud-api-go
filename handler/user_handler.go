package handler

import (
	"net/http"

	"waikiki/wkkcloud/postgresql"

	"github.com/gin-gonic/gin"
)

func (hs *HandlerStruct) LoginHandler(c *gin.Context) {
	var data postgresql.UserInfo
	c.ShouldBindJSON(&data)
	status, user := hs.pqHandler.ReadUser(data.Id, data.Password)

	switch status {
	case 200:
		c.JSON(http.StatusOK, gin.H{"user": user, "token": "abcd"})
		return
	case 422:
		c.JSON(http.StatusUnprocessableEntity, gin.H{"err": "아이디 혹은 비밀번호를 확인해주세요."})
		return
	}

	// if rst {
	// 	session := sessions.Default(c)
	// 	// str := session.ID()
	// 	session.Set("sessionID", sID)
	// 	session.Save()
	// 	// hs.pqHandler.SessionPASS(data.Id, "test")

	// 	// c.JSON(200, gin.H{"id": data.Id})
	// } else {
	// 	if sID == "not found" {
	// 		c.JSON(http.StatusOK, gin.H{"status": "not found"})
	// 		return
	// 	} else if sID == "passwords do not match" {
	// 		c.JSON(http.StatusOK, gin.H{"status": "not match"})
	// 		return
	// 	}
	// }
}

func (hs *HandlerStruct) CreateUserHandler(c *gin.Context) {
	var data postgresql.UserInfo
	c.ShouldBindJSON(&data)
	rst := hs.pqHandler.CreateUser(data.Id, data.Password, data.Name, data.Phone, data.Email)
	if rst {
		c.JSON(http.StatusOK, gin.H{"status": "created"})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"status": "failed create"})
		return
	}
}

func (hs *HandlerStruct) DeleteUserHandler(c *gin.Context) {
	var data postgresql.UserInfo
	c.ShouldBindJSON(&data)
	rst := hs.pqHandler.DeleteUser(data.Id)
	if rst {
		c.JSON(http.StatusOK, gin.H{"status": "deleted"})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"status": "failed delete"})
		return
	}
}

func (hs *HandlerStruct) UpdateUserHandler(c *gin.Context) {
	var data postgresql.UserInfo
	c.ShouldBindJSON(&data)
	rst := hs.pqHandler.UpdateUser(data.Id, data.Password)
	if rst {
		c.JSON(http.StatusOK, gin.H{"status": "updated"})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"status": "failed update"})
		return
	}
}
