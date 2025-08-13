package auth

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
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
	fmt.Printf(otp)
	if w.config.APIURL == "" || w.config.InstanceID == "" || w.config.AccessToken == "" {
		return fmt.Errorf("WhatsApp configuration is incomplete")
	}
	payload := map[string]interface{}{
		"number":       phone,
		"type":         "text",
		"message":      fmt.Sprintf("Your verification code is: %s", otp),
		"instance_id":  w.config.InstanceID,
		"access_token": w.config.AccessToken,
	}
	

	body, err := json.Marshal(payload)
	if err != nil {
		log.Printf("[ERROR] Failed to marshal payload: %v", err)
		return fmt.Errorf("marshal WhatsApp payload: %w", err)
	}

	req, err := http.NewRequest("POST", w.config.APIURL, bytes.NewReader(body))
	if err != nil {
		log.Printf("[ERROR] Failed to create HTTP request: %v", err)
		return fmt.Errorf("create WhatsApp request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{Timeout: 10 * time.Second, Transport: tr}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("[ERROR] Failed to send request: %v", err)
		return fmt.Errorf("send WhatsApp request: %w", err)
	}
	defer resp.Body.Close()

	respBody, _ := io.ReadAll(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("WhatsApp request failed: status %d, body: %s", resp.StatusCode, string(respBody))
	}

	return nil
}
