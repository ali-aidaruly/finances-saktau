package config

import (
	"os"

	"github.com/ali-aidaruly/finances-saktau/internal/telegram"
	"github.com/joeshaw/envdecode"
)

type Config struct {
	TelegramBot telegram.Config
}

func ParseConfig(v *Config) error {
	defer os.Clearenv()

	return envdecode.StrictDecode(v)
}
