package repo

import (
	"CleanArchitectureGo/internal/entities"
	"CleanArchitectureGo/pkg/logg"
	"database/sql"
	"errors"
	"fmt"
	"github.com/lib/pq"
	"reflect"
	"strings"
	"time"
)

type UserRepoInterface interface {
	CreateUser(user *entities.User) error
	GetUser(id int) (*entities.User, error)
	DeleteUser(id int) error
	UpdateUser(user *entities.User) (*entities.User, error)
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

		return nil, errors.New("ошибка получения пользователя")
	}

	return &user, nil
}

func (u *UserRepository) DeleteUser(id int) error {
	query := `UPDATE users SET is_deleted = true WHERE id = $1`

	_, err := u.db.Exec(query, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return errors.New("пользователя с таким id не найдено")
		}

		return errors.New("ошибка удаления пользователя")
	}

	return nil
}

func (u *UserRepository) UpdateUser(user *entities.User) (*entities.User, error) {
	var query strings.Builder
	var values []interface{}

	query.WriteString(`UPDATE users SET `)

	v := reflect.ValueOf(*user)
	t := v.Type()

	// Формирование запроса
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		tag := t.Field(i).Tag.Get("db")

		if tag == "" || tag == "id" {
			continue // если понадобиться менять id пользователя, то мы просто уберём это условие
		}

		isEmpty := false
		switch field.Kind() {
		case reflect.String:
			isEmpty = field.String() == ""
		case reflect.Bool:
			if !field.IsZero() {
				query.WriteString(fmt.Sprintf("%s = $%d, ", tag, len(values)+1))
				values = append(values, field.Bool())
			}
		case reflect.Struct:
			if field.Type() == reflect.TypeOf(time.Time{}) {
				timeValue := field.Interface().(time.Time)
				// Проверка на пустое значение времени
				if !timeValue.IsZero() {
					values = append(values, timeValue) // Явно передаем time.Time
					query.WriteString(fmt.Sprintf("%s = $%d, ", tag, len(values)))
				}
				continue
			}
		}

		// Если поле не пустое, добавляем его в запрос
		if !isEmpty {
			query.WriteString(fmt.Sprintf("%s = $%d, ", tag, len(values)+1))
			values = append(values, field.Interface())
		}
	}

	// Проверка на наличие полей для обновления
	if len(values) == 0 {
		return nil, errors.New("no fields to update")
	}

	// Формирование финального запроса
	finalQuery := strings.TrimSuffix(query.String(), ", ")
	finalQuery += fmt.Sprintf(" WHERE id = $%d RETURNING *", len(values)+1)
	values = append(values, user.ID)

	// Выполнение запроса
	row := u.db.QueryRow(finalQuery, values...)
	if err := row.Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.FromDateCreate,
		&user.FromDateUpdate,
		&user.IsDeleted,
		&user.IsBanned,
	); err != nil {
		return nil, errors.New("не удалось изменить данные: " + err.Error())
	}

	return user, nil
}
