package app

import (
	"atnlie/logger"
	"fmt"

	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApplication() {
	//map all router
	mapUrls()

	//init logger
	logger.Info("about to start the user-api service")

	router.Run(":8080")
}

func StartApplication2() {
	fmt.Println("=======================================")
	fmt.Println("Starting Service")
	fmt.Println("=======================================")

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "test",
		})
	})
	fmt.Println("=======================================")
	router.Run()

	/*
		router.GET("/someGet", getting)
		router.POST("/somePost", posting)
		router.PUT("/somePut", putting)
		router.DELETE("/someDelete", deleting)
		router.PATCH("/somePatch", patching)
		router.HEAD("/someHead", head)
		router.OPTIONS("/someOptions", options)
	*/
}