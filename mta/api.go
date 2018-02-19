package mta

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/gregmankes/mta-alexa/models"
	"github.com/gregmankes/mta-alexa/transit_realtime"
	"github.com/renstrom/fuzzysearch/fuzzy"
)

type Feed int

const (
	OneThroughSixAndSFeed Feed = 1
	ACEFeed               Feed = 26
	NQRWFeed              Feed = 16
	BDFMFeed              Feed = 21
	LFeed                 Feed = 2
	GFeed                 Feed = 31
	JZFeed                Feed = 36
	UndefinedFeed         Feed = 0
)

const (
	endpoint     = "http://datamine.mta.info/mta_esi.php"
	stopEndpoint = "http://mtaapi.herokuapp.com/stations"
)

var (
	feedMap = map[string]Feed{
		"1": OneThroughSixAndSFeed,
		"2": OneThroughSixAndSFeed,
		"3": OneThroughSixAndSFeed,
		"4": OneThroughSixAndSFeed,
		"5": OneThroughSixAndSFeed,
		"6": OneThroughSixAndSFeed,
		"S": OneThroughSixAndSFeed,
		"A": ACEFeed,
		"C": ACEFeed,
		"E": ACEFeed,
		"N": NQRWFeed,
		"Q": NQRWFeed,
		"R": NQRWFeed,
		"W": NQRWFeed,
		"L": LFeed,
		"G": GFeed,
		"J": JZFeed,
		"Z": JZFeed,
	}
	client = http.Client{
		Timeout: time.Duration(10) * time.Second,
	}
	directionMap = map[string]string{
		"north":  "N",
		"sourth": "S",
	}
)

func getFeed(lineName string) Feed {
	return feedMap[strings.ToUpper(lineName)]
}

func getStops() (*models.SubwayStopResults, error) {
	req, err := http.NewRequest("GET", stopEndpoint, nil)
	if err != nil {
		return nil, newError(fmt.Sprintf("Error creating the subway stop request: %s", err), HTTPErrorType)
	}
	res, err := client.Do(req)
	if err != nil {
		return nil, newError(fmt.Sprintf("Error getting the subway stops: %s", err), HTTPErrorType)
	}
	defer res.Body.Close()
	buf, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, newError(fmt.Sprintf("Error reading the subway stops: %s", err), HTTPErrorType)
	}
	stopResults := &models.SubwayStopResults{}
	err = json.Unmarshal(buf, stopResults)
	if err != nil {
		return nil, newError(fmt.Sprintf("Error reading the subway stops: %s", err), HTTPErrorType)
	}
	return stopResults, nil
}

func GetFeedData(apiKey, lineName, stop, direction string) ([]time.Duration, error) {
	feedID := getFeed(lineName)
	if feedID == UndefinedFeed {
		return nil, newError(fmt.Sprintf("Feed for line name %s is undefined", lineName), FeedUndefinedErrorType)
	}
	stopResults, err := getStops()
	if err != nil {
		return nil, err
	}
	cm := getClosestFromStopResults(stopResults, stop)
	stopMap := generateStopMap(stopResults)
	stopID := stopMap[cm][directionMap[strings.ToLower(direction)]]
	lineStopMap, err := sendReq(apiKey, feedID)
	if err != nil {
		return nil, err
	}
	now := time.Now()
	log.Printf("Now: %s, StopID: %s", now.String(), stopID)
	return getDurations(now, lineStopMap[strings.ToUpper(lineName)][stopID]), nil
}

func getClosestFromStopResults(stopResults *models.SubwayStopResults, stop string) string {
	wordsToTest := []string{}
	for _, stopResult := range stopResults.Results {
		wordsToTest = append(wordsToTest, strings.ToUpper(stopResult.Name))
	}
	fmt.Println(wordsToTest)
	matches := fuzzy.Find(stop, wordsToTest)
	if len(matches) == 0 {
		return ""
	}
	return matches[0]
}

func getDurations(now time.Time, stopTimes []time.Time) []time.Duration {
	durations := []time.Duration{}
	for _, stopTime := range stopTimes {
		durations = append(durations, stopTime.Sub(now))
	}
	return durations
}

func sendReq(apiKey string, feedID Feed) (map[string]map[string][]time.Time, error) {
	u, _ := url.Parse(endpoint)
	q := u.Query()
	q.Set("key", apiKey)
	q.Set("feed_id", fmt.Sprintf("%d", feedID))
	u.RawQuery = q.Encode()
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, newError(fmt.Sprintf("Error creating the request: %s", err), HTTPErrorType)
	}
	res, err := client.Do(req)
	if err != nil {
		return nil, newError(fmt.Sprintf("Error sending the request: %s", err), HTTPErrorType)
	}
	defer res.Body.Close()
	buf, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, newError(fmt.Sprintf("Error reading the response body: %s", err), HTTPErrorType)
	}
	feed := transit_realtime.FeedMessage{}
	err = proto.Unmarshal(buf, &feed)
	if err != nil {
		return nil, newError(fmt.Sprintf("Error reading the feed: %s", err), FeedReadErrorType)
	}
	return generateLineStopTimeMap(feed), nil
}

func generateStopMap(stopResults *models.SubwayStopResults) map[string]map[string]string {
	stopMap := make(map[string]map[string]string)
	for _, stopResult := range stopResults.Results {
		if _, ok := stopMap[strings.ToUpper(stopResult.Name)]; !ok {
			stopMap[strings.ToUpper(stopResult.Name)] = make(map[string]string)
		}
		stopMap[strings.ToUpper(stopResult.Name)][string(stopResult.ID[len(stopResult.ID)-1])] = stopResult.ID
	}
	return stopMap
}

func generateLineStopTimeMap(feed transit_realtime.FeedMessage) map[string]map[string][]time.Time {
	lineStopTimeMap := make(map[string]map[string][]time.Time)
	for _, entity := range feed.Entity {
		if entity.TripUpdate != nil {
			if _, ok := lineStopTimeMap[entity.TripUpdate.Trip.GetRouteId()]; !ok {
				lineStopTimeMap[entity.TripUpdate.Trip.GetRouteId()] = make(map[string][]time.Time)
			}
			for _, stopTimeUpdate := range entity.TripUpdate.StopTimeUpdate {
				if len(lineStopTimeMap[entity.TripUpdate.Trip.GetRouteId()][stopTimeUpdate.GetStopId()]) == 0 {
					lineStopTimeMap[entity.TripUpdate.Trip.GetRouteId()][stopTimeUpdate.GetStopId()] = []time.Time{
						time.Unix(stopTimeUpdate.GetArrival().GetTime(), 0),
					}
				} else {
					lineStopTimeMap[entity.TripUpdate.Trip.GetRouteId()][stopTimeUpdate.GetStopId()] = append(
						lineStopTimeMap[entity.TripUpdate.Trip.GetRouteId()][stopTimeUpdate.GetStopId()],
						time.Unix(stopTimeUpdate.GetArrival().GetTime(), 0),
					)
				}
			}
		}
	}
	for _, lineStopMap := range lineStopTimeMap {
		for _, lineStopTimes := range lineStopMap {
			sort.Slice(lineStopTimes, func(i, j int) bool {
				return lineStopTimes[i].Unix() < lineStopTimes[j].Unix()
			})
		}
	}
	return lineStopTimeMap
}
