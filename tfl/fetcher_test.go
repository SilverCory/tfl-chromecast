package tfl

import (
	"fmt"
	"os"
	"testing"
	"time"
)

func TestFetcher(t *testing.T) {
	reciver := make(chan *BusStop)
	errChan := make(chan error)

	os.Setenv("TFL_URLS", "https://tfl.gov.uk/bus/stop/490010535W/old-coulsdon-tudor-rose,https://tfl.gov.uk/bus/stop/490015305W1/the-glade")

	go func() {
		for {
			select {
			case bt := <-reciver:
				{
					if bt == nil {
						fmt.Print("\n\n=======================================\n\n")
					} else {
						fmt.Println(bt)
					}
				}
			case err := <-errChan:
				panic(err)
			}
		}
	}()

	f := New(reciver)

	go func() {
		time.Sleep(45 * time.Second)
		if err := f.Close(); err != nil {
			panic(err)
		}
	}()

	f.StartFetching(errChan)
}
