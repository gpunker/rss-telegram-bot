package types

import "time"

type User struct {
	ID          int       `json:"id"`
	TelegramID  int       `json:"telegram_id"`
	UserName    string    `json:"username"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func Ptr[T any](v T) *T { return &v }