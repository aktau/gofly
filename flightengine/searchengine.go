package flightengine

type SearchEngine interface {
	/* get the average price per month */
	AvgPrice(orig string, dest string, month int, year int) float64
}
