package main

import (
	"os"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	//export Name="OlafurArnalds"
	name := os.Getenv("name")
	country := os.Getenv("country")
	//resultfmt.Sprintf("Name:%s Country:%s", name, country)

	r.GET("/hello", func(c *gin.Context) {
		c.String(200, "Hello, World!"+" Name:%s ", name, " Country:%s", country)
	})

	api := r.Group("/api")

	api.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.Use(static.Serve("/", static.LocalFile("./views", true)))
	//Setting Port
	r.Run(":4444")
}
