package controllers

import (
	"JSON_CRUD_API/initalizers"
	models "JSON_CRUD_API/model"

	"github.com/gin-gonic/gin"
)

func PostsCreate(c *gin.Context) {

	// Get off body data

	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)

	//Create Post

	post := models.Post{Title: body.Title, Body: body.Body}
	result := initalizers.DB.Create(&post)
	if result.Error != nil {
		c.Status(400)
		return
	}

	//Return It

	c.JSON(200, gin.H{
		"Post": post,
	})
}

func PostIndex(c *gin.Context) {
	//Get the post
	var posts []models.Post
	initalizers.DB.Find(&posts)
	//respond with them
	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func PostShow(c *gin.Context) {
	//Get id off url
	id := c.Param("id")

	//Get the post
	var post []models.Post
	initalizers.DB.First(&post, id)
	//respond with them
	c.JSON(200, gin.H{
		"posts": post,
	})
}

func PostUpdate(c *gin.Context) {
	//Get the id of the url
	id := c.Param("id")

	//Get the data off the req body
	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)

	//find the post we want to update
	var post []models.Post
	initalizers.DB.First(&post, id)
	//update it
	initalizers.DB.Model(&post).Updates(models.Post{Title: body.Title, Body: body.Body})
	//Respond with thm
	c.JSON(200, gin.H{
		"posts": post,
	})
}

func PostDelete(c *gin.Context) {

	id := c.Param("id")
	var post []models.Post
	initalizers.DB.Delete(&post, id)
	c.JSON(200, gin.H{
		"posts": post,
	})
}
