package service

import (
	"CleanArchitectureGo/internal/entities"
	"CleanArchitectureGo/internal/repo"
	"time"
)

type UserServiceInterface interface {
	AddUser(user *entities.User) error
	RecUser(id int) (*entities.User, error)
	RemoveUser(id int) error
	RedactUser(user *entities.User, userID int) (*entities.User, error)
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

func (u *UserService) RemoveUser(id int) error {
	return u.repo.DeleteUser(id)
}

func (u *UserService) RedactUser(user *entities.User, userID int) (*entities.User, error) {

	user.ID = userID
	user.FromDateUpdate = time.Now()

	return u.repo.UpdateUser(user)
}
