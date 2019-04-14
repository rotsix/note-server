package main

import (
	"os"
	"server/pkg/config"
	"server/pkg/dbcli"

	"github.com/urfave/cli"
)

func main() {
	conf := config.Parse()
	force := false

	app := cli.NewApp()
	app.Commands = []cli.Command{
		{
			Name:  "create",
			Usage: "create given database, or all if unspecified",
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:        "f",
					Usage:       "force action (drop then create)",
					Destination: &force,
				},
			},
			Action: func(c *cli.Context) error {
				if force {
					if err := dbcli.Drop(c.Args().First(), conf); err != nil {
						panic(err)
					}
				}
				return dbcli.Create(c.Args().First(), conf)
			},
		},
		{
			Name:  "drop",
			Usage: "drop given database, or all if unspecified",
			Action: func(c *cli.Context) error {
				return dbcli.Drop(c.Args().First(), conf)
			},
		},
		{
			Name:  "fill",
			Usage: "fill given database with mock data, all dbs if unspecified",
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:        "f",
					Usage:       "force action (drop, create then fill)",
					Destination: &force,
				},
			},
			Action: func(c *cli.Context) error {
				if force {
					if err := dbcli.Drop(c.Args().First(), conf); err != nil {
						panic(err)
					}
					if err := dbcli.Create(c.Args().First(), conf); err != nil {
						panic(err)
					}
				}
				return dbcli.Fill(c.Args().First(), conf)
			},
		},
		{
			Name:  "all",
			Usage: "drop all databases, create new ones, and fill them with mock data (alias for 'fill -f')",
			Action: func(c *cli.Context) error {
				if err := dbcli.Drop("", conf); err != nil {
					panic(err)
				}
				if err := dbcli.Create("", conf); err != nil {
					panic(err)
				}
				if err := dbcli.Fill("", conf); err != nil {
					panic(err)
				}
				return nil
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		panic(err)
	}
}
