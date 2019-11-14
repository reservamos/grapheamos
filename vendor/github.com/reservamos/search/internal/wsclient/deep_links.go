package wsclient

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
	"time"
)

type deepLinkResponse struct {
	Response DeepLink `json:"body"`
	Status   int
}

// DeepLink string containing deep link structure
type DeepLink struct {
	Link string `json:"deep_link"`
}

// BusRedirectParams parameters to fetch redirection url
type BusRedirectParams struct {
	Origin          string
	Destination     string
	DepartsKey      string
	ReturnsKey      string
	DepartsDate     time.Time
	ReturnsDate     time.Time
	TransporterAbbr string
}

func (p BusRedirectParams) valid() bool {
	return p.Origin != "" && p.Destination != "" && p.DepartsKey != "" && !p.DepartsDate.IsZero()
}

// string stringified given parameters
func (p BusRedirectParams) string() (string, error) {
	v := url.Values{}
	if !p.valid() {
		return "", fmt.Errorf("Invalid Params")
	}
	v.Set("origin", p.Origin)
	v.Set("destination", p.Destination)
	v.Set("departs", fmt.Sprintf("%d", p.DepartsDate.Unix()*1000))
	v.Set("depart_key", p.DepartsKey)
	if !p.ReturnsDate.IsZero() && p.ReturnsKey != "" {
		v.Set("returns", fmt.Sprintf("%d", p.ReturnsDate.Unix()*1000))
		v.Set("return_key", p.ReturnsKey)
	}
	return v.Encode(), nil
}

// FlightRedirectParams parameters to fetch redirection url for a flight
type FlightRedirectParams struct {
	Origin      string
	Destination string
	Currency    string
	Passengers  []string
	DepartsKey  string
	ReturnsKey  string
	DepartsDate time.Time
	ReturnsDate time.Time
	AgentAbbr   string
}

func (p FlightRedirectParams) valid() bool {
	return p.Origin != "" && p.Destination != "" && p.DepartsKey != "" && !p.DepartsDate.IsZero() && len(p.Passengers) > 0 && p.Currency != ""
}

// string stringified given parameters
func (p FlightRedirectParams) string() (string, error) {
	v := url.Values{}
	if !p.valid() {
		return "", fmt.Errorf("Invalid Params")
	}
	v.Set("origin", p.Origin)
	v.Set("destination", p.Destination)
	v.Set("departs", p.DepartsDate.Format("2006-01-02T15:04:05"))
	v.Set("depart_key", p.DepartsKey)
	v.Set("currency", p.Currency)
	for _, val := range p.Passengers {
		v.Add("passengers[]", val)
	}
	if !p.ReturnsDate.IsZero() && p.ReturnsKey != "" {
		v.Set("returns", p.ReturnsDate.Format("2006-01-02T15:04:05.000"))
		v.Set("return_key", p.ReturnsKey)
	}
	return v.Encode(), nil
}

// GetBusDeepLink is the main request for fetching buses from the web service
func GetBusDeepLink(params BusRedirectParams) (*DeepLink, error) {
	var resp deepLinkResponse
	ws := NewWebService()

	URL := ws.URL
	queryParams, paramsError := params.string()
	if paramsError != nil {
		return &resp.Response, paramsError
	}
	URL.Path = "deep-link/" + params.TransporterAbbr
	URL.RawQuery = queryParams

	res, reqError := ws.DoWSRequest("GET", URL.String(), nil)
	if reqError != nil {
		return &resp.Response, reqError
	}

	defer res.Body.Close()
	body, readError := ioutil.ReadAll(res.Body)
	if readError != nil {
		return &resp.Response, readError
	}

	json.Unmarshal(body, &resp)
	return &resp.Response, nil
}

// GetFlightDeepLink gets deep link for flights
func GetFlightDeepLink(params FlightRedirectParams) (*DeepLink, error) {
	var resp deepLinkResponse
	ws := NewWebService()

	URL := ws.URL
	queryParams, paramsError := params.string()
	if paramsError != nil {
		return &resp.Response, paramsError
	}
	URL.Path = "airlines/deep-link/" + params.AgentAbbr
	URL.RawQuery = queryParams
	res, reqError := ws.DoWSRequest("GET", URL.String(), nil)
	if reqError != nil {
		return &resp.Response, reqError
	}

	defer res.Body.Close()
	body, readError := ioutil.ReadAll(res.Body)
	if readError != nil {
		return &resp.Response, readError
	}

	json.Unmarshal(body, &resp)
	return &resp.Response, nil
}
