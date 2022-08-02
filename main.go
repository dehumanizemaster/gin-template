package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Get Example
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// Post Example
	/*	POST /post?id=1234&page=1 HTTP/1.1
		Content-Type: application/x-www-form-urlencoded

		name=manu&message=this_is_great
	*/
	r.POST("/live", func(c *gin.Context) {
		//id := c.Query("id")
		streamkey := c.PostForm("streamkey")
		fmt.Printf(" streamkey is %s", streamkey)
	})

	r.GET("/root", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "root",
		})
	})

	//r.Run("127.0.0.1:10144") // defualt listen and serve on 0.0.0.0:8080
	r.Run("0.0.0.0:8080")
}
