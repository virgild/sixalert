package fetcher

import (
	"fmt"
	"strings"
	"time"

	"github.com/virgild/sixalert/alert"

	rss "github.com/jteeuwen/go-pkg-rss"
)

const TTC_RSS_URL = "https://ttc.ca/RSS/Service_Alerts/index.rss"

func FetchAndPrint() {
	fetcher := NewFetcher()
	err := fetcher.FetchCurrent()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		for n, alert := range fetcher.CurrentAlerts() {
			fmt.Printf("%d. %s\n", n+1, alert)
		}
	}
}

type Fetcher struct {
	sourceURL string
	alerts    []*alert.TTCAlert
}

func NewFetcher() *Fetcher {
	fetcher := &Fetcher{
		sourceURL: TTC_RSS_URL,
		alerts:    []*alert.TTCAlert{},
	}

	return fetcher
}

func (f *Fetcher) FetchCurrent() error {
	cacheTimeout := 10
	enforceCacheLimit := true
	feed := rss.NewWithHandlers(cacheTimeout, enforceCacheLimit, nil, f)
	err := feed.Fetch(f.sourceURL, nil)
	return err
}

func (f *Fetcher) CurrentAlerts() []*alert.TTCAlert {
	return f.alerts
}

func (f *Fetcher) ProcessChannels(feed *rss.Feed, newChannels []*rss.Channel) {

}

func (f *Fetcher) ProcessItems(feed *rss.Feed, ch *rss.Channel, newItems []*rss.Item) {
	f.alerts = []*alert.TTCAlert{}
	for _, item := range newItems {
		alert, err := parseItem(item)
		if err != nil {
			panic(err)
		} else {
			f.alerts = append(f.alerts, alert)
		}
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
