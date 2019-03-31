package main

import (
	"os"

	"github.com/iH8c0ff33/cosmicbox-api-server/server"
	"github.com/iH8c0ff33/cosmicbox-api-server/version"
	"github.com/urfave/cli"
)

var flags = []cli.Flag{
	cli.StringFlag{
		EnvVar: "COSMIC_DB_DRIVER,DB_DRIVER",
		Name:   "driver",
		Usage:  "database driver",
		Value:  "sqlite3",
	},
	cli.StringFlag{
		EnvVar: "COSMIC_DB_DATASOURCE,DB_CONFIG",
		Name:   "datasource",
		Usage:  "database driver config string",
		Value:  "cosmic.sqlite",
	},
}

func main() {
	app := cli.NewApp()

	app.Name = "Cosmic Box API Server"
	app.Usage = "Cosmic Box Event Store API and utilities"
	app.Version = version.Version.String()
	app.Authors = []cli.Author{
		{
			Name:  "Daniele Monteleone",
			Email: "daniele.monteleone.it@gmail.com",
		},
	}
	app.Copyright = "(c) 2019 Daniele Monteleone"

	app.Flags = flags

	app.Commands = []cli.Command{
		{
			Name:     "server",
			Aliases:  []string{"serve"},
			Category: "API",
			Usage:    "start the API server",
			Flags:    server.ServerFlags,
			Action:   server.Server,
		},
	}

	app.Run(os.Args)
}
