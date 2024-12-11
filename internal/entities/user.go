package entities

import "time"

type User struct {
	ID             int       `json:"id" db:"id"`
	Name           string    `json:"name" db:"name"`
	Email          string    `json:"email" db:"email"`
	Password       string    `json:"password" db:"password"`
	FromDateCreate time.Time `json:"fromDateCreate" db:"from_date_create"`
	FromDateUpdate time.Time `json:"fromDateUpdate" db:"from_date_update"`
	IsDeleted      bool      `json:"isDeleted" db:"is_deleted"`
	IsBanned       bool      `json:"isBanned" db:"is_banned"`
}
