package main

import (
    "github.com/gin-gonic/gin"
)

func main() {
		r := gin.Default()
r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run()
    router.LoadHTMLGlob("templates/*.html")

    router.GET("/", func(ctx *gin.Context){
        ctx.HTML(200, "index.html", gin.H{})
    })

    router.Run()
}