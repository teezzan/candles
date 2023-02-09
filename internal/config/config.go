package config

import "github.com/teezzan/ohlc/internal/util"

type Config struct {
	Database   DatabaseConfig
	Server     ServerConfig
	OHLCConfig OHLCConfig
}

type DatabaseConfig struct {
	Host     string
	Port     int
	Username string
	Password string
	Database string
}

type ServerConfig struct {
	Host string
	Port int
}

type OHLCConfig struct {
	DiscardInCompleteRow bool
}

func Init() *Config {
	return &Config{
		Database: DatabaseConfig{
			Host:     util.GetString("DB_HOST", "localhost"),
			Port:     util.GetInt("DB_PORT", 5432),
			Username: util.GetString("DB_USERNAME", "root"),
			Password: util.GetString("DB_PASSWORD", "root"),
			Database: util.GetString("DB_NAME", "ohlc"),
		},
		Server: ServerConfig{
			Host: util.GetString("SERVER_HOST", "localhost"),
			Port: util.GetInt("SERVER_PORT", 8090),
		},
		OHLCConfig: OHLCConfig{
			DiscardInCompleteRow: util.GetBool("OHLC_DISCARD_INCOMPLETE_ROW", false),
		},
	}

}
