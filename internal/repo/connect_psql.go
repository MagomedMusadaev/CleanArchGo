package repo

import (
	"CleanArchitectureGo/pkg/logg"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"os"
)

func ConnectPostgresDB() *sql.DB {
	hostPSQL := os.Getenv("POSTGRES_HOST")
	portPSQL := os.Getenv("POSTGRES_PORT")
	userPSQL := os.Getenv("POSTGRES_USER")
	passPSQL := os.Getenv("POSTGRES_PASS")
	dbnamePSQL := os.Getenv("POSTGRES_NAME")

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		hostPSQL, portPSQL, userPSQL, passPSQL, dbnamePSQL,
	)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		logg.Error("Не удалось подключиться к db - " + err.Error())
	}

	if err := db.Ping(); err != nil {
		logg.Error("Нет коннекта с db - " + err.Error())
	}

	logg.Info("Успешно подключено к PostgresDB")

	return db
}
