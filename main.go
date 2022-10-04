package main

import (
	"expentracker/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.GET("/", handler.Login)
	router.Run()
}
