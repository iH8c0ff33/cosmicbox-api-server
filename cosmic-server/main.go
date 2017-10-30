package main

import (
	"fmt"
	"os"
	"git.deutron.ml/iH8c0ff33/cosmicbox-api-server/version"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()

	app.Name = "cosmicbox-server"
	app.Version = version.Version.String()
	app.Usage = "cosmicbox server"
	app.Action = server
	app.Flags = flags

	if err := app.Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
