package post

import (
	"strconv"

	"github.com/gabrielgaspar447/go-blog-api/constants"
	"github.com/gabrielgaspar447/go-blog-api/models"
	"github.com/gabrielgaspar447/go-blog-api/utils"
	"github.com/gin-gonic/gin"
)

func createPostHandler(c *gin.Context) {
	var input models.Post

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(constants.HTTP_BadRequest, gin.H{"error": utils.GetErrorResponse(err)})
		return
	}

	input.UserID = c.GetUint("userId")

	err := createPostService(&input)

	if err != nil {
		statusCode, msg := utils.GetServiceErrorResponse(err)
		c.JSON(statusCode, gin.H{"error": msg})
		return
	}

	c.JSON(constants.HTTP_Created, gin.H{"data": input})
}

func listPostsHandler(c *gin.Context) {
	includeUser := c.Query("user") == "true"

	var posts []models.Post

	err := listPostsService(&posts, includeUser)

	if err != nil {
		statusCode, msg := utils.GetServiceErrorResponse(err)
		c.JSON(statusCode, gin.H{"error": msg})
		return
	}

	c.JSON(constants.HTTP_OK, gin.H{"data": posts})
}

func getPostByIdHandler(c *gin.Context) {
	includeUser := c.Query("user") == "true"
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(constants.HTTP_BadRequest, gin.H{"error": constants.InvalidId})
		return
	}

	var post models.Post

	err = getPostByIdService(&post, uint(id), includeUser)

	if err != nil {
		statusCode, msg := utils.GetServiceErrorResponse(err)
		c.JSON(statusCode, gin.H{"error": msg})
		return
	}

	c.JSON(constants.HTTP_OK, gin.H{"data": post})
}
