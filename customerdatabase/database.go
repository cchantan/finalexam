package customerdatabase

import (
	"database/sql"
	"os"

	_ "github.com/lib/pq"
)

func GetDBConn() (*sql.DB, error) {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		return nil, err
	}
	return db, nil
}

func Createtable() error {
	db, err := GetDBConn()
	if err != nil {
		return err
	}
	defer db.Close()

	createTb := `
		CREATE TABLE IF NOT EXISTS customers(
			id     SERIAL PRIMARY KEY,
			name   TEXT,
			email  TEXT,
			status TEXT
		);
		`
	_, err = db.Exec(createTb)
	if err != nil {
		return err
	}
	return nil
}
