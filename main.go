package main

import (
	"log"
	"os"
	"sort"

	"github.com/airheartdev/payload/internal"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:     "payload",
		Version:  internal.Version,
		Commands: internal.Commands(),
	}

	sort.Sort(cli.FlagsByName(app.Flags))

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
