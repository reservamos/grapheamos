package wsclient

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/reservamos/search/internal/config"
)

type GoogleMapsGeocodeRequestArgs struct {
	State   string
	City    string
	Country string
	Key     string
}

func NewGoogleMapsGeoCodeRequestArgs(city, state, country string) GoogleMapsGeocodeRequestArgs {
	return GoogleMapsGeocodeRequestArgs{
		State:   state,
		City:    city,
		Country: country,
		Key:     config.App.GoogleMapsKey,
	}
}

func (g GoogleMapsGeocodeRequestArgs) Method() string {
	return "geocode"
}

func (g GoogleMapsGeocodeRequestArgs) String() string {
	return "key=" + g.Key + "&address=" + htmlizeString(g.City) + "+" + htmlizeString(g.State) + "+" + htmlizeString(g.Country)
}

type GoogleMapsGeocodeResponse struct {
	Results []struct {
		Geometry struct {
			Location struct {
				Lat float64 `json:"lat"`
				Lng float64 `json:"lng"`
			} `json:"location"`
			Bounds struct {
				Northeast struct {
					Lat float64 `json:"lat"`
					Lng float64 `json:"lng"`
				} `json:"northeast"`
				Southwest struct {
					Lat float64 `json:"lat"`
					Lng float64 `json:"lng"`
				} `json:"southwest"`
			} `json:"bounds"`
		} `json:"geometry"`
	} `json:"results"`
	Status string `json:"status"`
}

type CityGeometryData struct {
	Location        RouteCoordinate
	NortheastBounds RouteCoordinate
	SouthwestBounds RouteCoordinate
}

func htmlizeString(s string) string {
	return strings.Replace(s, " ", "%20", -1)
}

type GoogleMapsGenericArgs interface {
	String() string
	Method() string
}

func GoogleMapsCommonRequest(args GoogleMapsGenericArgs) ([]byte, error) {
	mapsBaseURL := "https://maps.googleapis.com"
	ws := NewGenericWebService(mapsBaseURL, config.App.GoogleMapsKey)
	URL := ws.URL
	queryParams := args.String()
	URL.Path = fmt.Sprintf("maps/api/%s/json", args.Method())
	URL.RawQuery = queryParams

	response, reqErr := ws.DoWSRequest("GET", URL.String(), nil)
	log.Println("maps:", URL.String())
	if reqErr != nil {
		return nil, reqErr
	}
	defer response.Body.Close()
	rawResponse, readError := ioutil.ReadAll(response.Body)
	if readError != nil {
		return nil, readError
	}
	return rawResponse, nil
}

// GetGeocodeInformation performs request to integrations flight service
func GetGeocodeInformation(args GoogleMapsGeocodeRequestArgs) (*CityGeometryData, error) {
	body, err := GoogleMapsCommonRequest(args)
	if err != nil {
		return nil, err
	}
	var gmaps GoogleMapsGeocodeResponse
	json.Unmarshal(body, &gmaps)
	var res *CityGeometryData
	if gmaps.Status == "OK" && len(gmaps.Results) > 0 {
		res = &CityGeometryData{
			Location: RouteCoordinate{
				Lat: gmaps.Results[0].Geometry.Location.Lat,
				Lon: gmaps.Results[0].Geometry.Location.Lng,
			},
			NortheastBounds: RouteCoordinate{
				Lat: gmaps.Results[0].Geometry.Bounds.Northeast.Lat,
				Lon: gmaps.Results[0].Geometry.Bounds.Northeast.Lng,
			},
			SouthwestBounds: RouteCoordinate{
				Lat: gmaps.Results[0].Geometry.Bounds.Southwest.Lat,
				Lon: gmaps.Results[0].Geometry.Bounds.Northeast.Lng,
			},
		}
	} else {
		err = errors.New("Cannot parse response")
	}
	return res, err
}

func GetCitiesWithNoGeocodeData() {
	var cityAddresses []struct {
		City    string
		State   string
		Country string
	}
	config.SQL.Raw(`select c.name city, s.name state, m.name country
	FROM cities c
	join states s on s.id=c.state_id and (c.lat is null or c.long is null)
	join countries m on m.id=s.country_id`).Scan(&cityAddresses)

	for _, address := range cityAddresses {
		geo, err := GetGeocodeInformation(NewGoogleMapsGeoCodeRequestArgs(address.City, address.State, address.Country))
		if err != nil {
			log.Println(err)
			continue
		}
		tuple := fmt.Sprintf(`UPDATE cities set lat=%f, long=%f, bounds_ne_lat=%f, bounds_ne_long=%f, bounds_sw_lat=%f, bounds_sw_long=%f where name='%s' `,
			geo.Location.Lat, geo.Location.Lon, geo.NortheastBounds.Lat, geo.NortheastBounds.Lon,
			geo.SouthwestBounds.Lat, geo.SouthwestBounds.Lon, address.City)

		config.SQL.Exec(tuple)
	}
}
