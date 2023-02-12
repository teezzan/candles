package config

import "github.com/teezzan/ohlc/internal/util"

type Config struct {
	Database                  DatabaseConfig
	Server                    ServerConfig
	OHLCConfig                OHLCConfig
	S3Config                  S3Config
	SQSConfig                 SQSConfig
	CronJobFrequencyInMinutes int
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
	DiscardInCompleteRow  bool
	DefaultDataPointLimit int
}

type S3Config struct {
	Region               string
	Bucket               string
	PresignURLExpiryTime int
}

type SQSConfig struct {
	Region string
	Queue  string
}

func Init() *Config {
	return &Config{
		Database: DatabaseConfig{
			Host:     util.GetString("DB_HOST", defaultDBHost),
			Port:     util.GetInt("DB_PORT", defaultDBPort),
			Username: util.GetString("DB_USERNAME", defaultDBUsername),
			Password: util.GetString("DB_PASSWORD", defaultDBPassword),
			Database: util.GetString("DB_NAME", defaultDBName),
		},
		Server: ServerConfig{
			Host: util.GetString("SERVER_HOST", defaultServerHost),
			Port: util.GetInt("SERVER_PORT", defaultServerPort),
		},
		OHLCConfig: OHLCConfig{
			DiscardInCompleteRow:  util.GetBool("OHLC_DISCARD_INCOMPLETE_ROW", defaultDiscardInCompleteRow),
			DefaultDataPointLimit: util.GetInt("OHLC_DATA_POINT_LIMIT", defaultDataPointLimit),
		},
		S3Config: S3Config{
			Region:               util.GetString("S3_REGION", defaultS3Region),
			Bucket:               util.GetString("S3_BUCKET", defaultS3Bucket),
			PresignURLExpiryTime: util.GetInt("S3_PRESIGN_URL_EXPIRY_TIME", defaultS3PresignURLExpiryTime),
		},
		SQSConfig: SQSConfig{
			Region: util.GetString("SQS_REGION", defaultSQSRegion),
			Queue:  util.GetString("SQS_QUEUE", defaultSQSQueue),
		},
		CronJobFrequencyInMinutes: util.GetInt("CRON_JOB_FREQUENCY_IN_MINUTES", defaultCronJobFrequencyInMinutes),
	}

}
