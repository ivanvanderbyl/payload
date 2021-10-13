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
		Name:        "payload",
		Usage:       "Do it in production, but locally.",
		Description: "Payload is a simple tool for working with production data in your local environment.",
		Version:     internal.Version,
		Commands:    internal.Commands(),
	}

	sort.Sort(cli.FlagsByName(app.Flags))

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
