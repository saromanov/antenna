package main

import (
	"os"

	"github.com/saromanov/antenna/command/client"
	"github.com/urfave/cli/v2"
)

type InfoResponse struct {
}

func main() {
	app := &cli.App{
		Name:  "antenna",
		Usage: "make an explosive entrance",
		Commands: []*cli.Command{
			{
				Name:    "info",
				Aliases: []string{"i"},
				Usage:   "return info about running",
				Action: func(c *cli.Context) error {
					var resp InfoResponse
					err := client.Get("localhost:1255/v1/info", &resp)
					if err != nil {
						panic(err)
					}
					return nil
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		return
	}
}
