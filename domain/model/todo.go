package model

type Todo struct {
	ID int64 `db:"id"`
	UserID int64 `db:"user_id"`
	Name string `db:"name"`
}