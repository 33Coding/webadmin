package db

import (
	"github.com/go-pg/pg/v10"
	"os"
)

func ConnectORM() (*pg.DB, error) {

	opt, err := pg.ParseURL(os.Getenv("DB_URL"))

	if err != nil {
		return nil, err
	}

	db := pg.Connect(opt)

	err = CreateSchema(db)

	if err != nil {
		return nil, err
	}

	return db, nil
}
