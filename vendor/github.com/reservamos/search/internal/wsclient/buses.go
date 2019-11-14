package wsclient

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
	"time"

	"github.com/reservamos/search/perform/models/responses"
)

// BusesRequestArgs are the arguments needed to perform a request to the integrations WS
type BusesRequestArgs struct {
	Transporter           string
	Origin                string
	Destination           string
	Departure             time.Time
	OriginCoordinate      RouteCoordinate
	DestinationCoordinate RouteCoordinate
}

// GetBuses is the main request for fetching buses from the web service
func GetBuses(args BusesRequestArgs) (*responses.BusesResponse, error) {
	var resp responses.BusesResponse
	ws := NewWebService()

	URL := ws.URL
	queryParams, paramsError := args.String()
	if paramsError != nil {
		return &resp, paramsError
	}
	URL.Path = "trips/" + args.Transporter
	URL.RawQuery = queryParams

	res, reqError := ws.DoWSRequest("GET", URL.String(), nil)
	if reqError != nil {
		return &resp, reqError
	}

	defer res.Body.Close()
	body, readError := ioutil.ReadAll(res.Body)
	if readError != nil {
		return &resp, readError
	}

	json.Unmarshal(body, &resp)

	return &resp, nil
}

func (args BusesRequestArgs) String() (string, error) {
	if args.Origin == "" || args.Destination == "" {
		return "", fmt.Errorf("error: origin and destination are required")
	}
	v := url.Values{}
	v.Set("departure", fmt.Sprintf("%d", args.Departure.Unix()*1000))
	v.Set("origin", args.Origin)
	v.Set("destination", args.Destination)

	if (args.OriginCoordinate != RouteCoordinate{} && args.DestinationCoordinate != RouteCoordinate{}) {
		origin, errO := args.OriginCoordinate.String()
		if errO != nil {
			return "", errO
		}
		v.Set("route_coordinates.origin", origin)

		destination, errD := args.DestinationCoordinate.String()
		if errD != nil {
			return "", errD
		}
		v.Set("route_coordinates.destination", destination)
	}

	return v.Encode(), nil
}

func (c RouteCoordinate) String() (string, error) {
	if c.Lat == 0 || c.Lon == 0 {
		return "", fmt.Errorf("error: Lat and Lon are required")
	}
	return fmt.Sprintf("%f", c.Lat) + "," + fmt.Sprintf("%f", c.Lon), nil
}
