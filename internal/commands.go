package internal

import (
	"github.com/airheartdev/payload/internal/cmd/create"
	"github.com/urfave/cli/v2"
)

var Version = "0.0.1"

func Commands() []*cli.Command {
	return []*cli.Command{
		{
			Name:   "create",
			Usage:  "Dumps the database and stores it on Cloud Storage",
			Action: create.Create,
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:     "project",
					Required: true,
					Aliases:  []string{"p"},
					Usage:    "The project to store the database in",
					EnvVars:  []string{"GOOGLE_PROJECT_ID"},
				},
				&cli.StringFlag{
					Name:     "bucket",
					Required: true,
					Aliases:  []string{"b"},
					Usage:    "The bucket to store the database in",
					EnvVars:  []string{"BUCKET"},
				},

				&cli.StringFlag{
					Name:     "database",
					Required: true,
					Aliases:  []string{"db"},
					Usage:    "The name of the database",
					EnvVars:  []string{"DATABASE_NAME"},
				},

				&cli.StringFlag{
					Name:     "instance",
					Required: true,
					Aliases:  []string{"i"},
					Usage:    "The Cloud SQL database instance",
					EnvVars:  []string{"CLOUDSQL_INSTANCE"},
				},
			},
		},
		{
			Name:  "pull",
			Usage: "Pulls the latest database dump, creating it if not already stored",
		},
		{
			Name:  "sync",
			Usage: "Loads the latest database dump into your local database, pulling it if not already loaded",
		},
	}
}
