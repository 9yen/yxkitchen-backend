package conf

import (
	"os"

	"github.com/BurntSushi/toml"
)

type AppConf struct {
	Env       string `toml:"env"`
	Url       string `toml:"app_url"`
	ApiDomain string `toml:"api_domain"`
	AppName   string `toml:"app_name"`
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

type PagingConf struct {
	PerPage         int    `toml:"perpage"`
	UrlQueryPage    string `toml:"url_query_page"`
	UrlQuerySort    string `toml:"url_query_sort"`
	UrlQueryOrder   string `toml:"url_query_order"`
	UrlQueryPerPage string `toml:"url_query_per_page"`
}

type DatabaseConf struct {
	Connection string `toml:"db_connection"`
}

type MysqlConf struct {
	Host               string `toml:"db_host"`
	Port               int    `toml:"db_port"`
	Database           string `toml:"db_database"`
	Username           string `toml:"db_username"`
	Password           string `toml:"db_password"`
	Charset            string `toml:"db_charset"`
	MaxIdleConnections int    `toml:"db_max_idle_connections"`
	MaxOpenConnections int    `toml:"db_max_open_connections"`
	MaxLifeSeconds     int    `toml:"db_max_life_seconds"`
}

type SqliteConf struct {
	SqlFile string `toml:"db_sql_file"`
}

type RedisConf struct {
	Host     string `toml:"redis_host"`
	Port     int    `toml:"redis_port"`
	Username string `toml:"redis_username"`
	Password string `toml:"redis_password"`
	Database int    `toml:"redis_database"`
}
type Config struct {
	App      AppConf      `toml:"app"`
	Http     HttpConf     `toml:"http"`
	Logger   LoggerConf   `toml:"logger"`
	Paging   PagingConf   `toml:"paging"`
	Database DatabaseConf `toml:"database"`
	Mysql    MysqlConf    `toml:"mysql"`
	Sqlite   SqliteConf   `toml:"sqlite"`
	Redis    RedisConf    `toml:"redis"`
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
