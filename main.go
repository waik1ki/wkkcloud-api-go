package main

import (
	"os"
	"waikiki/wkkcloud/handler"
)

func main() {
	port := os.Getenv("PORT")
	r := handler.MakeHandler()

	err := r.Run(":" + port)
	if err != nil {
		panic(err)
	}
}
