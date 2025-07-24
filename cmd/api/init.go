package api

import (
	"context"
	"dummy-simpers-au/config"
	"dummy-simpers-au/database"
	"dummy-simpers-au/middleware"
	"dummy-simpers-au/router"
	"log"

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

func Init(ctx context.Context, env *config.EnvironmentVariable) (*Setup, error) {

	wrapDB := database.InitDB(env)
	repository := NewRepositories(wrapDB, env)

	service := NewServices(env, repository, wrapDB)

	handlers := NewHandlers(env, service)

	middleware := middleware.NewMiddleware(env)

	tp, err := InitTracer(ctx, env)
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
