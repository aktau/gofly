package flightengine

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

/* implements the SearchEngine interface */
type Momondo struct {
	version int
}

type AvgPriceResponse struct {
	OrigCode   string  `xml:"OrigCode"`
	DestCode   string  `xml:"DestCode"`
	Month      int     `xml:"Month"`
	Year       int     `xml:"Year"`
	PriceEUR   float64 `xml:"PriceEUR"`
	SupplierId string  `xml:"SupplierId"`
}

const (
	MomUrl      = "http://www.momondo.com/Momondo.asmx"
	MomAvgPrice = "WhereToGoGetUpdated"
)

func (a *Momondo) AvgPrice(orig string, dest string, month int, year int) (float64, error) {
	url :=
		fmt.Sprintf("%v/%v?origCode=%v&destCode=%v&year=%v&month=%v",
			MomUrl, MomAvgPrice, orig, dest, year, month)
	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}

	fmt.Println(string(body))

	decoded := &AvgPriceResponse{}
	err = xml.Unmarshal(body, decoded)
	if err != nil {
		return 0, err
	}

	fmt.Println(decoded)

	return decoded.PriceEUR, nil
}
