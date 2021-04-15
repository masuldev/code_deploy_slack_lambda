package main

import (
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

const slackURL = ""

func Handler(request SNSMessage) error {
	err := request.Validate()
	if err != nil {
		return err
	}

	var message Message
	err = json.Unmarshal([]byte(request.Records[0].SNS.Message), &message)
	if err != nil {
		return ErrInvalidMessage
	}

	severity := changeSeverity(message)

	attachment := Attachment{}
	attachment.AddColor(severity)
	attachment.AddField(AttachmentField{Title: "Task", Value: message.EventTriggerName, Short: true})
	attachment.AddField(AttachmentField{Title: "Status", Value: message.Status, Short: true})
	attachment.AddField(AttachmentField{Title: "Application", Value: message.ApplicationName, Short: true})
	attachment.AddField(AttachmentField{Title: "Deployment Group", Value: message.DeploymentGroupName, Short: true})
	attachment.AddField(AttachmentField{Title: "Region", Value: message.Region, Short: true})
	attachment.AddField(AttachmentField{Title: "Deployment Link", Value: fmt.Sprintf("https://%s.console.aws.amazon.com/codedeploy/home?region=%s#/deployments/%s", message.Region, message.Region, message.DeploymentId), Short: true})
	attachment.AddField(AttachmentField{Title: "Create Time", Value: string(message.CreateTime), Short: true})
	if message.CompleteTime != "" {
		attachment.AddField(AttachmentField{Title: "Complete Time", Value: string(message.CompleteTime), Short: true})
	} else {
		attachment.AddField(AttachmentField{Title: "Complete Time", Value: "", Short: true})
	}
	attachment.AddField(AttachmentField{Title: "Error Code", Value: message.ErrorInformation.ErrorCode, Short: true})
	attachment.AddField(AttachmentField{Title: "Error Message", Value: message.ErrorInformation.ErrorMessage, Short: true})

	payload := Payload{
		Username:    "AWS Lambda :: Outframe :: CodeDeploy",
		Text:        request.Records[0].SNS.Subject,
		IconEmoji:   ":smile_cat:",
		Attachments: []Attachment{attachment},
	}

	err = Send(slackURL, payload)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	lambda.Start(Handler)
}
