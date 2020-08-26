package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
)

type Db struct {
	Db *sql.DB
}

func (d *Db) runMigrations() error {
	_, err := d.Db.Exec(`create table if not exists patreon_user (
		user_id varchar(120) primary key
    )`)

	return err
}

func NewDb() (*Db, error) {
	sqlDb, err := sql.Open("postgres", "postgresql://localhost/test?user=test&password=test&sslmode=disable")
	if err != nil {
		return nil, err
	}

	db := &Db{
		Db: sqlDb,
	}

	err = db.runMigrations()
	if err != nil {
		return nil, errors.Wrap(err, "fail run migrations")
	}

	return db, nil
}
