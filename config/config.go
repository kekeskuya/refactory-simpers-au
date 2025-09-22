package config

import (
	"fmt"
	"os"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

const EnvFolder = "env"
const EnvFilename = ".env"

func LoadEnv() (env *EnvironmentVariable, err error) {
	envFile := fmt.Sprintf("%s/%s", EnvFolder, EnvFilename)

	v := viper.New()

	if _, err := os.Stat(envFile); err == nil {
		v.SetConfigFile(envFile)
		if err := v.ReadInConfig(); err != nil {
			log.Printf("Error reading .env file: %v", err)
			panic(err)
		}
		log.Info().Msg(".env file loaded successfully")
	} else {
		v.AutomaticEnv()
		log.Info().Msg(".env file not found, skipping loading")
	}

	err = v.Unmarshal(&env)
	if err != nil {
		log.Error().Err(err).Msg("viper error unmarshal config")
	}
	log.Info().Msg("Env Loaded")
	return
}

type EnvironmentVariable struct {
	App struct {
		Host string `mapstructure:"HOST"`
		Mode string `mapstructure:"MODE"`
	} `mapstructure:"APP"`
	DB struct {
		Timeout   time.Duration `mapstructure:"TIMEOUT"`
		SQLSERVER struct {
			Host     string `mapstructure:"HOST"`
			Username string `mapstructure:"USERNAME"`
			Password string `mapstructure:"PASSWORD"`
			Name     string `mapstructure:"NAME"`
			Port     string `mapstructure:"PORT"`
		} `mapstructure:"SQLSERVER"`
	} `mapstructure:"DB"`
	DB2 struct {
		Timeout    time.Duration `mapstructure:"TIMEOUT"`
		SQLSERVER2 struct {
			Host     string `mapstructure:"HOST"`
			Username string `mapstructure:"USERNAME"`
			Password string `mapstructure:"PASSWORD"`
			Name     string `mapstructure:"NAME"`
			Port     string `mapstructure:"PORT"`
		} `mapstructure:"SQLSERVER"` // 👈 must match .env
	} `mapstructure:"DB2"`
	Swagger struct {
		BasePath    string `mapstructure:"BASE_PATH"`
		Host        string `mapstructure:"HOST"`
		Title       string `mapstructure:"TITLE"`
		Description string `mapstructure:"DESCRIPTION"`
		Version     string `mapstructure:"VERSION"`
	} `mapstructure:"SWAGGER"`
	Tracer struct {
		Address     string `mapstructure:"ADDRESS"`
		ServiceName string `mapstructure:"SERVICE_NAME"`
	}
}
