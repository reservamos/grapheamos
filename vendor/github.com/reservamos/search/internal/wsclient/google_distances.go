package wsclient

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/reservamos/search/internal/config"
)

/*
{
   "destination_addresses" : [ "New York, NY, USA" ],
   "origin_addresses" : [ "Washington, DC, USA" ],
   "rows" : [
      {
         "elements" : [
            {
               "distance" : {
                  "text" : "225 mi",
                  "value" : 361715
               },
               "duration" : {
                  "text" : "3 hours 49 mins",
                  "value" : 13725
               },
               "status" : "OK"
            }
         ]
      }
   ],
   "status" : "OK"
}
*/

// GoogleMapsRequestArgs distance-matrix web service request params
type GoogleMapsDistanceMatrixRequestArgs struct {
	Origin       RouteCoordinate
	Destinations []RouteCoordinate
	Key          string
}

type GoogleMapsDistanceMatrixResponse struct {
	Rows []struct {
		Elements []struct {
			Distance struct {
				Value int `json:"value"`
			} `json:"distance"`
			Duration struct {
				Seconds int `json:"value"`
			} `json:"duration"`
			Status string `json:"status"`
		} `json:"elements"`
	} `json:"rows"`
}

// GetDuration performs request to integrations flight service
func GetGMatrixInformation(args GoogleMapsDistanceMatrixRequestArgs) ([]int, []int, error) {
	var gmaps GoogleMapsDistanceMatrixResponse
	mapsBaseURL := "https://maps.googleapis.com"
	ws := NewGenericWebService(mapsBaseURL, args.Key)
	URL := ws.URL
	queryParams, paramsError := args.String()
	if paramsError != nil {
		return nil, nil, paramsError
	}
	URL.Path = "maps/api/distancematrix/json"
	URL.RawQuery = queryParams
	response, reqErr := ws.DoWSRequest("GET", URL.String(), nil)
	log.Println("maps:", URL.String())
	if reqErr != nil {
		return nil, nil, reqErr
	}
	defer response.Body.Close()
	body, readError := ioutil.ReadAll(response.Body)
	if readError != nil {
		return nil, nil, readError
	}

	json.Unmarshal(body, &gmaps)

	durations, distances := gmaps.extractGMatrixData()
	return durations, distances, nil
}

func (g *GoogleMapsDistanceMatrixResponse) extractGMatrixData() ([]int, []int) {
	var durations []int
	var distances []int

	for _, r := range g.Rows {
		for _, e := range r.Elements {
			if e.Status == "OK" {
				durations = append(durations, e.Duration.Seconds/60)
				distances = append(distances, e.Distance.Value/1000)
			} else {
				durations = append(durations, -1)
				distances = append(distances, -1)
			}
		}
	}
	return durations, distances
}

func (args GoogleMapsDistanceMatrixRequestArgs) String() (string, error) {

	origin, oErr := args.Origin.String()

	var destinations []string
	for _, destination := range args.Destinations {
		tmp, dErr := destination.String()
		if dErr == nil {
			destinations = append(destinations, tmp)
		}
	}

	if oErr != nil || len(destinations) == 0 {
		return "", fmt.Errorf("error: origin and destination are required")
	}

	year, month, day := time.Now().AddDate(0, 0, 1).Date()
	currentTimeZone, _ := time.LoadLocation(config.App.Timezone)
	unixPeakTrafficTime := time.Date(year, month, day, 18, 0, 0, 0, currentTimeZone).Unix()

	v := url.Values{}
	v.Add("traffic_model", "pessimistic")
	v.Add("departure_time", strconv.FormatInt(unixPeakTrafficTime, 10))
	v.Add("key", args.Key)
	v.Add("origins", origin)
	v.Add("destinations", strings.Join(destinations, "|"))

	return v.Encode(), nil
}
