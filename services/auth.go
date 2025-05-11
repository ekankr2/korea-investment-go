package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// GetAccessToken은 접근 토큰을 발급받습니다
func (s *KISService) GetAccessToken() error {
	// 토큰이 유효한 경우 재사용
	if s.token != "" && time.Now().Before(s.tokenExp) {
		return nil
	}

	url := fmt.Sprintf("%s/oauth2/tokenP", s.baseURL)

	// 요청 데이터 준비
	data := map[string]string{
		"grant_type": "client_credentials",
		"appkey":     s.appKey,
		"appsecret":  s.appSecret,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	// 요청 생성
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	// 요청 전송
	resp, err := s.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// 응답 처리
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		return err
	}

	// 토큰 저장
	s.token = result["access_token"].(string)
	expiresIn := int(result["expires_in"].(float64))
	s.tokenExp = time.Now().Add(time.Duration(expiresIn) * time.Second)

	return nil
}
