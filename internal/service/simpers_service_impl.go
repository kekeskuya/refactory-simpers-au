package service

import (
	"dummy-simpers-au/config"
	"dummy-simpers-au/internal/repository"

	"github.com/jackc/pgx/v5/pgxpool"
)

type SimpersServiceImpl struct {
	env           *config.EnvironmentVariable
	db            *pgxpool.Pool
	personnelRepo repository.PersonnelRepository
}

func NewSimpersService(
	env *config.EnvironmentVariable,
	db *pgxpool.Pool,
	personnelRepo repository.PersonnelRepository,
) SimpersService {
	return &SimpersServiceImpl{
		env:           env,
		db:            db,
		personnelRepo: personnelRepo,
	}
}
