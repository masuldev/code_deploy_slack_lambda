package main

import (
	"log"
	"strings"
	"time"
)

type SNSMessage struct {
	Records []Record `json:"Records"`
}

type Record struct {
	EventSource          string `json:"EventSource"`
	EventVersion         string `json:"EventVersion"`
	EventSubscriptionArn string `json:"EventSubscriptionArn"`
	SNS                  SNS    `json:"Sns"`
}

type SNS struct {
	Type             string    `json:"Type"`
	MessageId        string    `json:"MessageId"`
	TopicArn         string    `json:"TopicArn"`
	Subject          string    `json:"Subject"`
	Message          string    `json:"Message"`
	Timestamp        time.Time `json:"Timestamp"`
	SignatureVersion string    `json:"SignatureVersion"`
	Signature        string    `json:"Signature"`
	SigningCertURL   string    `json:"SigningCertUrl"`
	UnsubscribeURL   string    `json:"UnsubscribeUrl"`
}

type Message struct {
	Region              string             `json:"region,omitempty"`
	AccountId           string             `json:"accountId,omitempty"`
	EventTriggerName    string             `json:"eventTriggerName,omitempty"`
	ApplicationName     string             `json:"applicationName,omitempty"`
	DeploymentId        string             `json:"deploymentId,omitempty"`
	DeploymentGroupName string             `json:"deploymentGroupName,omitempty"`
	CreateTime          string             `json:"createTime,omitempty"`
	CompleteTime        string             `json:"completeTime,omitempty"`
	DeploymentOverview  DeploymentOverview `json:"deploymentOverview,omitempty"`
	Status              string             `json:"status,omitempty"`
	ErrorInformation    ErrorInformation   `json:"errorInformation,omitempty"`
}

type DeploymentOverview struct {
	Succeeded  string `json:"Succeeded,omitempty"`
	Failed     string `json:"Failed,omitempty"`
	Skipped    string `json:"Skipped,omitempty"`
	InProgress string `json:"InProgress,omitempty"`
	Pending    string `json:"Pending,omitempty"`
}

type ErrorInformation struct {
	ErrorCode    string `json:"ErrorCode,omitempty"`
	ErrorMessage string `json:"ErrorMessage,omitempty"`
}

func (m SNSMessage) Validate() error {
	switch {
	case len(m.Records) != 1:
		return ErrInvalidRecords
	case m.Records[0].SNS.Subject == "":
		return ErrInvalidSubject
	default:
		return nil
	}

}

func toKST(date string) time.Time {
	t, err := time.Parse(time.UnixDate, date)
	if err != nil {
		log.Println(ErrTimeParse)
	}

	loc, _ := time.LoadLocation("Asia/Seoul")
	kst := t.In(loc)
	return kst
}

func changeSeverity(message Message) string {
	severity := "good"
	var dangerMessages = []string{
		" but with errors",
		" to RED",
		"During an aborted deployment",
		"FAILED",
		"Failed to deploy application",
		"Failed to deploy configuration",
		"has a dependent object",
		"is not authorized to perform",
		"Pending to Degraded",
		"Stack deletion failed",
		"Unsuccessful command execution",
		"You do not have permision",
		"Your quota allows for 0 more running instance",
	}

	var warningMessages = []string{
		" aborted operation.",
		" to YELLOW",
		"Adding instance ",
		"Degraded to Info",
		"Deleting SNS topic",
		"is currently running under desired capacity",
		"Ok to Info",
		"Ok to Warning",
		"Pending Initialization",
		"Removed instance ",
		"Rollback of environment",
	}

	for _, item := range dangerMessages {
		if (strings.Contains(message.Status, item) && strings.Contains(message.ErrorInformation.ErrorMessage, item) && strings.Contains(message.ErrorInformation.ErrorCode, item)) == true {
			severity = "danger"
			break
		}
	}

	if severity == "good" {
		for _, item := range warningMessages {
			if (strings.Contains(message.Status, item) && strings.Contains(message.ErrorInformation.ErrorMessage, item) && strings.Contains(message.ErrorInformation.ErrorCode, item)) == true {
				severity = "warning"
				break
			}
		}
	}
	return severity
}
