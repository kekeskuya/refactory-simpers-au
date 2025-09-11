package database

import (
	"refactory-simpers-au/config"
	"refactory-simpers-au/database/sqlserver"
)

type WrapDB struct {
	SQLserver *sqlserver.WrapDB
}

func InitDB(env *config.EnvironmentVariable) *WrapDB {
	sqlserverDB := sqlserver.InitDatabase(env)
	return &WrapDB{
		SQLserver: sqlserverDB,
	}
}
