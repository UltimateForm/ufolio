package turnstile

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/UltimateForm/ufolio/internal/config"
	"github.com/UltimateForm/ufolio/internal/rand"
)

var ErrUnexpectedStatus error = errors.New("unexpected status code")
var MAX_RETRIES int = 5

func VerifyToken(ctx context.Context, token string, idempotencyKey string, secretKey string) (*siteVerifyResponse, error) {

	body, err := json.Marshal(map[string]string{
		"secret":          secretKey,
		"response":        token,
		"idempotency_key": idempotencyKey,
	})
	if err != nil {
		return nil, err
	}
	log.Printf("sending token %s; body=%s\n", token, string(body))
	var httpResp *http.Response
	for i := range MAX_RETRIES {
		req, err := http.NewRequestWithContext(ctx, "POST", "https://challenges.cloudflare.com/turnstile/v0/siteverify", bytes.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
		if err != nil {
			return nil, err
		}
		httpResp, err = http.DefaultClient.Do(req)
		if err != nil {
			// breaking early on connectivity issues
			return nil, err
		}

		if httpResp.StatusCode == http.StatusOK {
			break
		}
		if i < MAX_RETRIES-1 {
			select {
			case <-ctx.Done():
				return nil, ctx.Err()
			case <-time.After(time.Duration(int(time.Second) * (i + 1))):
				continue
			}
		}
	}
	if httpResp == nil {
		return nil, errors.New("unexpected nil response")
	}
	defer httpResp.Body.Close()

	if httpResp.StatusCode != http.StatusOK {
		return nil, errors.Join(ErrUnexpectedStatus, errors.New(httpResp.Status))
	}

	var res siteVerifyResponse
	if err := json.NewDecoder(httpResp.Body).Decode(&res); err != nil {
		return nil, err
	}
	return &res, nil

}

var ErrInvalidToken error = errors.New("invalid token")

func ValidateToken(ctx context.Context, token string) error {
	idempotencyKey, err := rand.RandomStr(64)
	if err != nil {
		return err
	}
	res, err := VerifyToken(ctx, token, idempotencyKey, config.Secret.TurnstileSecretKey)
	if err != nil {
		return err
	}
	if !res.Success {
		return ErrInvalidToken
	}
	return nil
}
