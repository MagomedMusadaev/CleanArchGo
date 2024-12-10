package repo

import (
	"CleanArchitectureGo/internal/entities"
	"CleanArchitectureGo/pkg/logg"
	"database/sql"
	"errors"
	"fmt"
	"github.com/lib/pq"
)

type UserRepoInterface interface {
	CreateUser(user *entities.User) error
	GetUser(id int) (*entities.User, error)
}

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (u *UserRepository) CreateUser(user *entities.User) error {

	query := `INSERT INTO users (name, email, password, from_date_create, from_date_update) VALUES ($1, $2, $3, $4, $5)`

	_, err := u.db.Exec(query,
		user.Name,
		user.Email,
		user.Password,
		user.FromDateCreate,
		user.FromDateUpdate,
	)

	if err != nil {
		var pqErr *pq.Error
		// Проверка, является ли ошибка ошибкой PostgreSQL
		if errors.As(err, &pqErr) {
			// Проверка на уникальность email
			if pqErr.Code == "23505" {
				return errors.New("пользователь с таким Email уже существует")
			}

			// Если ошибка не связана с уникальностью, возвращаем её как есть
			logg.Error("Не удалось добавить пользователя в psql: " + err.Error())
			return errors.New("не удалось создать пользователя")
		}
	}

	return nil
}

func (u *UserRepository) GetUser(id int) (*entities.User, error) {
	var user entities.User

	query := `SELECT name, email, password, from_date_create, from_date_update FROM users WHERE id = $1`

	err := u.db.QueryRow(query, id).Scan(
		&user.Name,
		&user.Email,
		&user.Password,
		&user.FromDateCreate,
		&user.FromDateUpdate,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("пользователя с таким id не найдено")
		}
		fmt.Println(err)
		return nil, errors.New("ошибка получения пользователя")
	}

	return &user, nil
}
