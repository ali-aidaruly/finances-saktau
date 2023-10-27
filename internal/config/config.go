package config

import (
	"os"
	"time"

	"github.com/ali-aidaruly/finances-saktau/pkg/logger"

	"github.com/ali-aidaruly/finances-saktau/internal/telegram"
	"github.com/joeshaw/envdecode"
)

type Config struct {
	TelegramBot telegram.Config
	Database    DbConfig
	Logger      logger.Config
}

type DbConfig struct {
	ConnectionString      string        `env:"SQL_CONNECTION_STRING,required"`
	Driver                string        `env:"SQL_DRIVER,required"`
	MaxIdleConnections    int           `env:"SQL_MAX_IDLE_CONNECTIONS,required"`
	MaxOpenConnections    int           `env:"SQL_MAX_OPEN_CONNECTIONS,required"`
	ConnectionMaxLifetime time.Duration `env:"SQL_CONNECTION_MAX_LIFETIME,required"`
	Tracing               bool          `env:"SQL_TRACING"`
}

func ParseConfig(v *Config) error {
	defer os.Clearenv()

	return envdecode.StrictDecode(v)
}
