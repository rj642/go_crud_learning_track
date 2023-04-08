package controllers

import (
	"example.com/m/v2/initializers"
	"example.com/m/v2/models"
	"github.com/gin-gonic/gin"
)

func PostsCreate(c *gin.Context) {
	// Get data off req body
	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)

	// Create a post
	post := models.Post{Title: body.Title, Body: body.Body}

	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}

	// Return it
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsIndex(c *gin.Context) {
	// Get the posts
	var posts []models.Post
	initializers.DB.Find(&posts)

	// Respond with them
	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func GetPostById(c *gin.Context) {
	// Get id off url
	id := c.Param("id")

	// Get the post
	var post models.Post
	initializers.DB.Find(&post, id)

	// Return the result
	c.JSON(200, gin.H{
		"post": post,
	})

}

func UpdatePost(c *gin.Context) {
	// Get the id of the url
	id := c.Param("id")

	// Get the request from the body
	var body struct {
		Title string
		Body  string
	}

	c.Bind(&body)

	// Find the post were updating
	var post models.Post
	initializers.DB.Find(&post, id)

	// Update it
	initializers.DB.Model(&post).Updates(models.Post{
		Title: body.Title,
		Body:  body.Body,
	})

	// Send response to user

	c.JSON(200, gin.H{
		"post": post,
	})

}

func DeletePostById(c *gin.Context) {
	// Get the post id to delete
	id := c.Param("id")

	// Check if post exists or not
	var post models.Post
	initializers.DB.Find(&post, id)

	if post.DeletedAt.Valid {
		// Respond
		c.JSON(404, gin.H{
			"success": false,
		})
	} else {
		// Delete the post from database
		initializers.DB.Delete(&models.Post{}, id)

		// Respond
		c.Status(200)
	}
}
