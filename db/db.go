package db

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// DB o'zgaruvchisi
var DB *sqlx.DB

// Ma'lumotlar bazasini ulash
func InitDB(dataSource string) error {
	var err error
	DB, err = sqlx.Connect("postgres", dataSource)
	if err != nil {
		return err
	}
	return DB.Ping()
}

// Qisqa URL ni olish
func GetURL(shortCode string) (string, error) {
	var longURL string
	err := DB.Get(&longURL, "SELECT long_url FROM urls WHERE short_code=$1", shortCode)
	return longURL, err
}

// Qisqa URL yaratish
func SaveURL(shortCode, longURL string) error {
	_, err := DB.Exec("INSERT INTO urls (short_code, long_url) VALUES ($1, $2)", shortCode, longURL)
	return err
}
