package tfl

import (
	"net/http"
	"strings"
	"tfl-chromecast/env"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/pkg/errors"
)

// Fetcher is the interface to TFL.
type Fetcher struct {
	receiver chan *BusStop
	close    chan struct{}
	client   *http.Client
}

const (
	StopNameSelector    = "#station-stops-and-piers > div.ssp > h1"
	StopTowardsSelector = "#station-stops-and-piers > div.ssp > div > p"

	BaseSelector        = "#full-width-content > div:nth-child(4) > div.main > div.station-details > div:nth-child(5) > div"
	BasePlusOneSelector = "div"
	OrderedListSelector = "div > div.live-board.initial-board-container.bus > div > ol"
)

func New(receiver chan *BusStop) *Fetcher {
	return &Fetcher{
		receiver: receiver,
		close:    make(chan struct{}),
		client: &http.Client{
			Timeout: time.Second * 3,
		},
	}
}

func (f *Fetcher) Close() error {
	select {
	case f.close <- struct{}{}:
		return nil
	case <-time.After(3 * time.Second):
		return errors.New("timeout on close")
	}
}

func (f *Fetcher) StartFetching(errChan chan error) {
	for {
		select {

		case <-f.close:
			return

		case <-time.After(env.Get().RefreshTime):
			f.receiver <- nil // Clear the display.
			for _, url := range env.Get().TfLURLs {
				if err := f.getBusTime(url); err != nil {
					errChan <- err
				}
			}

		}
	}
}

func (f *Fetcher) getBusTime(URL string) (err error) {
	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		return errors.Wrap(err, "create request")
	}

	resp, err := f.client.Do(req)
	if err != nil {
		return errors.Wrap(err, "executing request")
	}

	defer func() {
		err1 := resp.Body.Close()
		if err == nil {
			err = errors.Wrap(err1, "closing body")
		}
	}()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return errors.Wrap(err, "creating NewDocumentFromReader")
	}

	var stop BusStop

	// Find an set the bus stop name
	stop.Name = strings.TrimSpace(doc.Find(StopNameSelector).Text())
	if stop.Name == "" {
		return errors.Errorf("unable to find bus stop name for url %q", URL)
	}

	// Find an set the bus stop towards text
	stop.Towards = strings.TrimSpace(doc.Find(StopTowardsSelector).Text())

	// Get the base, of errors or the list.
	base := doc.Find(BaseSelector)

	// Check for an error on the page (no busses in the next x mins..)
	basePlusOne := base.Find(BasePlusOneSelector)
	if basePlusOne.Size() == 0 {
		stop.Message = base.Children().After("strong").Text()
	}

	// Get and process the actual list.
	list := base.Find(OrderedListSelector).ChildrenFiltered("li")
	for i := 0; i < list.Size(); i++ {
		elem := goquery.NewDocumentFromNode(list.Get(i))
		busTime := BusTime{}
		busTime.RouteNumber = strings.TrimSpace(elem.Find("span.live-board-route.line-text").Text())
		busTime.Destination = strings.TrimSpace(elem.Find("span.live-board-destination > span.train-destination").Text())
		busTime.Arrival = strings.TrimSpace(elem.Find("span.live-board-eta").Text())

		stop.BusTimes = append(stop.BusTimes, busTime)
	}

	// Blocking, post to display, timeout after refresh time has passed.
	select {
	case f.receiver <- &stop:
		return nil
	case <-time.After(env.Get().RefreshTime):
		return errors.New("timeout on send bus time")
	}
}
