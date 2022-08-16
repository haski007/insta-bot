package main

import (
	"fmt"
	"os"

	"github.com/haski007/insta-bot/pkg/run"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

const (
	Version = "local"

	serviceName = "instabot"
)

func main() {
	app := cli.App{
		Name:    serviceName,
		Usage:   "Telegram bot to publish instagram posts",
		Version: Version,
		Flags:   flags(),
		Action:  action(),
	}

	if err := app.Run(os.Args); err != nil {
		logrus.Panicln(err)
	}
}

func flags() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:    "service_name",
			Value:   serviceName,
			EnvVars: []string{"SERVICE_NAME"},
		},
		&cli.StringFlag{
			Name:    "log_level",
			Value:   "INFO",
			EnvVars: []string{"LOG_LEVEL"},
		},
		&cli.StringFlag{
			Name:    "metrics_server_addr",
			Usage:   "Address of prometheus listener.",
			Value:   ":9091",
			EnvVars: []string{"METRICS_SERVER_ADDR"},
		},
		&cli.StringFlag{
			Name:    "config",
			Aliases: []string{"c"},
			EnvVars: []string{"CONFIG_PATH"},
		},
	}
}

func action() cli.ActionFunc {
	return func(cliCTX *cli.Context) error {
		args := run.Args{
			ServiceName: cliCTX.String("service_name"),
			ConfigFile:  cliCTX.String("config"),
			MetricsAddr: cliCTX.String("metrics_server_addr"),
			LogLevel:    run.LogLevel(cliCTX.String("log_level")),
		}
		if err := args.Validate(); err != nil {
			return fmt.Errorf("validate args err: %w", err)
		}

		if err := Run(cliCTX.Context, args); err != nil {
			return fmt.Errorf("run app err: %w", err)
		}

		return nil
	}
}
