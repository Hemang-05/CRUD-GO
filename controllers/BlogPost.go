package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/hemang-go/go-crud/initializers"
	"github.com/hemang-go/go-crud/models"
)

func PostBlog(c *gin.Context) {
	//Get data  off req body
	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)

	//create post
	post := models.BlogPost{Title: body.Title, Body: body.Body}

	result := initializers.DB.Create(&post)
	// return

	if result.Error != nil {
		c.Status(404)
		return
	}

	c.JSON(200, gin.H{
		"post": post,
	})
}

func BlogRead(c *gin.Context) {
	// Get all records
	var result []models.BlogPost

	initializers.DB.Find(&result)

	c.JSON(200, gin.H{
		"post": result,
	})
}

func SingleShow(c *gin.Context) {
	// get id from url
	id := c.Param("id")

	// Get  record by id
	var result models.BlogPost

	initializers.DB.First(&result, id)

	c.JSON(200, gin.H{
		"post": result,
	})
}

func UpdateBlog(c *gin.Context) {

	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)

	// find  id
	id := c.Param("id")

	//   get data to request post
	// post := models.BlogPost{Title: body.Title, Body: body.Body}
	// result_ := initializers.DB.Create(&post)

	// find by id
	var result models.BlogPost
	initializers.DB.First(&result, id)
	//update
	initializers.DB.First(&result)

	result.Title = "jinzhu 2"
	result.Body = "100"
	initializers.DB.Save(&result)

	// respond
	c.JSON(200, gin.H{
		"post": result,
	})
}

func DeleteBlog(c *gin.Context) {
	// get id
	id := c.Param("id")

	// delete by id
	var result models.BlogPost
	initializers.DB.Delete(&result, id)

	// send response
	c.JSON(200, gin.H{
		"post": result,
	})

}
