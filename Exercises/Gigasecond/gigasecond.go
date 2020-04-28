package gigasecond

import "time"

// Gigasecond is 10^9 seconds
const gigasecond = 1e9

// AddGigasecond is a func which multiples the number of seconds to a gigasecond
func AddGigasecond( t time.Time) time.Time {
	return t.Add(time.Second * gigasecond)
}

