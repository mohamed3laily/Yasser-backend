package auth

import (
	"bytes"
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
	log.Printf("[DEBUG] Preparing to send WhatsApp OTP to %s with OTP: %s", phone, otp)
	log.Printf("[DEBUG] WhatsApp API URL: %s", w.config.APIURL)
	log.Printf("[DEBUG] WhatsApp Instance ID: %s", w.config.InstanceID)
	log.Printf("[DEBUG] WhatsApp Access Token: %s", w.config.AccessToken)
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

	log.Printf("[DEBUG] Sending request to: %s", w.config.APIURL)
	log.Printf("[DEBUG] Request body: %s", string(body))

	req, err := http.NewRequest("POST", w.config.APIURL, bytes.NewReader(body))
	if err != nil {
		log.Printf("[ERROR] Failed to create HTTP request: %v", err)
		return fmt.Errorf("create WhatsApp request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("[ERROR] Failed to send request: %v", err)
		return fmt.Errorf("send WhatsApp request: %w", err)
	}
	defer resp.Body.Close()

	respBody, _ := io.ReadAll(resp.Body)

	log.Printf("[DEBUG] WhatsApp API Response Status: %s", resp.Status)
	log.Printf("[DEBUG] WhatsApp API Response Body: %s", string(respBody))

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("WhatsApp request failed: status %d, body: %s", resp.StatusCode, string(respBody))
	}

	log.Printf("[SUCCESS] WhatsApp OTP sent successfully to %s", phone)
	return nil
}
