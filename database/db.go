package database

import (
	"dummy-simpers-au/config"
	"dummy-simpers-au/database/postgres"
)

type WrapDB struct {
	Postgres *postgres.WrapDB
}

func InitDB(env *config.EnvironmentVariable) *WrapDB {
	postgresDB := postgres.InitDatabase(env)

	return &WrapDB{
		Postgres: postgresDB,
	}
}
