package main

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type AttachmentField struct {
	Title string `json:"title"`
	Value string `json:"value"`
	Short bool `json:"short"`
}

type AttachmentAction struct {
	Type string `json:"type"`
	Text string `json:"text"`
	Url string `json:"url"`
}

type Attachment struct {
	Fallback     *string   `json:"fallback,omitempty"`
	Color        *string   `json:"color,omitempty"`

	CallbackID   *string   `json:"callback_id,omitempty"`

	AuthorName   *string   `json:"author_name,omitempty"`
	AuthorLink   *string   `json:"author_link,omitempty"`
	AuthorIcon   *string   `json:"author_icon,omitempty"`

	PreText      *string   `json:"pretext,omitempty"`
	Title        *string   `json:"title,omitempty"`
	TitleLink    *string   `json:"title_link,omitempty"`
	Text         *string   `json:"text,omitempty"`

	ImageURL     *string             `json:"image_url,omitempty"`
	Fields       []*AttachmentField  `json:"fields,omitempty"`
	TS    *int64                     `json:"ts,omitempty"`
	MarkdownIn   *[]string           `json:"mrkdwn_in,omitempty"`
	Actions      []*AttachmentAction `json:"actions,omitempty"`

	Footer       *string   `json:"footer,omitempty"`
	FooterIcon   *string   `json:"footer_icon,omitempty"`
}

type Payload struct {
	Username        string       `json:"username,omitempty"`
	IconEmoji       string       `json:"icon_emoji,omitempty"`
	IconURL         string       `json:"icon_url,omitempty"`
	Channel         string       `json:"channel,omitempty"`
	ThreadTimestamp string       `json:"thread_ts,omitempty"`
	Text            string       `json:"text,omitempty"`
	Attachments     []Attachment `json:"attachments,omitempty"`
	Parse           string       `json:"parse,omitempty"`
}

func (a *Attachment) AddField(field AttachmentField) *Attachment {
	a.Fields = append(a.Fields, &field)
	return a
}

func (a *Attachment) AddColor(color string) *Attachment {
	a.Color = &color
	return a
}

func (a *Attachment) AddAction(action AttachmentAction) *Attachment {
	a.Actions = append(a.Actions, &action)
	return a
}

func Send(slackURL string, payload Payload) error {
	payloadJson, _ := json.Marshal(payload)
	resp, err := http.Post(slackURL, "application/json", bytes.NewBuffer(payloadJson))
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return ErrStatusCode
	}

	return nil
}