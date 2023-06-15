package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hemang-go/go-crud/controllers"
	"github.com/hemang-go/go-crud/initializers"
)

func init() {
	initializers.ConnectToDB()
	initializers.Load_Env()
}
func main() {
	r := gin.Default()
	r.POST("/posts", controllers.PostBlog)
	r.GET("/readposts", controllers.BlogRead)
	r.GET("/readposts/:id", controllers.SingleShow)
	r.PUT("/updateposts/:id", controllers.UpdateBlog)
	r.DELETE("/delete/:id", controllers.DeleteBlog)
	r.Run()
}
