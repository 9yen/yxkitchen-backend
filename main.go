package main

import (
	"os"
	"yxkitchen-backend/conf"
	"yxkitchen-backend/pkg/logger"

	"github.com/urfave/cli/v2"
)

func start() {
	logger.InitLogger(conf.C.Logger)
}

func main() {
	app := cli.NewApp()
	app.Name = "YuexianKitchen Backend"
	app.Usage = "YuexianKitchen Backend Server"
	app.Version = "1.0.0"
	app.Authors = []*cli.Author{
		{
			Name:  "9yen",
			Email: "3267666759@qq.com",
		},
	}
	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:    "config",
			Aliases: []string{"c"},
			Value:   "settings/conf-dev.toml",
			Usage:   "config file",
		},
	}
	app.Action = func(ctx *cli.Context) error {
		if err := conf.LoadConfig(ctx.String("config")); err != nil {
			logger.FatalString("Main", "load config", err.Error())
		}
		start()
		return nil
	}
	if err := app.Run(os.Args); err != nil {
		logger.FatalString("Main", "app run", err.Error())
	}

}
