package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"korea-investment-go/lib/redis"
	"log"
	"net/http"
	"time"
)

func (s *KISService) GetAccessToken() error {
	tokenKey := fmt.Sprintf("token:kis:%s", s.appKey)

	tokenData, err := redis.Get(tokenKey)

	if err == nil && tokenData != "" {
		var tokenInfo map[string]interface{}
		if err := json.Unmarshal([]byte(tokenData), &tokenInfo); err == nil {
			s.token = tokenInfo["access_token"].(string)
			expStr := tokenInfo["expires_at"].(string)
			expTime, err := time.Parse(time.RFC3339, expStr)

			if err == nil && time.Now().Before(expTime) {
				s.tokenExp = expTime
				return nil
			}
		}
	}

	url := fmt.Sprintf("%s/oauth2/tokenP", s.baseURL)

	data := map[string]string{
		"grant_type": "client_credentials",
		"appkey":     s.appKey,
		"appsecret":  s.appSecret,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := s.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		return err
	}

	s.token = result["access_token"].(string)
	expiresIn := int(result["expires_in"].(float64))
	s.tokenExp = time.Now().Add(time.Duration(expiresIn) * time.Second)

	tokenInfo := map[string]interface{}{
		"access_token": s.token,
		"expires_at":   s.tokenExp.Format(time.RFC3339),
	}

	tokenJSON, err := json.Marshal(tokenInfo)

	expiration := time.Duration(expiresIn-60) * time.Second
	log.Printf("redis.Set 호출 직전: tokenKey=%s", tokenKey)

	err = redis.Set(tokenKey, string(tokenJSON), expiration)
	if err != nil {
		log.Printf("Failed to store token in Redis: %v", err)
	} else {
		log.Printf("Successfully stored token in Redis with key: %s", tokenKey)

		// Verify it was stored
		storedToken, err := redis.Get(tokenKey)
		if err != nil {
			log.Printf("Failed to verify token in Redis: %v", err)
		} else {
			log.Printf("Verified token in Redis: %s", storedToken)
		}
	}

	return nil
}
