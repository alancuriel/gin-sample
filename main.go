package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello chika",
		})
	})

	r.LoadHTMLGlob("hello.html")

	r.GET("/hello", func(ctx *gin.Context) {
		ctx.HTML(200, "hello.html", gin.H{})
	})
	r.Run()

}
