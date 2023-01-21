package services

import (
	"context"

	"github.com/gabrielgaspar447/go-blog-api/auth"
	"github.com/gabrielgaspar447/go-blog-api/database"
	"github.com/gabrielgaspar447/go-blog-api/errs"
	"github.com/gabrielgaspar447/go-blog-api/models"
	"golang.org/x/crypto/bcrypt"
)

type UserServiceInterface interface {
	Signup(ctx context.Context, input *models.User) (string, error)
	Login(ctx context.Context, input *models.LoginDTO) (string, error)
	List(ctx context.Context, includePosts bool) (*[]models.User, error)
	GetById(ctx context.Context, id uint, includePosts bool) (*models.User, error)
	DeleteSelf(ctx context.Context, id uint) error
}

type userService struct {
	userRepository database.UserRepositoryInterface
}

func NewUserService(r database.UserRepositoryInterface) UserServiceInterface {
	return &userService{
		r,
	}
}

func (s *userService) Signup(ctx context.Context, input *models.User) (string, error) {
	user, err := s.userRepository.FindByEmail(ctx, input.Email)
	if err != nil {
		return "", err
	}

	if user.Id != 0 {
		return "", errs.ErrUserAlreadyExists
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(input.Password), 8)
	if err != nil {
		return "", err
	}

	input.Password = string(hash)

	err = s.userRepository.Create(ctx, input)
	if err != nil {
		return "", err
	}

	return auth.SignJWT(input.Id)
}

func (s *userService) Login(ctx context.Context, input *models.LoginDTO) (string, error) {
	user, err := s.userRepository.FindByEmail(ctx, input.Email)
	if err != nil {
		return "", err
	}

	if user.Id == 0 {
		return "", errs.ErrUserNotFound
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil {
		return "", errs.ErrInvalidPassword
	}

	return auth.SignJWT(user.Id)
}

func (s *userService) List(ctx context.Context, includePosts bool) (*[]models.User, error) {
	users, err := s.userRepository.List(ctx, includePosts)
	if err != nil {
		return nil, err
	}

	for _, user := range *users {
		user.SanitizeToJson()
	}

	return users, nil
}

func (s *userService) GetById(ctx context.Context, id uint, includePosts bool) (*models.User, error) {
	user, err := s.userRepository.GetById(ctx, id, includePosts)
	if err != nil {
		return nil, err
	}

	if user.Id == 0 {
		return nil, errs.ErrUserNotFound
	}

	user.SanitizeToJson()

	return user, nil
}

func (s *userService) DeleteSelf(ctx context.Context, id uint) error {
	return s.userRepository.DeleteById(ctx, id)
}
