// services/kis_service.go
package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type KISService struct {
	client    *http.Client
	baseURL   string
	appKey    string
	appSecret string
	token     string
	tokenExp  time.Time
}

func NewKISService(appKey, appSecret string) *KISService {
	return &KISService{
		client:    &http.Client{Timeout: 10 * time.Second},
		baseURL:   "https://openapi.koreainvestment.com:9443", // 실전투자용 URL
		appKey:    appKey,
		appSecret: appSecret,
	}
}

func (s *KISService) GetAccessToken() error {
	if s.token != "" && time.Now().Before(s.tokenExp) {
		return nil
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

	return nil
}

func (s *KISService) GetStockPrice(stockCode string) (map[string]interface{}, error) {
	if err := s.GetAccessToken(); err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%s/uapi/domestic-stock/v1/quotations/inquire-price", s.baseURL)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Add("FID_COND_MRKT_DIV_CODE", "J") // 주식 시장 구분 코드
	q.Add("FID_INPUT_ISCD", stockCode)   // 주식 종목 코드
	req.URL.RawQuery = q.Encode()

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", s.token))
	req.Header.Set("appkey", s.appKey)
	req.Header.Set("appsecret", s.appSecret)
	req.Header.Set("tr_id", "FHKST01010100") // 주식 현재가 시세 조회 TR ID

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	return result, nil
}

func (s *KISService) GetAccountBalance(accountNo string) (map[string]interface{}, error) {
	if err := s.GetAccessToken(); err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%s/uapi/domestic-stock/v1/trading/inquire-balance", s.baseURL)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Add("CANO", accountNo[:8])         // 계좌번호 앞 8자리
	q.Add("ACNT_PRDT_CD", accountNo[8:]) // 계좌상품코드
	q.Add("INQR_DVSN", "02")             // 조회구분(01: 대출일별, 02: 종목별)
	q.Add("UNIT_CLS", "01")              // 단위구분(01: 원화, 02: 달러)
	req.URL.RawQuery = q.Encode()

	// 헤더 설정
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", s.token))
	req.Header.Set("appkey", s.appKey)
	req.Header.Set("appsecret", s.appSecret)
	req.Header.Set("tr_id", "TTTC8434R") // 계좌 잔고 조회 TR ID

	// 요청 전송
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// 응답 처리
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	return result, nil
}
