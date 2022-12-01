package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/api/v2010"
)

var (
	accountSid string
	authToken string
	fromPhone string
	toPhone string
	client *twilio.RestClient
)

func sendMessage(msg string) {
	params := openapi.CreateMessageParams{}
	params.SetTo(toPhone)
	params.SetFrom(fromPhone)
	params.SetBody(msg)

	response, err := client.Api.CreateMessage(&params)

	if err != nil {
		fmt.Printf("Error whilst sending message: %s\n", err.Error())
		return
	}

	fmt.Printf("Message SID: %s\n", *response.Sid)	
}

func Init() {
	err := godotenv.Load(".env")

	if err != nil {
		fmt.Printf("Error loading .env: %s\n", err.Error())
	}

	accountSid = os.Getenv("ACCOUNT_SID")
	authToken = os.Getenv("AUTH_TOKEN")
	fromPhone = os.Getenv("FROM_PHONE")
	toPhone = os.Getenv("TO_PHONE")

	client = twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: accountSid,
		Password: authToken,
	})
}

func main() {
	Init();
	msg := "Hello, sending from your stupid little go demo app, hope this works"
	sendMessage(msg)
}