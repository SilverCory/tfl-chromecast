package tfl

// BusStop is a bus stop.
// The message will only be populated if there is an error?
type BusStop struct {
	Name     string
	Towards  string
	Message  string
	BusTimes []BusTime
}

func (bs BusStop) String() string {
	ret := bs.Name + "\n" + bs.Towards + "\n"
	if bs.Message != "" {
		ret += "\t" + bs.Message
	}

	for _, v := range bs.BusTimes {
		ret += "\t" + v.String() + "\n"
	}

	return ret
}

// BusTime is the time a specific bus going to a specific destination will arrive at a BusStop.
type BusTime struct {
	RouteNumber string
	Destination string
	Arrival     string
}

func (bt BusTime) String() string {
	return bt.RouteNumber + " - " + bt.Destination + " - " + bt.Arrival
}
