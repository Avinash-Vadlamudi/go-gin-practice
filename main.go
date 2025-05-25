package main

import (
	"github.com/gin-gonic/gin"
)

func v1EndPointHandler(c *gin.Context) {
	c.String(200, "v1: %s %s", c.Request.Method, c.Request.URL.Path)
}

func v2EndPointHandler(c *gin.Context) {
	c.String(200, "v2: %s %s", c.Request.Method, c.Request.URL.Path)
}


func main() {
	router := gin.Default()
	// router.GET("/hello", func (c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "Hello there!",
	// 	})
	// })

	// router .GET("/os", func (c *gin.Context) {
	// 	c.String(200, "Operating System: %s", runtime.GOOS)
	// })


	v1 := router.Group("/v1")
	v1.GET("/products", v1EndPointHandler)


	v2 := router.Group("/v2")
	v2.GET("/products", v2EndPointHandler)


	router.Run(":3000")
}