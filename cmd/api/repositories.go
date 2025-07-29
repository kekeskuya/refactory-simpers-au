package api

import (
	"dummy-simpers-au/config"
	"dummy-simpers-au/database"
	"dummy-simpers-au/internal/repository"
)

type Repositories struct {
	PersonelRepository repository.PersonelRepository
	NPWPRepository     repository.NPWPRepository
	AsabriRepository   repository.AsabriRepository
	PasporRepository   repository.PasporRepository
	LampiranRepository repository.LampiranRepository
}

func NewRepositories(wrapDB *database.WrapDB, env *config.EnvironmentVariable) Repositories {
	return Repositories{
		PersonelRepository: repository.NewPersonelRepository(wrapDB, env),
		NPWPRepository:     repository.NewNPWPRepository(wrapDB, env),
		AsabriRepository:   repository.NewAsabriRepository(wrapDB, env),
		PasporRepository:   repository.NewPasporRepository(wrapDB, env),
		LampiranRepository: repository.NewLampiranRepository(wrapDB, env),
	}
}
