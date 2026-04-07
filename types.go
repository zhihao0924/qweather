package qweather

type codedResponse interface {
	getCode() string
}

type baseResponse struct {
	Code string `json:"code"`
}

func (r baseResponse) getCode() string {
	return r.Code
}

type Refer struct {
	Sources []string `json:"sources"`
	License []string `json:"license"`
}

type Location struct {
	Name      string `json:"name"`
	ID        string `json:"id"`
	Lat       string `json:"lat"`
	Lon       string `json:"lon"`
	Adm2      string `json:"adm2"`
	Adm1      string `json:"adm1"`
	Country   string `json:"country"`
	TZ        string `json:"tz"`
	UTCOffset string `json:"utcOffset"`
	IsDST     string `json:"isDst"`
	Type      string `json:"type"`
	Rank      string `json:"rank"`
	FxLink    string `json:"fxLink"`
}

type Metadata struct {
	Tag          string   `json:"tag"`
	ZeroResult   bool     `json:"zeroResult"`
	Attributions []string `json:"attributions"`
}

type RGBAColor struct {
	Code  string  `json:"code,omitempty"`
	Red   int     `json:"red"`
	Green int     `json:"green"`
	Blue  int     `json:"blue"`
	Alpha float64 `json:"alpha"`
}
