package qweather

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
)

type WeatherAlertCurrentResponse struct {
	Metadata Metadata             `json:"metadata"`
	Alerts   []WeatherAlertRecord `json:"alerts"`
}

type WeatherAlertRecord struct {
	ID            string                  `json:"id"`
	SenderName    string                  `json:"senderName"`
	IssuedTime    string                  `json:"issuedTime"`
	MessageType   WeatherAlertMessageType `json:"messageType"`
	LatestChange  string                  `json:"latestChange"`
	EventType     WeatherAlertEventType   `json:"eventType"`
	Urgency       string                  `json:"urgency"`
	Severity      string                  `json:"severity"`
	Certainty     string                  `json:"certainty"`
	Icon          string                  `json:"icon"`
	Color         RGBAColor               `json:"color"`
	EffectiveTime string                  `json:"effectiveTime"`
	OnsetTime     string                  `json:"onsetTime"`
	ExpireTime    string                  `json:"expireTime"`
	Headline      string                  `json:"headline"`
	Description   string                  `json:"description"`
	Criteria      string                  `json:"criteria"`
	ResponseTypes []string                `json:"responseTypes"`
	Instruction   string                  `json:"instruction"`
}

type WeatherAlertMessageType struct {
	Code       string   `json:"code"`
	Supersedes []string `json:"supersedes"`
}

type WeatherAlertEventType struct {
	Name string `json:"name"`
	Code string `json:"code"`
}

func (c *Client) WeatherAlertCurrent(ctx context.Context, latitude, longitude float64, localTime bool, lang string) (*WeatherAlertCurrentResponse, error) {
	var out WeatherAlertCurrentResponse
	path := fmt.Sprintf("/weatheralert/v1/current/%s/%s", formatCoordinate(latitude), formatCoordinate(longitude))
	if err := c.doJSON(ctx, http.MethodGet, path, map[string]string{
		"localTime": strconv.FormatBool(localTime),
		"lang":      lang,
	}, &out); err != nil {
		return nil, err
	}
	return &out, nil
}
