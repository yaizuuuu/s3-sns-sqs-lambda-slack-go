package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/yaizuuuu/s3-sns-sqs-lambda-slack-go/handlers/notifier/slack"
	"os"
)

var client *slack.Client

func main() {
	lambda.Start(handler)
}

func init() {
	client = slack.NewClient(
		slack.Config{
			URL: os.Getenv("WEBHOOK_URL"),
			Channel: os.Getenv("CHANNEL"),
			Username: os.Getenv("USER_NAME"),
			IconEmoji: os.Getenv("ICON"),
		},
	)
}

func handler(snsEvent events.SNSEvent) error {
	record := snsEvent.Records[0]
	snsRecord := snsEvent.Records[0].SNS
	fmt.Printf("[%s %s] Message = %s \n", record.EventSource, snsRecord.Timestamp, snsRecord)

	if err := client.PostMessage(snsRecord.Message); err != nil {
		return err
	}

	return nil
}
