package service

import (
	"CleanArchitectureGo/internal/entities"
	"CleanArchitectureGo/internal/repo"
	"time"
)

type UserServiceInterface interface {
	AddUser(user *entities.User) error
	RecUser(id int) (*entities.User, error)
}

type UserService struct {
	repo repo.UserRepoInterface
}

func NewUserService(repo repo.UserRepoInterface) *UserService {
	return &UserService{repo: repo}
}

func (u *UserService) AddUser(user *entities.User) error {

	user.FromDateCreate = time.Now()
	user.FromDateUpdate = user.FromDateCreate

	return u.repo.CreateUser(user)
}

func (u *UserService) RecUser(id int) (*entities.User, error) {
	return u.repo.GetUser(id)
}
