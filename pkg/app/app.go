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
