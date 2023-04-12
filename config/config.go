package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Port   string
	JWTKey string
	JWTTTL int
	PgURL  string
}

func New() (*Config, error) {
	viper.SetConfigFile("./.env")
	viper.ReadInConfig()

	url := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		viper.GetString("DB_HOST"),
		viper.GetString("DB_PORT"),
		viper.GetString("DB_USER"),
		viper.GetString("DB_PASSWORD"),
		viper.GetString("DB_NAME"),
		viper.GetString("DB_SS_MODE"),
	)

	cfg := Config{
		Port:   viper.GetString("PORT"),
		JWTKey: viper.GetString("JWT_SECRET_KEY"),
		JWTTTL: viper.GetInt("JWT_TTL"),
		PgURL:  url,
	}

	return &cfg, nil
}
