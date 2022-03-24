package handler

import (
	"waikiki/wkkcloud/postgresql"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization, Origin")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, DELETE, POST")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func MakeHandler() *gin.Engine {
	router := gin.Default()

	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("auth_session", store))
	router.Use(CORSMiddleware())

	h := &HandlerStruct{
		pqHandler: postgresql.NewHandler(),
	}

	user := router.Group("/user")
	{
		user.POST("/register", h.CreateUserHandler)
		user.DELETE("/delete", h.DeleteUserHandler)
		user.PUT("/update", h.UpdateUserHandler)
		user.POST("/login", h.LoginHandler)
	}
	file := router.Group("/file")
	{
		file.GET("/get", h.GetFilesHandler)
		file.GET("/get/recent", h.GetRecentFilesHandler)

		file.POST("/get/files", h.GetFileByContainerHandler)
		file.POST("/upload", h.UploadFileHandler)
		file.POST("/delete", h.DeleteFileHandler)

		file.POST("/download", h.DownloadBlobHandler)
		file.POST("/download/recent", h.DownloadRecentBlobHandler)
	}
	container := router.Group("/container")
	{
		container.POST("/create", h.CreateContainerHandler)
		container.POST("/delete", h.DeleteContainerHandler)
		container.GET("/get", h.GetContainersHandler)
	}
	overview := router.Group("/overview")
	{
		overview.GET("/get", h.GetOverviewHandler)
		overview.POST("/init", h.InitOverviewHandler)
	}
	router.GET("/auth", h.SessionTestHandler)

	return router
}
