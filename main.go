package main

import (
	"log"
	"net/http"

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

// type Product struct {
// 	Id int `json:"id" xml:"Id" yaml:"id"`
// 	Name string `json:"name" xml:"Name" yaml:"name"`
// }

// type PrintJob struct {
// 	JobId int `json:"jobId" binding:"required,gte=10000"`
// 	Pages int `json:"pages" binding:"required,gte=1,lte=100"`
// }


func FindUserAgent() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println(c.GetHeader("User-Agent"))
	
		c.Next()
	}
}

func CookieTool() gin.HandlerFunc {
	return func(c *gin.Context) {
		if cookie, err := c.Cookie("my_cookie"); err == nil {
			if cookie == "ok" {
				c.Next()
				return
			}
		}

		c.JSON(http.StatusForbidden, gin.H {
			"error" : "Forbidden with no cookie",
		})

		c.Abort() // Stop further processing
	}
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

	// router.GET("/productJSON", func(ctx *gin.Context) {
	// 	product := Product{
	// 		Id: 1,
	// 		Name: "Apple",
	// 	}

	// 	ctx.JSON(200, product)
	// })
	
	// router.GET("/productXML", func(ctx *gin.Context) {
	// 	product := Product{
	// 		Id: 2,
	// 		Name: "Banana",
	// 	}

	// 	ctx.XML(200, product)
	// })
	
	// router.GET("/productYAML", func(ctx *gin.Context) {
	// 	product := Product{
	// 		Id: 3,
	// 		Name: "Banana",
	// 	}

	// 	ctx.YAML(200, product)
	// })

	// router.POST("/print", func(c *gin.Context) {
	// 	var p PrintJob
	// 	if err := c.ShouldBindJSON(&p); err != nil {
	// 		c.JSON(400, gin.H{
	// 			"error": "Invalid print job",
	// 			"details": err.Error(),
	// 		})
	// 		return
	// 	}
	// 	c.JSON(200, gin.H{
	// 		"message": fmt.Sprintf("Print job #%v started!", p.JobId),
	// 	})
	// })



	// router.Use(cors.Default())

	// router.GET("/", func (c* gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "CORS works!",
	// 	})
	// })


	// router.Use(FindUserAgent())

	// router.GET("/", func (c *gin.Context) {
	// 	c.JSON(200, gin.H {
	// 		"message": "Middleware works!",
	// 	})
	// })

	router.GET("/login", func (c *gin.Context) {
		c.SetCookie("my_cookie", "ok", 300, "/", "localhost", false, true)

		c.String(200, "Login successful! Cookie set." +
			" You can now access protected routes.")
	})

	router.GET("/home", CookieTool(), func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to the home page!",
		})
	})



	router.Run(":3000")
}