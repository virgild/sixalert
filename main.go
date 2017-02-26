package main

import (
	"os"

	"github.com/urfave/cli"

	"github.com/virgild/sixalert/fetcher"
)

const TTC_RSS_URL = "https://ttc.ca/RSS/Service_Alerts/index.rss"

func main() {
	app := cli.NewApp()
	app.Name = "sixalert"
	app.Version = "0.1.0"
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
		fetcher.RssFetchCommand(TTC_RSS_URL),
	}

	app.Run(os.Args)
}
