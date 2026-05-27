package mail

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/UltimateForm/ufolio/internal/config"
)

const url string = "https://api.resend.com/emails"

var logger = log.New(log.Default().Writer(), "[resend] ", log.Default().Flags())

func SendEmail(ctx context.Context, email, subject, body string) error {
	logger.Println("sending email", email, subject, body)

	payload := struct {
		From    string `json:"from"`
		To      string `json:"to"`
		Subject string `json:"subject"`
		Html    string `json:"html"`
		ReplyTo string `json:"reply_to"`
	}{
		From:    config.Secret.ResendFromEmail,
		To:      config.Secret.ResendToEmail,
		Subject: subject,
		Html:    body,
		ReplyTo: email,
	}
	jsonBytes, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewReader(jsonBytes))
	if err != nil {
		return err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", config.Secret.ResendApiKey))
	req.Header.Add("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", res.StatusCode)
	}
	decoder := json.NewDecoder(res.Body)
	var result map[string]interface{}
	if err := decoder.Decode(&result); err != nil {
		return err
	}

	logger.Printf("response %+v\n", result)
	return nil
}
