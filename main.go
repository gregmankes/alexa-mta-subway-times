package main

import (
	"context"
	"log"
	"os"

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
	apiKey = os.Getenv("API_KEY")
	if apiKey == "" {
		log.Fatalf("Error getting the api key")
	}
	applicationID := os.Getenv("APPLICATION_ID")
	if applicationID == "" {
		log.Fatalf("Error getting the application ID")
	}
	axa = alexa.Alexa{
		ApplicationID: applicationID,
		RequestHandler: &handlers.CommuteTimes{
			APIKey: apiKey,
		},
	}
	lambda.Start(Handler)
}
