package api

import (
	"dummy-simpers-au/config"
	"dummy-simpers-au/database"
	"dummy-simpers-au/internal/service"
)

type Services struct {
	SimpersService service.SimpersService
}

func NewServices(
	env *config.EnvironmentVariable,
	r Repositories,
	db *database.WrapDB,
) Services {
	simpersService := service.NewSimpersService(env, db.Postgres.Conn, r.PersonelRepository, r.NPWPRepository, r.AsabriRepository, r.PasporRepository, r.LampiranRepository)
	return Services{
		SimpersService: simpersService,
	}
}
