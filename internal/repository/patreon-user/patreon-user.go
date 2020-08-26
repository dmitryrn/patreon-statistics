package patreon_user

import (
	"database/sql"
	"github.com/pkg/errors"
	"log"
	db "patreon-statistics/internal/db"
	"patreon-statistics/internal/domain"
)

type PatreonUserRepository interface {
	GetOne(userId string) (*domain.PatreonUser, error)
}

type patreonUserService struct {
	db *db.Db
}

func (s *patreonUserService) GetOne(userId string) (*domain.PatreonUser, error) {
	row := s.db.Db.QueryRow("select user_id from patreon_user where user_id = $1 limit 1", userId)

	user := &domain.PatreonUser{}
	err := row.Scan(&user.UserId)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("not found user in db", userId)
			return nil, err
		}
		return nil, errors.Wrap(err, "fail scan row")
	}

	return user, nil
}

func NewPatreonUserRepository(db *db.Db) PatreonUserRepository {
	return &patreonUserService{
		db: db,
	}
}
