package server

import (
	"github.com/urfave/cli"
	"github.com/iH8c0ff33/cosmicbox-api-server/store"
	"github.com/iH8c0ff33/cosmicbox-api-server/store/datastore"
)

func setupStore(c *cli.Context) store.Store {
	return datastore.New(
		c.String("driver"),
		c.String("datasource"),
	)
}
