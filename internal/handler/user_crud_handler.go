package handler

import (
	"CleanArchitectureGo/internal/entities"
	"CleanArchitectureGo/internal/service"
	"CleanArchitectureGo/pkg/logg"
	"CleanArchitectureGo/pkg/utils"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
)

type UserHandlerInterface interface {
	CreateUser(w http.ResponseWriter, r *http.Request)
	GetUser(w http.ResponseWriter, r *http.Request)
}

type UserHandler struct {
	service service.UserServiceInterface
}

func NewUserService(service service.UserServiceInterface) *UserHandler {
	return &UserHandler{service: service}
}

func (s *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user entities.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		utils.DecodeErr(w, err, http.StatusBadRequest)
		logg.Error("Не удалось декодить полученные данные - " + err.Error())
		return
	}

	if err := s.service.AddUser(&user); err != nil {
		utils.DecodeErr(w, err, http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("user " + user.Name + " created"))
}

func (s *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {

	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		utils.DecodeErr(w, errors.New("ID пользователя не указан"), http.StatusBadRequest)
		return
	}

	userID, err := strconv.Atoi(idStr)
	if err != nil {
		utils.DecodeErr(w, errors.New("uncorrected param"), http.StatusBadRequest)
		logg.Error("Невалидный параметр - " + err.Error())
		return
	}

	user, err := s.service.RecUser(userID)
	if err != nil {
		utils.DecodeErr(w, err, http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}
