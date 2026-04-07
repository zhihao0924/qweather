package airquality

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"qweather/common"
	"qweather/internal/sdk"
)

type Service struct {
	client *sdk.Client
}

func New(client *sdk.Client) *Service {
	return &Service{client: client}
}

func (s *Service) Current(ctx context.Context, query common.CoordinateQuery) (*CurrentResponse, error) {
	var out CurrentResponse
	path := fmt.Sprintf("/airquality/v1/current/%s/%s", formatCoordinate(query.Latitude), formatCoordinate(query.Longitude))
	if err := s.client.DoJSON(ctx, http.MethodGet, path, sdk.NewParams().Set("lang", query.Lang), &out); err != nil {
		return nil, err
	}
	return &out, nil
}

func formatCoordinate(v float64) string {
	return strconv.FormatFloat(v, 'f', 2, 64)
}
