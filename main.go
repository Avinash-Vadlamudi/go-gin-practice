package main

import (
	"github.com/gin-gonic/gin"
)

// func v1EndPointHandler(c *gin.Context) {
// 	c.String(200, "v1: %s %s", c.Request.Method, c.Request.URL.Path)
// }

// func v2EndPointHandler(c *gin.Context) {
// 	c.String(200, "v2: %s %s", c.Request.Method, c.Request.URL.Path)
// }

// func add(c *gin.Context) {
// 	x, _ := strconv.ParseFloat(c.Param("x"), 64)
// 	y, _ := strconv.ParseFloat(c.Param("y"), 64)

// 	c.String(200, fmt.Sprintf("Sum: %f", x + y))
// }

// type AddParams struct {
// 	X float64 `json:"x"`
// 	Y float64 `json:"y"`
// }

// func add(c *gin.Context) {
// 	var ap AddParams
// 	err := c.ShouldBindJSON(&ap)
// 	if err != nil {
// 		c.JSON(400, gin.H {
// 			"error": "Calculator error",
// 			"check": err,
// 		})
// 		return
// 	}

// 	c.JSON(200, gin.H {
// 		"Sum": ap.X + ap.Y,
// 		"check2": err,
// 	})
// }

type Product struct {
	Id int `json:"id" xml:"Id" yaml:"id"`
	Name string `json:"name" xml:"Name" yaml:"name"` 
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


	// v1 := router.Group("/v1")
	// v1.GET("/products", v1EndPointHandler)


	// v2 := router.Group("/v2")
	// v2.GET("/products", v2EndPointHandler)

	// router.GET("/add/:x/:y", add)

	// router.POST("/add", add)

	router.GET("/productJSON", func(ctx *gin.Context) {
		product := Product{
			Id: 1,
			Name: "Apple",
		}

		ctx.JSON(200, product)
	})
	
	router.GET("/productXML", func(ctx *gin.Context) {
		product := Product{
			Id: 2,
			Name: "Banana",
		}

		ctx.XML(200, product)
	})
	
	router.GET("/productYAML", func(ctx *gin.Context) {
		product := Product{
			Id: 3,
			Name: "Banana",
		}

		ctx.YAML(200, product)
	})


	router.Run(":3000")
}