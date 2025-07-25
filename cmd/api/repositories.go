package api

import (
	"dummy-simpers-au/config"
	"dummy-simpers-au/database"
	"dummy-simpers-au/internal/repository"
)

type Repositories struct {
	PersonelRepository repository.PersonelRepository
	NPWPRepository     repository.NPWPRepository
}

func NewRepositories(wrapDB *database.WrapDB, env *config.EnvironmentVariable) Repositories {
	return Repositories{
		PersonelRepository: repository.NewPersonelRepository(wrapDB, env),
		NPWPRepository:     repository.NewNPWPRepository(wrapDB, env),
	}
}
