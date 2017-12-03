/*
Package gigasecond implements a simple "AddGigasecond" function for
checking the time one Gigasecond from a given input time.

It may be expanded in future to provide additional Gigasecond-related
functionality, or not.

It accepts a time.Time as input, and returns similar.

*/
package gigasecond

import "time"

// AddGigasecond adds 10^9 seconds to input time
func AddGigasecond(t time.Time) time.Time {
	return t.Add(1e9 * time.Second)
}
