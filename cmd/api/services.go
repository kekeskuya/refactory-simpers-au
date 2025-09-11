package api

import (
	"refactory-simpers-au/config"
	"refactory-simpers-au/database"
	"refactory-simpers-au/internal/service"
)

type Services struct {
	SimpersService service.SimpersService
}

func NewServices(
	env *config.EnvironmentVariable,
	r Repositories,
	db *database.WrapDB,
) Services {
	simpersService := service.NewSimpersService(env, db.SQLserver.Conn, r.PersonelRepository, r.NPWPRepository, r.AsabriRepository, r.PasporRepository, r.LampiranRepository)
	return Services{
		SimpersService: simpersService,
	}
}
