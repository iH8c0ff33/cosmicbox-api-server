package server

import (
	"git.deutron.ml/iH8c0ff33/cosmicbox-api-server/store"
	"git.deutron.ml/iH8c0ff33/cosmicbox-api-server/store/datastore"
	"github.com/urfave/cli"
)

func setupStore(c *cli.Context) store.Store {
	return datastore.New(
		c.String("driver"),
		c.String("datasource"),
	)
}
