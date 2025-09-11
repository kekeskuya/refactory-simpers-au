package sqlserver

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"time"

	"refactory-simpers-au/config"

	_ "github.com/microsoft/go-mssqldb"
	"github.com/rs/zerolog/log"
)

type WrapDB struct {
	Conn *sql.DB
}

func buildDSN(env *config.EnvironmentVariable) string {
	host := os.Getenv("SQLSERVER_HOST")
	port := os.Getenv("SQLSERVER_PORT")
	user := os.Getenv("SQLSERVER_USER")
	pass := os.Getenv("SQLSERVER_PASSWORD")
	name := os.Getenv("SQLSERVER_DB")

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
