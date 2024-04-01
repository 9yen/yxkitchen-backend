package app

import (
	"yxkitchen-backend/conf"
)

func IsDev() bool {
	return conf.C.App.Env == "dev"
}

func IsProduction() bool {
	return conf.C.App.Env == "production"
}

func IsTesting() bool {
	return conf.C.App.Env == "testing"
}

// URL 传参 path 拼接站点的 URL
func URL(path string) string {
	return conf.C.App.Url + path
}

// V1URL 拼接带 v1 标示 URL
func V1URL(path string) string {
	return URL("/v1/" + path)
}
