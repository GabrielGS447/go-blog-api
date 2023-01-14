package handlers

import (
	"net/http"
	"strconv"

	"github.com/gabrielgaspar447/go-blog-api/errs"
	"github.com/gabrielgaspar447/go-blog-api/models"
	"github.com/gabrielgaspar447/go-blog-api/services"
	"github.com/gabrielgaspar447/go-blog-api/utils"
	"github.com/gin-gonic/gin"
)

func PostCreate(c *gin.Context) {
	var input models.Post

	if err := c.ShouldBindJSON(&input); err != nil {
		handlePostErrors(c, err)
		return
	}

	input.UserId = c.GetUint("userId")

	err := services.PostCreate(&input)

	if err != nil {
		handlePostErrors(c, err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": input})
}

func PostList(c *gin.Context) {
	includeUser := c.Query("user") == "true"

	posts, err := services.PostList(includeUser)

	if err != nil {
		handlePostErrors(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": posts})
}

func PostGetById(c *gin.Context) {
	includeUser := c.Query("user") == "true"
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errs.ErrInvalidId})
		return
	}

	post, err := services.PostGetById(uint(id), includeUser)

	if err != nil {
		handlePostErrors(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": post})
}

func PostSearch(c *gin.Context) {
	query := c.Query("q")
	includeUser := c.Query("user") == "true"

	posts, err := services.PostSearch(query, includeUser)

	if err != nil {
		handlePostErrors(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": posts})
}

func PostUpdate(c *gin.Context) {
	var input models.Post
	userId := c.GetUint("userId")

	if err := c.ShouldBindJSON(&input); err != nil {
		handlePostErrors(c, err)
		return
	}

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errs.ErrInvalidId})
		return
	}

	input.Id = uint(id)

	err = services.PostUpdate(&input, userId)

	if err != nil {
		handlePostErrors(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": input})
}

func PostDelete(c *gin.Context) {
	userId := c.GetUint("userId")
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errs.ErrInvalidId})
		return
	}

	err = services.PostDelete(uint(id), userId)

	if err != nil {
		handlePostErrors(c, err)
		return
	}

	c.Status(http.StatusNoContent)
}

func handlePostErrors(c *gin.Context, err error) {
	if valErrs := utils.GetValidationErrors(err); valErrs != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": valErrs})
		return
	}

	switch err {
	case errs.ErrPostNotFound:
		c.JSON(http.StatusNotFound, gin.H{"error": errs.ErrPostNotFound.Error()})
		return
	case errs.ErrPostNotOwned:
		c.JSON(http.StatusForbidden, gin.H{"error": errs.ErrPostNotOwned.Error()})
		return
	default:
		c.JSON(http.StatusInternalServerError, gin.H{"error": errs.ErrUnknown.Error()})
		return
	}
}
