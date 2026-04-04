package databaseclient

import (
	"database/sql"
	"fmt"

	"github.com/anibaka9/faresniper/internal/database"
	_ "github.com/mattn/go-sqlite3"
	"github.com/pressly/goose/v3"
)

type Client struct {
	DB      *sql.DB
	Queries *database.Queries
}

func NewClient() (Client, error) {
	db, err := sql.Open("sqlite3", "./app.db")
	if err != nil {
		return Client{}, fmt.Errorf("Ошибка подключения к БД: %v", err)
	}

	// 2. Автоматический запуск миграций при старте (Goose)
	if err := goose.SetDialect("sqlite3"); err != nil {
		return Client{}, fmt.Errorf("Ошибка установки диалекта: %v", err)
	}
	if err := goose.Up(db, "sql/schema"); err != nil {
		return Client{}, fmt.Errorf("Ошибка применения миграций: %v", err)
	}

	dbQueries := database.New(db)
	c := Client{Queries: dbQueries, DB: db}
	return c, nil
}
