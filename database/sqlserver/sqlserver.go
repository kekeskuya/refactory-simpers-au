package sqlserver

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"refactory-simpers-au/config"

	_ "github.com/microsoft/go-mssqldb"
	"github.com/rs/zerolog/log"
)

type WrapDB struct {
	Conn *sql.DB
}

func buildDSN(env *config.EnvironmentVariable) string {
	host := env.DB.SQLSERVER.Host
	port := env.DB.SQLSERVER.Port
	user := env.DB.SQLSERVER.Username
	pass := env.DB.SQLSERVER.Password
	name := env.DB.SQLSERVER.Name

	// Debugging output
	fmt.Println("DEBUG USER:", user)
	fmt.Println("DEBUG PASS:", pass)
	fmt.Println("DEBUG HOST:", host)
	fmt.Println("DEBUG PORT:", port)
	fmt.Println("DEBUG NAME:", name)

	fmt.Println("DEBUG HOST:", host)

	if port == "" {
		port = "1433"
	}

	return fmt.Sprintf(
		"sqlserver://%s:%s@%s:%s?database=%s&encrypt=disable&TrustServerCertificate=true",
		user, pass, host, port, name,
	)
}

func InitDatabase(env *config.EnvironmentVariable) *WrapDB {
	dsn := buildDSN(env)

	db, err := sql.Open("sqlserver", dsn)
	if err != nil {
		log.Fatal().Err(err).Msg("[x] failed to open SQL Server connection")
	}

	db.SetMaxOpenConns(50)
	db.SetMaxIdleConns(10)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(30 * time.Minute)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := db.PingContext(ctx); err != nil {
		log.Fatal().Err(err).Msg("[x] failed to ping SQL Server")
	}

	log.Info().Msg("[√] SQL Server connected")
	return &WrapDB{Conn: db}
}

func buildDSN2(env *config.EnvironmentVariable) string {
	host2 := env.DB2.SQLSERVER2.Host
	port2 := env.DB2.SQLSERVER2.Port
	user2 := env.DB2.SQLSERVER2.Username
	pass2 := env.DB2.SQLSERVER2.Password
	name2 := env.DB2.SQLSERVER2.Name
	// Debugging output
	fmt.Println("DEBUG 2 USER:", user2)
	fmt.Println("DEBUG 2 PASS:", pass2)
	fmt.Println("DEBUG 2 HOST:", host2)
	fmt.Println("DEBUG 2 PORT:", port2)
	fmt.Println("DEBUG 2 NAME:", name2)
	fmt.Println("DEBUG 2 HOST:", host2)

	if port2 == "" {
		port2 = "1433"
	}
	return fmt.Sprintf(
		"sqlserver://%s:%s@%s:%s?database=%s&encrypt=disable&TrustServerCertificate=true",
		user2, pass2, host2, port2, name2,
	)
}

func InitDatabase2(env *config.EnvironmentVariable) *WrapDB {
	dsn := buildDSN2(env)

	db, err := sql.Open("sqlserver", dsn)
	if err != nil {
		log.Fatal().Err(err).Msg("[x] failed to open SQL Server connection")
	}

	db.SetMaxOpenConns(50)
	db.SetMaxIdleConns(10)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(30 * time.Minute)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := db.PingContext(ctx); err != nil {
		log.Fatal().Err(err).Msg("[x] failed to ping SQL Server 2")
	}

	log.Info().Msg("[√] SQL Server 2 connected")
	return &WrapDB{Conn: db}
}
