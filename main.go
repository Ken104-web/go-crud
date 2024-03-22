package main

import (
	initializers "go-crud/init"
	"github.com/gin-gonic/gin"
)

func init(){
  initializers.LoadInitVar()
  initializers.ConnectToDo()

}

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() 
}


