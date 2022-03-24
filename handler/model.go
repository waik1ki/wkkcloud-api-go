package handler

import (
	"mime/multipart"

	"waikiki/wkkcloud/postgresql"
)

type HandlerStruct struct {
	pqHandler postgresql.Handler
	// app2 upload.AzureStorageHandler
}

type Files struct {
	ContainerId string   `json:"containerId"`
	File        []string `json:"file"`
}

var fm map[int][]*multipart.FileHeader
