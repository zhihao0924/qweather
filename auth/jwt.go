package auth

import (
	"context"
	"crypto"
	"crypto/ed25519"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"sync"
	"time"
)

const (
	defaultTokenTTL    = 15 * time.Minute
	defaultTokenSkew   = 30 * time.Second
	maxTokenTTL        = 24 * time.Hour
	tokenRefreshWindow = 1 * time.Minute
)

type JWTConfig struct {
	CredentialID  string
	ProjectID     string
	PrivateKeyPEM []byte
	TTL           time.Duration
	Now           func() time.Time
}

type TokenProvider struct {
	credentialID string
	projectID    string
	privateKey   ed25519.PrivateKey
	ttl          time.Duration
	now          func() time.Time

	mu          sync.Mutex
	cachedToken string
	expiresAt   time.Time
}

func NewTokenProvider(cfg JWTConfig) (*TokenProvider, error) {
	if cfg.CredentialID == "" {
		return nil, fmt.Errorf("qweather: credential id is required")
	}
	if cfg.ProjectID == "" {
		return nil, fmt.Errorf("qweather: project id is required")
	}
	if len(cfg.PrivateKeyPEM) == 0 {
		return nil, fmt.Errorf("qweather: private key is required")
	}

	privateKey, err := parseEd25519PrivateKey(cfg.PrivateKeyPEM)
	if err != nil {
		return nil, err
	}

	ttl := cfg.TTL
	switch {
	case ttl <= 0:
		ttl = defaultTokenTTL
	case ttl > maxTokenTTL:
		return nil, fmt.Errorf("qweather: jwt ttl must be <= %s", maxTokenTTL)
	}

	nowFn := cfg.Now
	if nowFn == nil {
		nowFn = time.Now
	}

	return &TokenProvider{
		credentialID: cfg.CredentialID,
		projectID:    cfg.ProjectID,
		privateKey:   privateKey,
		ttl:          ttl,
		now:          nowFn,
	}, nil
}

func (p *TokenProvider) Token(context.Context) (string, error) {
	now := p.now().UTC()

	p.mu.Lock()
	defer p.mu.Unlock()

	if p.cachedToken != "" && now.Before(p.expiresAt.Add(-tokenRefreshWindow)) {
		return p.cachedToken, nil
	}

	token, exp, err := p.sign(now)
	if err != nil {
		return "", err
	}

	p.cachedToken = token
	p.expiresAt = exp
	return token, nil
}

func (p *TokenProvider) sign(now time.Time) (string, time.Time, error) {
	iat := now.Add(-defaultTokenSkew).Unix()
	exp := now.Add(p.ttl).Unix()

	header := map[string]string{
		"alg": "EdDSA",
		"kid": p.credentialID,
	}
	payload := map[string]any{
		"sub": p.projectID,
		"iat": iat,
		"exp": exp,
	}

	headerBytes, err := json.Marshal(header)
	if err != nil {
		return "", time.Time{}, fmt.Errorf("qweather: encode jwt header: %w", err)
	}
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return "", time.Time{}, fmt.Errorf("qweather: encode jwt payload: %w", err)
	}

	encodedHeader := base64.RawURLEncoding.EncodeToString(headerBytes)
	encodedPayload := base64.RawURLEncoding.EncodeToString(payloadBytes)
	signingInput := encodedHeader + "." + encodedPayload
	signature := ed25519.Sign(p.privateKey, []byte(signingInput))

	token := signingInput + "." + base64.RawURLEncoding.EncodeToString(signature)
	return token, time.Unix(exp, 0).UTC(), nil
}

func ParseEd25519PrivateKey(pemBytes []byte) (crypto.PrivateKey, error) {
	return parseEd25519PrivateKey(pemBytes)
}

func parseEd25519PrivateKey(pemBytes []byte) (ed25519.PrivateKey, error) {
	block, _ := pem.Decode(pemBytes)
	if block == nil {
		return nil, fmt.Errorf("qweather: invalid pem private key")
	}

	key, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("qweather: parse pkcs8 private key: %w", err)
	}

	privateKey, ok := key.(ed25519.PrivateKey)
	if !ok {
		return nil, fmt.Errorf("qweather: private key is not ed25519")
	}

	return privateKey, nil
}
