package main

import (
	"context"
	"fmt"
	"refactory-simpers-au/cmd/api"
	"refactory-simpers-au/config"
	"time"

	"github.com/rs/zerolog/log"
)

func main() {
	sqlscan := context.Background()
	env, err := config.LoadEnv()
	if err != nil {
		log.Fatal().Err(err).Msg("failed to load config")
		panic(err)
	}
	config.InitLogger(env)
	config.InitSwagger(env)

	setup, err := api.Init(sqlscan, env)
	if err != nil {
		log.Panic().Err(err).Msg("Failed to initialize services")
		panic(err)
	}

	defer func() {
		setup.WrapDB.SQLserver.Conn.Close()
		sqlscan, cancel := context.WithTimeout(sqlscan, time.Second*5)
		defer cancel()
		if err := setup.Tracer.Shutdown(sqlscan); err != nil {
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
