package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gabrielgs447/go-blog-api/errs"
	"github.com/gabrielgs447/go-blog-api/models"
	"github.com/gabrielgs447/go-blog-api/services"
	"github.com/gabrielgs447/go-blog-api/utils"
	"github.com/gin-gonic/gin"
)

type UserHandlerInterface interface {
	Signup(c *gin.Context)
	Login(c *gin.Context)
	List(c *gin.Context)
	GetById(c *gin.Context)
	DeleteSelf(c *gin.Context)
}

type userHandler struct {
	userService services.UserServiceInterface
}

func NewUserHandler(s services.UserServiceInterface) UserHandlerInterface {
	return &userHandler{
		s,
	}
}

func (h *userHandler) Signup(c *gin.Context) {
	var input models.CreateUserDTO

	if err := c.ShouldBindJSON(&input); err != nil {
		handleUsersErrors(c, err)
		return
	}

	token, err := h.userService.Signup(c.Request.Context(), input.ToModel())

	if err != nil {
		handleUsersErrors(c, err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": token})
}

func (h *userHandler) Login(c *gin.Context) {
	var input models.LoginDTO

	if err := c.ShouldBindJSON(&input); err != nil {
		handleUsersErrors(c, err)
		return
	}

	token, err := h.userService.Login(c.Request.Context(), &input)

	if err != nil {
		handleUsersErrors(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": token})
}

func (h *userHandler) List(c *gin.Context) {
	includePosts := c.Query("posts") == "true"

	users, err := h.userService.List(c.Request.Context(), includePosts)

	if err != nil {
		handleUsersErrors(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": users})
}

func (h *userHandler) GetById(c *gin.Context) {
	includePosts := c.Query("posts") == "true"
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		handleUsersErrors(c, errs.ErrInvalidId)
		return
	}

	user, err := h.userService.GetById(c.Request.Context(), uint(id), includePosts)

	if err != nil {
		handleUsersErrors(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}

func (h *userHandler) DeleteSelf(c *gin.Context) {
	id := c.GetUint("userId")

	err := h.userService.DeleteSelf(c.Request.Context(), id)

	if err != nil {
		handleUsersErrors(c, err)
		return
	}

	c.Status(http.StatusNoContent)
}

func handleUsersErrors(c *gin.Context, err error) {
	if valErrs := utils.GetValidationErrors(err); valErrs != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": valErrs})
		return
	}

	switch {
	case errors.Is(err, errs.ErrUserAlreadyExists):
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	case errors.Is(err, errs.ErrInvalidId):
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	case errors.Is(err, errs.ErrUserNotFound):
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	case errors.Is(err, errs.ErrInvalidCredentials):
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	default:
		c.JSON(http.StatusInternalServerError, gin.H{"error": errs.ErrUnknown.Error()})
		return
	}
}
