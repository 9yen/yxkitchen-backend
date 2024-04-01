package conf

import (
	"os"

	"github.com/BurntSushi/toml"
)

type AppConf struct {
	Env string `toml:"env"`
}

type HttpConf struct {
	Addr string `toml:"addr"`
}

type LoggerConf struct {
	Level     string `toml:"log_level"`
	Type      string `toml:"log_type"`
	Filename  string `toml:"log_name"`
	MaxSize   int    `toml:"log_max_size"`
	MaxBackup int    `toml:"log_max_backup"`
	MaxAge    int    `toml:"log_max_age"`
	Compress  bool   `toml:"log_compress"`
}

type Config struct {
	App    AppConf    `toml:"app"`
	Http   HttpConf   `toml:"http"`
	Logger LoggerConf `toml:"logger"`
}

var C Config

func LoadConfig(filename string) error {

	content, err := os.ReadFile(filename)
	if err != nil {
		return err
	}
	if _, err := toml.Decode(string(content), &C); err != nil {
		return err
	}
	return nil
}
