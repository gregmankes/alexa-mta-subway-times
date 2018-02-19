package handlers

import (
	"fmt"
	"log"
	"strings"

	alexa "github.com/ericdaugherty/alexa-skills-kit-golang"
	"github.com/gregmankes/mta-alexa/mta"
)

const (
	cardTitle          = "NYC Commute"
	directionSlot      = "direction"
	stopSlot           = "stop"
	subwaySlot         = "subway"
	maxSubwayDurations = 3
)

type CommuteTimes struct {
	APIKey string
}

// OnSessionStarted called when a new session is created.
func (c *CommuteTimes) OnSessionStarted(request *alexa.Request, session *alexa.Session, response *alexa.Response) error {
	log.Printf("OnSessionStarted requestId=%s, sessionId=%s", request.RequestID, session.SessionID)
	return nil
}

func (c *CommuteTimes) OnLaunch(request *alexa.Request, session *alexa.Session, response *alexa.Response) error {
	log.Printf("OnLaunch requestId=%s, sessionId=%s", request.RequestID, session.SessionID)

	speechText := "Welcome to commute time! You can say: ask commute for the enter_subway_stop enter_subway_line going enter_subway_line. Or you can just say: ask commute enter_subway_stop enter_subway_line enter_subway_line."
	response.SetSimpleCard(cardTitle, speechText)
	response.SetOutputText(speechText)
	response.SetRepromptText(speechText)

	response.ShouldSessionEnd = true

	return nil
}

func (c *CommuteTimes) OnIntent(request *alexa.Request, session *alexa.Session, response *alexa.Response) error {
	log.Printf("OnIntent requestId=%s, sessionId=%s, intent=%s", request.RequestID, session.SessionID, request.Intent.Name)
	switch request.Intent.Name {
	case "CommuteIntent":
		log.Println("CommuteIntent triggered")
		speechText, err := c.formatCommuteIntentOutput(request.Intent)
		if err != nil {
			return fmt.Errorf("Error getting mta results: %s", err)
		}
		response.SetSimpleCard(cardTitle, speechText)
		response.SetOutputText(speechText)

		log.Printf("Set Output speech, value now: %s", response.OutputSpeech.Text)
	case "AMAZON.HelpIntent":
		log.Println("AMAZON.HelpIntent triggered")
		// speechText := "You can say hello to me!"

		// response.SetSimpleCard("HelloWorld", speechText)
		// response.SetOutputText(speechText)
		// response.SetRepromptText(speechText)
	default:
		return fmt.Errorf("Invalid Intent")
	}
	return nil
}

func (c *CommuteTimes) formatCommuteIntentOutput(intent alexa.Intent) (string, error) {
	direction, ok := intent.Slots[directionSlot]
	if !ok {
		return "", fmt.Errorf("Error getting the direction")
	}
	stop, ok := intent.Slots[stopSlot]
	if !ok {
		return "", fmt.Errorf("Error getting the stop")
	}
	subway, ok := intent.Slots[subwaySlot]
	if !ok {
		return "", fmt.Errorf("Error getting the subway")
	}
	if subway.Value == "" || stop.Value == "" || direction.Value == "" {
		return "", fmt.Errorf("Error: not enough values for getting the mta subway times")
	}
	log.Printf("Subway: %s, Stop: %s, Direction %s", subway.Value, stop.Value, direction.Value)
	durations, err := mta.GetFeedData(c.APIKey, subway.Value, stop.Value, direction.Value)
	if err != nil {
		return "", fmt.Errorf("Error getting the subway durations: %s", err)
	}
	durationStrings := []string{}
	for i := 0; i < len(durations) && i < maxSubwayDurations; i++ {
		durationStrings = append(durationStrings, fmt.Sprintf("%1.2f", durations[i].Minutes()))
	}
	minuteString := strings.Join(durationStrings, ", ")
	if minuteString == "" {
		return "", fmt.Errorf("Error getting the subway durations: no durations returned from mta")
	}
	responseString := fmt.Sprintf("The next %d %s trains at the %s stop going %s are coming in %s minutes",
		maxSubwayDurations,
		subway.Value,
		stop.Value,
		direction.Value,
		minuteString,
	)
	return responseString, nil
}

func (c *CommuteTimes) OnSessionEnded(request *alexa.Request, session *alexa.Session, response *alexa.Response) error {
	log.Printf("OnSessionEnded requestId=%s, sessionId=%s", request.RequestID, session.SessionID)
	return nil
}
