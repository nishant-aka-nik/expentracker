package main

import (
	"expentracker/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.GET("/", handler.Index)
	router.POST("/login", handler.Login)

	router.Run(":5555")
}
