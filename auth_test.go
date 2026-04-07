package qweather

import (
	"context"
	"crypto/ed25519"
	"crypto/x509"
	"encoding/pem"
	"strings"
	"testing"
	"time"
)

func TestJWTTokenProviderGeneratesToken(t *testing.T) {
	t.Parallel()

	_, privateKey, err := ed25519.GenerateKey(nil)
	if err != nil {
		t.Fatalf("GenerateKey() error = %v", err)
	}

	der, err := x509.MarshalPKCS8PrivateKey(privateKey)
	if err != nil {
		t.Fatalf("MarshalPKCS8PrivateKey() error = %v", err)
	}

	privateKeyPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: der,
	})

	provider, err := NewJWTTokenProvider(JWTConfig{
		CredentialID:  "credential-id",
		ProjectID:     "project-id",
		PrivateKeyPEM: privateKeyPEM,
		TTL:           10 * time.Minute,
		Now: func() time.Time {
			return time.Unix(1_700_000_000, 0)
		},
	})
	if err != nil {
		t.Fatalf("NewJWTTokenProvider() error = %v", err)
	}

	token, err := provider.Token(context.Background())
	if err != nil {
		t.Fatalf("Token() error = %v", err)
	}

	if count := strings.Count(token, "."); count != 2 {
		t.Fatalf("unexpected token format: %q", token)
	}
}
