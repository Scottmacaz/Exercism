//Package gigasecond contains a method for determining the time you turned a gigasecond
package gigasecond

// import path for the time package from the standard library
import "time"

// AddGigasecond takes in a time and returns a new time which is the time + 1 gigasecond
func AddGigasecond(t time.Time) time.Time {
	return t.Add(time.Second * 1000000000)
}
