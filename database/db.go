package database

import (
	"refactory-simpers-au/config"
	"refactory-simpers-au/database/sqlserver"
)

type WrapDB struct {
	SQLserver  *sqlserver.WrapDB
	SQLserver2 *sqlserver.WrapDB
}

// func InitDB(env *config.EnvironmentVariable) *WrapDB {
// 	sqlserverDB := sqlserver.InitDatabase(env)
// 	return &WrapDB{
// 		SQLserver: sqlserverDB,
// 	}
// }

// func InitDB2(env *config.EnvironmentVariable) *WrapDB {
// 	sqlserverDB2 := sqlserver.InitDatabase2(env)
// 	return &WrapDB{
// 		SQLserver2: sqlserverDB2,
// 	}
// }

func InitDB(env *config.EnvironmentVariable) *WrapDB {
	sqlserverDB := sqlserver.InitDatabase(env)
	sqlserverDB2 := sqlserver.InitDatabase2(env)

	return &WrapDB{
		SQLserver:  sqlserverDB,
		SQLserver2: sqlserverDB2,
	}
}
