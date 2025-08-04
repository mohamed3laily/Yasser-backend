package auth

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type WhatsAppConfig struct {
	APIURL      string
	InstanceID  string
	AccessToken string
}

type WhatsAppSender struct {
	config WhatsAppConfig
}

func NewWhatsAppSender(cfg WhatsAppConfig) *WhatsAppSender {
	return &WhatsAppSender{config: cfg}
}

func (w *WhatsAppSender) SendOTP(phone, otp string) error {
	log.Printf("Sending WhatsApp OTP to %s: %s", phone, otp)

	payload := map[string]interface{}{
		"number":       phone,
		"type":         "text",
		"message":      fmt.Sprintf("Your verification code is: %s", otp),
		"instance_id":  w.config.InstanceID,
		"access_token": w.config.AccessToken,
	}

	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("marshal WhatsApp payload: %w", err)
	}

	req, err := http.NewRequest("POST", w.config.APIURL, bytes.NewReader(body))
	if err != nil {
		return fmt.Errorf("create WhatsApp request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("send WhatsApp request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("WhatsApp request failed: status %d", resp.StatusCode)
	}

	log.Printf("WhatsApp OTP sent successfully to %s", phone)
	return nil
}
