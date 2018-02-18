package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/aws/aws-lambda-go/lambda"
	alexa "github.com/ericdaugherty/alexa-skills-kit-golang"
	"github.com/gregmankes/mta-alexa/handlers"
)

var (
	apiKey string
	axa    alexa.Alexa
)

func Handler(ctx context.Context, requestEnv *alexa.RequestEnvelope) (interface{}, error) {
	return axa.ProcessRequest(requestEnv)
}

func main() {
	buf, err := ioutil.ReadFile("config.json")
	if err != nil {
		log.Fatalf("Error loading config file: %s", err)
		return
	}
	cfg := map[string]interface{}{}
	err = json.Unmarshal(buf, &cfg)
	if err != nil {
		log.Fatalf("Error unmarshaling the config file")
		return
	}
	apiKey = cfg["apiKey"].(string)
	applicationID := cfg["applicationId"].(string)
	axa = alexa.Alexa{
		ApplicationID: applicationID,
		RequestHandler: handlers.CommuteTimes{
			APIKey: apiKey,
		},
	}
	lambda.Start(Handler)
}
