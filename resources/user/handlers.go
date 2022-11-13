package user

import (
	"strconv"

	"github.com/gabrielgaspar447/go-blog-api/constants"
	"github.com/gabrielgaspar447/go-blog-api/models"
	"github.com/gabrielgaspar447/go-blog-api/utils"
	"github.com/gin-gonic/gin"
)

func signupHandler(c *gin.Context) {
	var input models.User

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(constants.HTTP_BadRequest, gin.H{"error": utils.GetErrorResponse(err)})
		return
	}

	token, err := signupService(&input)

	if err != nil {
		statusCode, msg := utils.GetServiceErrorResponse(err)
		c.JSON(statusCode, gin.H{"error": msg})
		return
	}

	c.JSON(constants.HTTP_Created, gin.H{"data": token})
}

func loginHandler(c *gin.Context) {
	var input models.LoginDTO

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(constants.HTTP_InternalServerError, gin.H{"error": constants.SomethingWentWrong})
		return
	}

	token, err := loginService(&input)

	if err != nil {
		statusCode, msg := utils.GetServiceErrorResponse(err)
		c.JSON(statusCode, gin.H{"error": msg})
		return
	}

	c.JSON(constants.HTTP_OK, gin.H{"data": token})
}

func listUsersHandler(c *gin.Context) {
	includePosts := c.Query("posts") == "true"

	var users []models.User

	err := listUsersService(&users, includePosts)

	if err != nil {
		statusCode, msg := utils.GetServiceErrorResponse(err)
		c.JSON(statusCode, gin.H{"error": msg})
		return
	}

	c.JSON(constants.HTTP_OK, gin.H{"data": users})
}

func getUserByIdHandler(c *gin.Context) {
	includePosts := c.Query("posts") == "true"
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(constants.HTTP_BadRequest, gin.H{"error": constants.InvalidId})
		return
	}

	var user models.User

	err = getUserByIdService(&user, uint(id), includePosts)

	if err != nil {
		statusCode, msg := utils.GetServiceErrorResponse(err)
		c.JSON(statusCode, gin.H{"error": msg})
		return
	}

	c.JSON(constants.HTTP_OK, gin.H{"data": user})
}

func deleteUserHandler(c *gin.Context) {
	id := c.GetUint("userId")

	err := deleteUserByIdService(id)

	if err != nil {
		statusCode, msg := utils.GetServiceErrorResponse(err)
		c.JSON(statusCode, gin.H{"error": msg})
		return
	}

	c.Status(constants.HTTP_NoContent)
}
