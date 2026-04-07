package geo

import "qweather/common"

type CityLookupParams struct {
	Location string
	Adm      string
	Range    string
	Number   int
	Lang     string
}

type TopCitiesParams struct {
	Range  string
	Number int
	Lang   string
}

type CityLookupResponse struct {
	common.BaseResponse
	Location []common.Location `json:"location"`
	Refer    common.Refer      `json:"refer"`
}

type TopCitiesResponse struct {
	common.BaseResponse
	TopCityList []common.Location `json:"topCityList"`
	Refer       common.Refer      `json:"refer"`
}
