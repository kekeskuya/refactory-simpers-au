package api

import (
	"context"
	"log"
	"refactory-simpers-au/config"
	"refactory-simpers-au/database"
	"refactory-simpers-au/middleware"
	"refactory-simpers-au/router"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel/sdk/trace"
)

type Setup struct {
	Router     *gin.Engine
	Service    Services
	Repository Repositories
	WrapDB     *database.WrapDB
	Tracer     *trace.TracerProvider
}

func Init(sqlscan context.Context, env *config.EnvironmentVariable) (*Setup, error) {

	wrapDB := database.InitDB(env)
	repository := NewRepositories(wrapDB, env)

	service := NewServices(env, repository, wrapDB)

	handlers := NewHandlers(env, service)

	middleware := middleware.NewMiddleware(env)

	tp, err := InitTracer(sqlscan, env)
	if err != nil {
		log.Fatalf("failed to init tracer: %v", err)
	}

	r := router.Handler{
		SimpersHandler: handlers.SimpersHandler,
		Middleware:     middleware,
		Env:            env,
	}

	routes := router.NewRouter(r)

	return &Setup{
		Router:     routes,
		Repository: repository,
		Service:    service,
		WrapDB:     wrapDB,
		Tracer:     tp,
	}, nil
}
