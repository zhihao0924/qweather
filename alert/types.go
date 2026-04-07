package alert

import "github.com/zhihao0924/qweather/common"

type Query struct {
	common.CoordinateQuery
	LocalTime bool
}

type CurrentResponse struct {
	Metadata common.Metadata `json:"metadata"`
	Alerts   []Record        `json:"alerts"`
}

type Record struct {
	ID            string           `json:"id"`
	SenderName    string           `json:"senderName"`
	IssuedTime    string           `json:"issuedTime"`
	MessageType   MessageType      `json:"messageType"`
	LatestChange  string           `json:"latestChange"`
	EventType     EventType        `json:"eventType"`
	Urgency       string           `json:"urgency"`
	Severity      string           `json:"severity"`
	Certainty     string           `json:"certainty"`
	Icon          string           `json:"icon"`
	Color         common.RGBAColor `json:"color"`
	EffectiveTime string           `json:"effectiveTime"`
	OnsetTime     string           `json:"onsetTime"`
	ExpireTime    string           `json:"expireTime"`
	Headline      string           `json:"headline"`
	Description   string           `json:"description"`
	Criteria      string           `json:"criteria"`
	ResponseTypes []string         `json:"responseTypes"`
	Instruction   string           `json:"instruction"`
}

type MessageType struct {
	Code       string   `json:"code"`
	Supersedes []string `json:"supersedes"`
}

type EventType struct {
	Name string `json:"name"`
	Code string `json:"code"`
}
