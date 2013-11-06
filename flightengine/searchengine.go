package flightengine

import (
	"time"
)

/* all methods may block, run them in a goroutine if you
 * don't want that */
type SearchEngine interface {
	/* get the average price per month */
	PriceAvg(orig string, dest string, month int, year int) (float64, error)

	/* get the best price for a route */
	Price(orig string, dest string, dep time.Time, ret time.Time, live string) (float64, error)
	PriceOneWay(orig string, dest string, dep time.Time, live string) (float64, error)
}
