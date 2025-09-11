package api

import (
	"refactory-simpers-au/config"
	"refactory-simpers-au/internal/handler"
)

type Handlers struct {
	SimpersHandler handler.SimpersHandler
}

func NewHandlers(env *config.EnvironmentVariable, s Services) Handlers {
	return Handlers{
		SimpersHandler: handler.NewSimpersHandler(env, s.SimpersService),
	}
}
