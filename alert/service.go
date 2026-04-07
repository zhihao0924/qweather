package alert

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/zhihao0924/qweather/internal/sdk"
)

type Service struct {
	client *sdk.Client
}

func New(client *sdk.Client) *Service {
	return &Service{client: client}
}

func (s *Service) Current(ctx context.Context, query Query) (*CurrentResponse, error) {
	var out CurrentResponse
	path := fmt.Sprintf("/weatheralert/v1/current/%s/%s", formatCoordinate(query.Latitude), formatCoordinate(query.Longitude))
	params := sdk.NewParams().
		SetBool("localTime", query.LocalTime).
		Set("lang", query.Lang)
	if err := s.client.DoJSON(ctx, http.MethodGet, path, params, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

func formatCoordinate(v float64) string {
	return strconv.FormatFloat(v, 'f', 2, 64)
}
