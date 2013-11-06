package flightengine

/**
 * initial exploration performed with cURL, response format can be XML or
 * JSON, but for JSON one needs to send JSON as well, all of these queries
 * provide the same information:
 *
 * When communicating in XML, momondo uses some standard datatypes, like
 * dateType, defined here:
 * http://www.w3.org/TR/2004/REC-xmlschema-2-20041028/datatypes.html#dateTime
 *
 * WhereToGoGetUpdated:
 * curl -H "Accept: application/json" --verbose http://www.momondo.com/Momondo.asmx/WhereToGoGetUpdated\?origCode\=BRU\&destCode\=LIM\&year\=2014\&month\=03
 * curl --data "origCode=BRU&destCode=LIM&year=2014&month=03" http://www.momondo.com/Momondo.asmx/WhereToGoGetUpdated
 * curl -H "Content-Type:application/json; charset=UTF-8" -H "Accept: application/json" --data '{"origCode":"BRU", "destCode":"LIM", "year": 2014, "month":3}' http://www.momondo.com/Momondo.asmx/WhereToGoGetUpdated
 *
 * a strange service that seems exposed on the internet:
 * http://213.42.52.80/dcawebservices/FIS/flights.asmx
 */

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

/* implements the SearchEngine interface */
type Momondo struct {
	version int
}

/* { */
/*     "adultCount": "1", */
/*     "childAges": [], */
/*     "country": null, */
/*     "culture": null, */
/*     "segList": [ */
/*         { */
/*             "Depart": "20140307", */
/*             "Destination": "LIM", */
/*             "Origin": "BRU", */
/*             "__type": "SkyGate.Momondo.Engine.FlightService.SearchSegment" */
/*         }, */
/*         { */
/*             "Depart": "20140321", */
/*             "Destination": "BRU", */
/*             "Origin": "LIM", */
/*             "__type": "SkyGate.Momondo.Engine.FlightService.SearchSegment" */
/*         } */
/*     ], */
/*     "ticketType": "ECO" */
/* } */
type priceRequestSegment struct {
	Depart      string `json:"Depart"`
	Destination string `json:"Destination"`
	Origin      string `json:"Origin"`
	SegType     string `json:"__type"`
}

type priceRequest struct {
	AdultCount int                   `json:"adultCount,string"`
	ChildAges  []int                 `json:"childAges"`
	Country    *string               `json:"country"`
	Culture    *string               `json:"culture"`
	SegList    []priceRequestSegment `json:"segList"`
	TicketType string                `json:"ticketType"`
}

type PriceResponse struct {
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
	MomBase           = "http://www.momondo.com"
	MomLegacyServ     = "Momondo.asmx"
	MomFlightServ     = "FlightWS.asmx"
	MomAvgPriceOp     = "WhereToGoGetUpdated"
	MomFlightSearchOp = "StartFlightSearch"
)

func (a *Momondo) PriceAvg(orig string, dest string, month int, year int) (float64, error) {
	url := fmt.Sprintf("%v/%v/%v?origCode=%v&destCode=%v&year=%v&month=%v",
		MomBase, MomLegacyServ, MomAvgPriceOp, orig, dest, year, month)

	decoded := &AvgPriceResponse{}
	err := momFetch(url, decoded)

	return decoded.PriceEUR, err
}

/**
 * curl --verbose \
 *   -H "Content-Type: application/json; charset=UTF-8" \
 *   --data @StartFlightSearch.json \
 *   http://www.momondo.com/FlightWS.asmx/StartFlightSearch
 */
func (a *Momondo) Price(orig string, dest string, dep time.Time, ret time.Time, live string) (float64, error) {
	url := fmt.Sprintf("%v/%v/%v", MomBase, MomFlightServ, MomFlightSearchOp)

	leg1 := priceRequestSegment{"20140307", "LIM", "BRU", "SkyGate.Momondo.Engine.FlightService.SearchSegment"}
	leg2 := priceRequestSegment{"20140321", "BRU", "LIM", "SkyGate.Momondo.Engine.FlightService.SearchSegment"}
	send := &priceRequest{0, nil, nil, nil, []priceRequestSegment{leg1, leg2}, "ECO"}
	/* resp := &PriceResponse{} */
	resp := make(map[string]interface{})
	err := momPost(url, send, &resp)

	fmt.Println("decoded map -> ", resp)

	return 400, err
}

func (a *Momondo) PriceOneWay(orig string, dest string, dep time.Time, live string) (float64, error) {
	return 400, nil
}

/* helper function that fetches data and unmarshals the json */
func momPost(url string, payload interface{}, recv interface{}) error {
	/* http request */
	payloadBuffer, err := json.MarshalIndent(payload, "", "\t")
	if err != nil {
		return err
	}
	fmt.Println("momondo: POST -> ", string(payloadBuffer))

	resp, err := http.Post(url, "application/json", bytes.NewReader(payloadBuffer))
	if err != nil {
		return err
	}

	/* read the entire response */
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	fmt.Println("momondo: POST <- ", string(body))

	/* unmarshal the XML response */
	err = json.Unmarshal(body, recv)
	if err != nil {
		return err
	}

	return nil
}

/* helper function that fetches data end unmarshals the XML */
func momFetch(url string, data interface{}) error {
	/* http request */
	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	/* read the entire response */
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	fmt.Println(string(body))

	/* unmarshal the XML response */
	err = xml.Unmarshal(body, data)
	if err != nil {
		return err
	}

	fmt.Println(data)

	return nil
}
