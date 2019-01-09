package main

import (
	"github.com/gin-gonic/gin"
	"github.com/my3/cloud-native-go/ping"

	)

func main() {
	r := gin.Default()
	r.GET("/ping", ping.RespondToPing)
	r.Run()
}
