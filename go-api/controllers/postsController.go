package controllers

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/go-api/initializers"
	"gitlab.com/go-api/models"
)

func CreatePost(c *gin.Context) {
	var body struct {
		Body string
		Title string
	}
	c.Bind(&body)

	post := models.Post{Title: body.Title, Body: body.Body}
	result := initializers.DB.Create(&post)
	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"post": post,
	})
}

func GetALLPosts(c *gin.Context)  {

	var posts []models.Post
	initializers.DB.Find(&posts)

	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func GetPostByID(c *gin.Context)  {

	id := c.Param("id")

	var post models.Post
	initializers.DB.First(&post,id)

	c.JSON(200, gin.H{
		"post": post,
	})
}

func UpdatePost(c *gin.Context)  {

	id := c.Param("id")

	var body struct {
		Body string
		Title string
	}
	c.Bind(&body)

	var post models.Post
	initializers.DB.First(&post,id)

	initializers.DB.Model(&post).Updates(models.Post{
		Title: body.Body,
		Body: body.Body,
	})

	c.JSON(200, gin.H{
		"post": post,
	})
}

func DeletePost(c *gin.Context)  {

	id := c.Param("id")

	var post models.Post
	initializers.DB.First(&post,id)

	initializers.DB.Delete(&models.Post{}, id)

	c.JSON(200, gin.H{
		"deleted post": post,
	})
}