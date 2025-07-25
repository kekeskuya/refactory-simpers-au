package main

import (
	"context"
	"dummy-simpers-au/cmd/api"
	"dummy-simpers-au/config"
	"fmt"
	"time"

	"github.com/rs/zerolog/log"
)

func main() {
	ctx := context.Background()
	env, err := config.LoadEnv()
	if err != nil {
		log.Fatal().Err(err).Msg("failed to load config")
		panic(err)
	}
	config.InitLogger(env)
	config.InitSwagger(env)

	setup, err := api.Init(ctx, env)
	if err != nil {
		log.Panic().Err(err).Msg("Failed to initialize services")
		panic(err)
	}

	defer func() {
		setup.WrapDB.Postgres.Conn.Close()
		ctx, cancel := context.WithTimeout(ctx, time.Second*5)
		defer cancel()
		if err := setup.Tracer.Shutdown(ctx); err != nil {
			log.Printf("error shutting down tracer: %v", err)
		}
	}()

	go func() {
		err = setup.Router.Run(env.App.Host)
		if err != nil {
			log.Info().Msg(fmt.Sprintf("Listening on %s", env.App.Host))
		}
	}()

	select {}

}
