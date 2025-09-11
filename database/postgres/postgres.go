package postgres

import (
	"context"
	"refactory-simpers-au/config"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog/log"
)

type WrapDB struct {
	Conn *pgxpool.Pool
}

func InitDatabase(env *config.EnvironmentVariable) *WrapDB {
	connStr := fmt.Sprintf("postgres://%s:%s@%s/%s",
		env.DB.Postgres.Username,
		env.DB.Postgres.Password,
		env.DB.Postgres.Host,
		env.DB.Postgres.Name,
	)

	config, err := pgxpool.ParseConfig(connStr)
	if err != nil {
		log.Fatal().Err(err).Str("database", env.DB.Postgres.Name).Msg("[x] failed to parse connection config for postgres")
		panic(err)
	}

	conn, err := pgxpool.New(context.Background(), config.ConnString())
	if err != nil {
		log.Fatal().Err(err).Str("database", env.DB.Postgres.Name).Msg("[x] failed to connect to postgres")
		panic(err)
	}

	if err := conn.Ping(context.Background()); err != nil {
		log.Fatal().Err(err).Str("database", env.DB.Postgres.Name).Msg("[x] failed to ping postgres")
		panic(err)
	}

	return &WrapDB{
		Conn: conn,
	}
}
