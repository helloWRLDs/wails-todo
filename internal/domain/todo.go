package domain

import "time"

type Todo struct {
	ID        int64     `db:"id"`
	Body      string    `db:"body"`
	IsDone    bool      `db:"is_done"`
	Priority  int32     `db:"priority"`
	CreatedAt time.Time `db:"created_at"`
}
