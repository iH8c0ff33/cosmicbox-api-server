package server

import (
	"encoding/base64"
	"net/http"
	"time"

	"github.com/gin-gonic/contrib/ginrus"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	"gitlab.com/iH8c0ff33/cosmicbox-api-server/router"
	"gitlab.com/iH8c0ff33/cosmicbox-api-server/router/middleware"
	"gitlab.com/iH8c0ff33/cosmicbox-api-server/router/middleware/auth"
)

var ServerFlags = []cli.Flag{
	cli.BoolFlag{
		EnvVar: "COSMIC_DEBUG",
		Name:   "debug",
		Usage:  "starts the server in debug mode",
	},
	cli.StringFlag{
		EnvVar: "COSMIC_SERVER_ADDR",
		Name:   "server_addr",
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
	cli.StringFlag{
		EnvVar: "OAUTH_REDIRECT_URL,REDIRECT_URL",
		Name:   "redirect_url",
		Usage:  "oauth2 redirect url",
		Value:  "http://localhost:9000/auth",
	},
	cli.StringFlag{
		EnvVar: "OAUTH_CREDENTIALS_FILE,CREDENTIALS_FILE",
		Name:   "credentials_file",
		Usage:  "oauth2 credentials file",
		Value:  "credentials.json",
	},
	cli.StringFlag{
		EnvVar: "COSMIC_SECRET,SECRET",
		Name:   "secret",
		Usage:  "base64 encoded secret",
		Value:  "8yU9RYGeNCHa6y3siMpxFxj5m4/WACtq",
	},
}

func Server(c *cli.Context) error {

	if c.Bool("debug") {
		logrus.SetLevel(logrus.DebugLevel)
	} else {
		logrus.SetLevel(logrus.WarnLevel)
	}

	scopes := []string{
		"https://www.googleapis.com/auth/userinfo.email",
	}

	mstore := setupStore(c)
	secret, err := base64.StdEncoding.DecodeString(c.String("secret"))
	if err != nil {
		logrus.Errorln(err)
		logrus.Fatalln("server: failed to parse base64 secret")
	}
	auth.Setup(
		c.String("redirect_url"),
		c.String("credentials_file"),
		scopes,
		secret,
	)

	handler := router.Load(
		ginrus.Ginrus(logrus.StandardLogger(), time.RFC3339, true),
		middleware.Version,
		middleware.Store(c, mstore),
		auth.Session("cb"),
	)

	return http.ListenAndServe(c.String("server_addr"), handler)
}
