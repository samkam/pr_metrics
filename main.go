package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"io/ioutil"
)
import "fmt"
func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/listen", func(c *gin.Context) {
		c.JSON(200, gin.H{ "message": "Hello world",})

		fmt.Println("hello")
	})
	r.POST("/listen", listen)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
func listen(c *gin.Context) {
	fmt.Println("I show up")
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	fmt.Printf("%s",body)
	fmt.Println()

	c.JSON(http.StatusOK, gin.H{"message":"hi"})
	return
}