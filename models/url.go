package models

import "time"

// URL modeli
type URL struct {
	ID        int       `db:"id"`
	ShortCode string    `db:"short_code"`
	LongURL   string    `db:"long_url"`
	CreatedAt time.Time `db:"created_at"`
}
