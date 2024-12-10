package entities

import "time"

type User struct {
	ID             int       `json:"id"`
	Name           string    `json:"name"`
	Email          string    `json:"email"`
	Password       string    `json:"password"`
	FromDateCreate time.Time `json:"fromDateCreate"`
	FromDateUpdate time.Time `json:"fromDateUpdate"`
	IsDeleted      bool      `json:"isDeleted"`
	IsBanned       bool      `json:"isBanned"`
}
