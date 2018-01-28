package server

import (
	"github.com/urfave/cli"
	"gitlab.com/iH8c0ff33/cosmicbox-api-server/store"
	"gitlab.com/iH8c0ff33/cosmicbox-api-server/store/datastore"
)

func setupStore(c *cli.Context) store.Store {
	return datastore.New(
		c.String("driver"),
		c.String("datasource"),
	)
}
