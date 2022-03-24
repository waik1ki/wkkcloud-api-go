package postgresql

import (
	"database/sql"
	"time"
)

type UserInfo struct {
	Index     int       `json:"index"`
	Sessionid string    `json:"sessionid"`
	Id        string    `json:"id"`
	Password  string    `json:"password"`
	Name      string    `json:"name"`
	Phone     string    `json:"phone"`
	Email     string    `json:"email"`
	CreateAt  time.Time `json:"created_at"`
}

type ContainerInfo struct {
	Index int    `json:"index"`
	Name  string `json:"name"`
	Id    string `json:"id"`
}

type FileInfo struct {
	Name        string    `json:"name"`
	Size        int64     `json:"size"`
	Ftype       string    `json:"ftype"`
	Author      string    `json:"author"`
	Container   string    `json:"container"`
	DownloadURL string    `json:"downloadUrl"`
	CreateAt    time.Time `json:"created_at"`
}

type CategoryInfo struct {
	Index string `json:"index"`
	Name  string `json:"name"`
	Size  int    `json:"size"`
	Count int    `json:"count"`
}

type connectDB struct {
	db *sql.DB
}

type Handler interface {
	//------------------- user -------------------//
	CreateUser(id string, password string, name string, phone string, email string) bool
	ReadUser(id string, password string) (int, UserInfo)
	UpdateUser(id string, password string) bool
	DeleteUser(id string) bool
	SessionPASS(sessionID string) (bool, UserInfo)
	//------------------- container -------------------//
	CreateContainer(containerName string, containerID string) bool
	DeleteContainer(containerName string) bool
	GetContainers() []*ContainerInfo
	FindContainer(name string) ContainerInfo
	//------------------- file -------------------//
	InsertFile(name string, size int64, ftype string, author string, container string, downloadUrl string) bool
	DeleteFile(name string) bool
	GetFiles() []*FileInfo
	GetRecentUploadFiles() []*FileInfo
	GetFileByContainer(containerName string) []*FileInfo
	//------------------- overview -------------------//
	GetOverview() []*CategoryInfo
	UpdateOverview(ftype string, size int64, option string)
	InitOverview()

	Close()
}

func NewHandler() Handler {
	return InitPostgreSQL()
}
