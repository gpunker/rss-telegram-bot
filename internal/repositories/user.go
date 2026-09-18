package repositories

import (
	"database/sql"
	"fmt"
	"strings"

	types "github.com/gpunker/rss-telegram-bot/internal"
)

type UserParams struct {
	ID         *int
	TelegramID *int64
	UserName   *string
}

func FindUser(filter UserParams) (types.User, error) {
	db := GetConnection()

	// TODO сделать простенький генератор запросов
	query := "SELECT * FROM users"

	var conditions []string
	var args []any

	if filter.ID != nil {
		conditions = append(conditions, "id = ?")
		args = append(args, *filter.ID)
	}

	if filter.TelegramID != nil {
		conditions = append(conditions, "telegram_id = ?")
		args = append(args, *filter.TelegramID)
	}

	if filter.UserName != nil {
		conditions = append(conditions, "username = ?")
		args = append(args, *filter.UserName)
	}

	if len(conditions) > 0 {
		query += " WHERE " + strings.Join(conditions, " AND ")
	}

	var user types.User

	row := db.QueryRow(query, args...)
	if err := row.Scan(&user.ID, &user.TelegramID, &user.UserName, &user.CreatedAt, &user.UpdatedAt); err != nil {
		if err == sql.ErrNoRows {
			return user, err
		}
		return user, fmt.Errorf("user: %v", err)
	}

	return user, nil
}

func InsertUser(params UserParams) (types.User, error) {
	db := GetConnection()

	query := `
		INSERT INTO users (telegram_id, username)
	  VALUES (?, ?);
	`

	result, err := db.Exec(query, *params.TelegramID, *params.UserName)

	if err != nil {
		return types.User{}, fmt.Errorf("InsertUser: %v", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return types.User{}, fmt.Errorf("InsertUser: %v", err)
	}

	return types.User{
		ID: int(id),
		TelegramID: int(*params.TelegramID),
		UserName: *params.UserName,
	}, nil
}
