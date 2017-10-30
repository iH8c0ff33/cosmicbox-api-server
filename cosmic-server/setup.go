package main

import (
	"github.com/urfave/cli"
	"git.deutron.ml/iH8c0ff33/cosmicbox-api-server/store"
	"git.deutron.ml/iH8c0ff33/cosmicbox-api-server/store/datastore"	
)

func setupStore(c *cli.Context) store.Store {
	return datastore.New(
		c.String("driver"),
		c.String("datasource"),
	)
}