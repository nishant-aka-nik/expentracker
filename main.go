package main

import (
	"context"
	"expentracker/gateway"
	"log"

	"google.golang.org/api/gmail/v1"
)

func main() {
	ctx := context.Background()
	gmailService := getGmailService(ctx)
	log.Println("gmail serivce = ", gmailService)

	//hello
}

func getGmailService(ctx context.Context) *gmail.Service {
	gatewayClient := gateway.NewGateway(ctx)
	gmailService, err := gatewayClient.GmailClient()
	if err != nil {
		log.Println("Error = ", err.Error())
	}
	return gmailService
}
