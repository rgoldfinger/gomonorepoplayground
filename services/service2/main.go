package main

import (
	"fmt"
	"net/http"

	"example.com/monorepo/pkg/base"
	"example.com/monorepo/pkg/package1"
	"github.com/gin-gonic/gin"
)


func main() {
	fmt.Println("service 2")
	fmt.Println(package1.RevString("fooooo"))
	r := base.Init("service1")
	r.GET("/ping", func(c *gin.Context) {
	  c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	  })
	})
	r.Run()
}
