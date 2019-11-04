package main

import (
	"log"
	"os"
	"server/pkg/config"
	"server/pkg/dbcli"

	"github.com/urfave/cli"
)

func main() {
	if err := config.Parse(); err != nil {
		log.Println("config parsing: ", err)
		panic(err)
	}
	force := false

	app := cli.NewApp()
	app.Commands = []cli.Command{
		{
			Name:  "create",
			Usage: "create [database [table]] all if unspecified",
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:        "f",
					Usage:       "force action (drop; create)",
					Destination: &force,
				},
			},
			Action: func(c *cli.Context) error {
				if force {
					if err := dbcli.Drop(c.Args().First(), c.Args().Get(1)); err != nil {
						log.Println(err)
					}
				}
				return dbcli.Create(c.Args().First(), c.Args().Get(1))
			},
		},
		{
			Name:  "drop",
			Usage: "drop [database [table]] all if unspecified",
			Action: func(c *cli.Context) error {
				return dbcli.Drop(c.Args().First(), c.Args().Get(1))
			},
		},
		{
			Name:  "fill",
			Usage: "fill [database [table]] all if unspecified",
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:        "f",
					Usage:       "force action (drop; create; fill)",
					Destination: &force,
				},
			},
			Action: func(c *cli.Context) error {
				if force {
					if err := dbcli.Drop(c.Args().First(), c.Args().Get(1)); err != nil {
						log.Println(err)
					}
					if err := dbcli.Create(c.Args().First(), c.Args().Get(1)); err != nil {
						panic(err)
					}
				}
				return dbcli.Fill(c.Args().First(), c.Args().Get(1))
			},
		},
		{
			Name:  "all",
			Usage: "drop; create; fill",
			Action: func(c *cli.Context) error {
				if err := dbcli.Drop("", ""); err != nil {
					log.Println(err)
				}
				if err := dbcli.Create("", ""); err != nil {
					panic(err)
				}
				if err := dbcli.Fill("", ""); err != nil {
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
