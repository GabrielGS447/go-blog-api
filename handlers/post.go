package handlers

import (
	"net/http"
	"strconv"

	"github.com/gabrielgs447/go-blog-api/errs"
	"github.com/gabrielgs447/go-blog-api/models"
	"github.com/gabrielgs447/go-blog-api/services"
	"github.com/gabrielgs447/go-blog-api/utils"
	"github.com/gin-gonic/gin"
)

type PostHandlerInterface interface {
	Create(c *gin.Context)
	List(c *gin.Context)
	GetById(c *gin.Context)
	Search(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type postHandler struct {
	postService services.PostServiceInterface
}

func NewPostHandler(s services.PostServiceInterface) PostHandlerInterface {
	return &postHandler{
		s,
	}
}

func (h *postHandler) Create(c *gin.Context) {
	var input models.Post

	if err := c.ShouldBindJSON(&input); err != nil {
		handlePostsErrors(c, err)
		return
	}

	input.UserId = c.GetUint("userId")

	err := h.postService.Create(c.Request.Context(), &input)

	if err != nil {
		handlePostsErrors(c, err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": input})
}

func (h *postHandler) List(c *gin.Context) {
	includeUser := c.Query("user") == "true"

	posts, err := h.postService.List(c.Request.Context(), includeUser)

	if err != nil {
		handlePostsErrors(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": posts})
}

func (h *postHandler) GetById(c *gin.Context) {
	includeUser := c.Query("user") == "true"
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		handlePostsErrors(c, errs.ErrInvalidId)
		return
	}

	post, err := h.postService.GetById(c.Request.Context(), uint(id), includeUser)

	if err != nil {
		handlePostsErrors(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": post})
}

func (h *postHandler) Search(c *gin.Context) {
	query := c.Query("q")
	includeUser := c.Query("user") == "true"

	posts, err := h.postService.Search(c.Request.Context(), query, includeUser)

	if err != nil {
		handlePostsErrors(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": posts})
}

func (h *postHandler) Update(c *gin.Context) {
	var input models.Post
	userId := c.GetUint("userId")

	if err := c.ShouldBindJSON(&input); err != nil {
		handlePostsErrors(c, err)
		return
	}

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		handlePostsErrors(c, errs.ErrInvalidId)
		return
	}

	input.Id = uint(id)

	err = h.postService.Update(c.Request.Context(), &input, userId)

	if err != nil {
		handlePostsErrors(c, err)
		return
	}

	c.Status(http.StatusNoContent)
}

func (h *postHandler) Delete(c *gin.Context) {
	userId := c.GetUint("userId")
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		handlePostsErrors(c, errs.ErrInvalidId)
		return
	}

	err = h.postService.Delete(c.Request.Context(), uint(id), userId)

	if err != nil {
		handlePostsErrors(c, err)
		return
	}

	c.Status(http.StatusNoContent)
}

func handlePostsErrors(c *gin.Context, err error) {
	if valErrs := utils.GetValidationErrors(err); valErrs != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": valErrs})
		return
	}

	switch err {
	case errs.ErrInvalidId:
		c.JSON(http.StatusBadRequest, gin.H{"error": errs.ErrInvalidId.Error()})
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
