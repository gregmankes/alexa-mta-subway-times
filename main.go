package main

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

var apiKey string

func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{}, nil
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
	lambda.Start(Handler)
}
