package main

import (
	"JSON_CRUD_API/controllers"
	"JSON_CRUD_API/initalizers"

	"github.com/gin-gonic/gin"
)

func init() {
	initalizers.LoadEnvvariables()
	initalizers.ConnectToDB()
}

func main() {
	r := gin.Default()
	r.POST("/", controllers.PostsCreate)
	r.GET("/posts", controllers.PostIndex)
	r.GET("/posts/:id", controllers.PostShow)
	r.PUT("/posts/:id", controllers.PostUpdate)
	r.DELETE("/posts/:id", controllers.PostDelete)
	r.Run()
}
