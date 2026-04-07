package sdk

import (
	"net/url"
	"strconv"
	"strings"
)

type Params struct {
	values url.Values
}

func NewParams() Params {
	return Params{values: make(url.Values)}
}

func (p Params) Set(key, value string) Params {
	if strings.TrimSpace(value) == "" {
		return p
	}
	p.values.Set(key, value)
	return p
}

func (p Params) SetInt(key string, value int) Params {
	if value <= 0 {
		return p
	}
	p.values.Set(key, strconv.Itoa(value))
	return p
}

func (p Params) SetBool(key string, value bool) Params {
	p.values.Set(key, strconv.FormatBool(value))
	return p
}
