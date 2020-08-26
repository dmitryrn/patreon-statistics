package main

import (
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"patreon-statistics/internal"
)

func main() {
	app := &cli.App{
		Name:  "boom",
		Usage: "make an explosive entrance",
		Action: func(c *cli.Context) error {
			err := internal.NewApp()

			return err
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
