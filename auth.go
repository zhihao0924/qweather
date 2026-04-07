package qweather

import (
	"crypto"

	qwauth "github.com/zhihao0924/qweather/auth"
)

type JWTConfig = qwauth.JWTConfig
type JWTTokenProvider = qwauth.TokenProvider

func NewJWTTokenProvider(cfg JWTConfig) (*JWTTokenProvider, error) {
	return qwauth.NewTokenProvider(cfg)
}

func ParseEd25519PrivateKey(pemBytes []byte) (crypto.PrivateKey, error) {
	return qwauth.ParseEd25519PrivateKey(pemBytes)
}
