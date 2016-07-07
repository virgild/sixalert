package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/virgild/sixalert/alert"

	rss "github.com/jteeuwen/go-pkg-rss"
	"github.com/urfave/cli"
)

func rssFetchCommand(sourceURL string) cli.Command {
	cmd := cli.Command{
		Name: "rss",
		Action: func(c *cli.Context) error {
			rssFetchAndPrint(sourceURL)
			return nil
		},
	}
	return cmd
}

func rssFetchAndPrint(sourceURL string) {
	cacheTimeout := 10
	enforceCacheLimit := true
	f := &fetcher{}
	feed := rss.NewWithHandlers(cacheTimeout, enforceCacheLimit, f, f)
	if err := feed.Fetch(sourceURL, nil); err != nil {
		panic(err)
	}
}

type fetcher struct {
}

func (f *fetcher) ProcessChannels(feed *rss.Feed, newChannels []*rss.Channel) {

}

func (f *fetcher) ProcessItems(feed *rss.Feed, ch *rss.Channel, newItems []*rss.Item) {
	for index, item := range newItems {
		alert, err := parseItem(item)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%d: %s - (%s) %s\n", index+1, alert.Timestamp.Format("3:04pm"), alert.Affecting, alert.Content)
	}
}

func parseItem(item *rss.Item) (*alert.TTCAlert, error) {
	rssExtensions := item.Extensions["http://purl.org/rss/1.0/"]
	contentValue := rssExtensions["description"][0].Value
	contentElements := strings.Split(contentValue, "\n- Affecting: ")
	contentText := contentElements[0]

	var affectingText string
	if len(contentElements) > 1 {
		affectingText = contentElements[1]
	}

	var pubDate time.Time
	var err error
	dcExtensions := item.Extensions["http://purl.org/dc/elements/1.1/"]
	if len(dcExtensions["date"]) > 0 {
		pubDate, err = time.Parse("2006-01-02T15:04:05.000-07:00", dcExtensions["date"][0].Value)
	}

	return alert.NewTTCAlert(contentText, affectingText, pubDate), err
}
