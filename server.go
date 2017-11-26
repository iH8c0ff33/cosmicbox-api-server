package main

import (
	"net/http"
	"time"

	"git.deutron.ml/iH8c0ff33/cosmicbox-api-server/router"
	"git.deutron.ml/iH8c0ff33/cosmicbox-api-server/router/middleware"
	"github.com/gin-gonic/contrib/ginrus"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

var serverFlags = []cli.Flag{
	cli.BoolFlag{
		EnvVar: "COSMIC_DEBUG",
		Name:   "debug",
		Usage:  "starts the server in debug mode",
	},
	cli.StringFlag{
		EnvVar: "COSMIC_SERVER_ADDR",
		Name:   "server-addr",
		Usage:  "server address",
		Value:  ":9000",
	},
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

func server(c *cli.Context) error {

	if c.Bool("debug") {
		logrus.SetLevel(logrus.DebugLevel)
	} else {
		logrus.SetLevel(logrus.WarnLevel)
	}

	mstore := setupStore(c)

	handler := router.Load(
		ginrus.Ginrus(logrus.StandardLogger(), time.RFC3339, true),
		middleware.Version,
		middleware.Store(c, mstore),
	)

	return http.ListenAndServe(c.String("server-addr"), handler)
}
