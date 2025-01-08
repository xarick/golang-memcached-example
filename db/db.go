package db

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var DB *sqlx.DB

func InitDB(dataSource string) error {
	var err error
	DB, err = sqlx.Connect("postgres", dataSource)
	if err != nil {
		return err
	}
	return DB.Ping()
}

func GetURL(shortCode string) (string, error) {
	var longURL string
	err := DB.Get(&longURL, "SELECT long_url FROM urls WHERE short_code=$1", shortCode)
	return longURL, err
}

func SaveURL(shortCode, longURL string) error {
	_, err := DB.Exec("INSERT INTO urls (short_code, long_url) VALUES ($1, $2)", shortCode, longURL)
	return err
}
