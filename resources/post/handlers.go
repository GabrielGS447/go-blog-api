package post

import (
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
