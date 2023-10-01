package main

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/go-api/controllers"
	"gitlab.com/go-api/initializers"
)

func init()  {
initializers.LoadEnvVariables()
initializers.ConnectToDB()
}

func main() {
	r := gin.Default()
	r.POST("/posts", controllers.CreatePost)
	r.GET("/posts", controllers.GetALLPosts)
	r.GET("/posts/:id", controllers.GetPostByID)
	r.PUT("/posts/:id", controllers.UpdatePost)
	r.DELETE("/posts/:id", controllers.DeletePost)
	r.Run()

}