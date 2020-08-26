package service

import (
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
	"patreon-statistics/internal/domain"
	patreon_user "patreon-statistics/internal/repository/patreon-user"
)

type PatreonUserService interface {
	GetOne(userId string) (*domain.PatreonUser, error)
}

type patreonUserService struct {
	patreonUserRepository patreon_user.PatreonUserRepository
}

func (s *patreonUserService) GetOne(userId string) (*domain.PatreonUser, error) {
	user, err := s.patreonUserRepository.GetOne(userId)
	if err == nil {
		return user, nil
	}
	if err != sql.ErrNoRows {
		return nil, errors.Wrap(err, fmt.Sprintf("failed get patreon user from DB, userId: %v", userId))
	}

	return &domain.PatreonUser{
		UserId: userId,
	}, nil
}

func NewPatreonUserService(r patreon_user.PatreonUserRepository) *patreonUserService {
	return &patreonUserService{
		patreonUserRepository: r,
	}
}
