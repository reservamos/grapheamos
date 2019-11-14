package wsclient

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
	"time"

	"github.com/reservamos/search/internal/config"
	"github.com/reservamos/search/perform/models/responses"
)

// FlightsRequestArgs flights web service request params
type FlightsRequestArgs struct {
	Transporter string
	Db          time.Time
	De          time.Time
	Origin      string
	Destination string
	Passengers  []string
	Discount    FlightDiscount
}

// FlightDiscount is the discount format to send request params to integrations service
type FlightDiscount struct {
	Type     string
	Category string
	Code     string
}

// GetFlights performs request to integrations flight service
func GetFlights(args FlightsRequestArgs) (*responses.FlightsResponse, error) {
	var fr responses.FlightsResponse
	ws := NewWebService()
	URL := ws.URL
	queryParams, paramsError := args.String()
	if paramsError != nil {
		return &fr, paramsError
	}
	URL.Path = "airlines/search/" + args.Transporter
	URL.RawQuery = queryParams
	response, reqErr := ws.DoWSRequest("GET", URL.String(), nil)
	if reqErr != nil {
		return &fr, reqErr
	}
	defer response.Body.Close()
	body, readError := ioutil.ReadAll(response.Body)
	if readError != nil {
		return &fr, readError
	}

	json.Unmarshal(body, &fr)

	return &fr, nil
}

func (args FlightsRequestArgs) String() (string, error) {
	if args.Origin == "" || args.Destination == "" {
		return "", fmt.Errorf("error: origin and destination are required")
	}
	v := url.Values{}
	v.Set("db", args.Db.Format("02-01-2006"))
	v.Set("origin", args.Origin)
	v.Set("destination", args.Destination)

	var passengers string
	for _, p := range args.Passengers {
		passengers = passengers + "&passengers[]=" + p
	}

	if args.Discount.Code != "" {
		v.Set("discount.type", args.Discount.Type)
		v.Set("discount.category", args.Discount.Category)
		v.Set("discount.code", args.Discount.Code)
	}

	v.Set("currency", config.App.Currency)

	return v.Encode() + passengers, nil
}
