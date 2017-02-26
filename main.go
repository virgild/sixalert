package main

import (
	"os"

	"github.com/urfave/cli"

	"github.com/virgild/sixalert/fetcher"
)

func main() {
	app := cli.NewApp()
	app.Name = "sixalert"
	app.Version = "0.1.1"
	app.Authors = []cli.Author{
		cli.Author{
			Name:  "Virgil Dimaguila",
			Email: "virgild@gmail.com",
		},
	}
	app.Usage = "Runs the sixalert program"

	app.Action = func(c *cli.Context) error {
		cli.ShowAppHelp(c)
		return nil
	}

	app.Commands = []cli.Command{
		fetcher.RssFetchCommand(),
	}

	app.Run(os.Args)
}
