package api

import (
	"dummy-simpers-au/config"
	"dummy-simpers-au/database"
	"dummy-simpers-au/internal/repository"
)

type Repositories struct {
	PersonnelRepository repository.PersonnelRepository
}

func NewRepositories(wrapDB *database.WrapDB, env *config.EnvironmentVariable) Repositories {
	return Repositories{
		PersonnelRepository: repository.NewPersonnelRepository(wrapDB, env),
	}
}
