package wsclient

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/url"
)

// Rome2RioRequestArgs rome2rio rest api request params
type Rome2RioRequestArgs struct {
	Origin       RouteCoordinate
	Destination  RouteCoordinate
	Key          string
	CurrencyCode string
}

type Rome2RioResponse struct {
	Routes []struct {
		Name          string             `json:"name"`
		Distance      float64            `json:"distance"`
		TotalDuration int                `json:"totalDuration"`
		Prices        []IndicativePrices `json:"indicativePrices"`
	} `json:"routes"`
}

type IndicativePrices struct {
	Name      string `json:"name"`
	Price     int    `json:"price"`
	PriceLow  int    `json:"priceLow"`
	PriceHigh int    `json:"priceHigh"`
	Currency  string `json:"currency"`
}

func (response *Rome2RioResponse) WalkDistanceAndDuration() (float64, int, error) {
	for _, route := range response.Routes {
		if route.Name == "Walk" {
			return route.Distance, route.TotalDuration, nil
		}
	}
	return 0, 0, fmt.Errorf("Walking route not found")
}

func (response *Rome2RioResponse) UberPrice() (*IndicativePrices, float64, error) {
	for i := range response.Routes {
		if response.Routes[i].Name == "Uber" {
			for j := range response.Routes[i].Prices {
				if response.Routes[i].Prices[j].Name == "UberX" {
					return &response.Routes[i].Prices[j], response.Routes[i].Distance, nil
				}
			}
		}
	}
	return nil, 0, fmt.Errorf("UberX price not found")
}

func (response *Rome2RioResponse) TaxiPrice() (*IndicativePrices, float64, error) {
	for i := range response.Routes {
		if response.Routes[i].Name == "Taxi" {
			if len(response.Routes[i].Prices) > 0 {
				return &response.Routes[i].Prices[0], response.Routes[i].Distance, nil
			}
		}
	}
	return nil, 0, fmt.Errorf("Taxi price not found")
}

// GetPrices performs request to rome2rio api
func GetPrices(args Rome2RioRequestArgs) (*Rome2RioResponse, error) {
	var serviceResponse Rome2RioResponse
	rome2rioBaseURL := "http://free.rome2rio.com"
	ws := NewGenericWebService(rome2rioBaseURL, args.Key)
	URL := ws.URL
	queryParams, paramsError := args.String()
	if paramsError != nil {
		return &serviceResponse, paramsError
	}
	URL.Path = "api/1.4/json/Search"
	URL.RawQuery = queryParams
	log.Println("rome2rio:", URL.String())
	response, reqErr := ws.DoWSRequest("GET", URL.String(), nil)
	if reqErr != nil {
		return &serviceResponse, reqErr
	}
	defer response.Body.Close()
	body, readError := ioutil.ReadAll(response.Body)
	if readError != nil {
		return &serviceResponse, readError
	}

	json.Unmarshal(body, &serviceResponse)
	log.Println(serviceResponse)
	return &serviceResponse, nil
}

func (args Rome2RioRequestArgs) String() (string, error) {
	origin, oErr := args.Origin.String()
	destination, dErr := args.Destination.String()
	if oErr != nil || dErr != nil {
		return "", fmt.Errorf("error: origin and destination are required")
	}
	v := url.Values{}
	v.Set("currencyCode", args.CurrencyCode)
	v.Set("key", args.Key)
	v.Set("oPos", origin)
	v.Set("dPos", destination)

	return v.Encode(), nil
}
