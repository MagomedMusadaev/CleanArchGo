package http

import (
	"CleanArchitectureGo/internal/handler"
	"github.com/gorilla/mux"
)

func InitRoutes(r *mux.Router, handler *handler.UserHandler) {
	r.HandleFunc("/user", handler.GetUser).Methods("GET")
	r.HandleFunc("/user", handler.CreateUser).Methods("POST")
	r.HandleFunc("/user", handler.UpdateUser).Methods("PATCH")
	r.HandleFunc("/user", handler.DeleteUser).Methods("DELETE")
}
