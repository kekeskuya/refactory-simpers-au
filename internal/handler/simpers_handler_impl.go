package handler

import (
	"dummy-simpers-au/config"
	"dummy-simpers-au/internal/service"
)

type SimpersHandlerImpl struct {
	Env            *config.EnvironmentVariable
	SimpersService service.SimpersService
}

func NewSimpersHandler(
	env *config.EnvironmentVariable,
	simpersService service.SimpersService,
) SimpersHandler {
	return &SimpersHandlerImpl{
		Env:            env,
		SimpersService: simpersService,
	}
}
