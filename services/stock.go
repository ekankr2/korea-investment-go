package services

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// GetStockPrice 함수는 주식 시세를 조회합니다
func (s *KISService) GetStockPrice(stockCode string) (map[string]interface{}, error) {
	// 토큰 확인 및 갱신
	if err := s.GetAccessToken(); err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%s/uapi/domestic-stock/v1/quotations/inquire-price", s.baseURL)

	// 요청 생성
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	// 쿼리 파라미터 추가
	q := req.URL.Query()
	q.Add("FID_COND_MRKT_DIV_CODE", "J") // 주식 시장 구분 코드
	q.Add("FID_INPUT_ISCD", stockCode)   // 주식 종목 코드
	req.URL.RawQuery = q.Encode()

	// 헤더 설정
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", s.token))
	req.Header.Set("appkey", s.appKey)
	req.Header.Set("appsecret", s.appSecret)
	req.Header.Set("tr_id", "FHKST01010100") // 주식 현재가 시세 조회 TR ID

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

// GetDailyStockPrice 함수는 주식의 일/주/월 시세를 조회합니다
func (s *KISService) GetDailyStockPrice(stockCode, startDate, endDate, periodCode string) (map[string]interface{}, error) {
	// 토큰 확인 및 갱신
	if err := s.GetAccessToken(); err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%s/uapi/domestic-stock/v1/quotations/inquire-daily-price", s.baseURL)

	// 요청 생성
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	// 쿼리 파라미터 추가
	q := req.URL.Query()
	q.Add("FID_COND_MRKT_DIV_CODE", "UN")    // 조건 시장 분류 코드 J:KRX, NX:NXT, UN:통합
	q.Add("FID_INPUT_ISCD", stockCode)       // 주식 종목 코드
	q.Add("FID_PERIOD_DIV_CODE", periodCode) // 기간 분류 코드 (D:일봉, W:주봉, M:월봉)
	q.Add("FID_ORG_ADJ_PRC", "0")            // 기준가 조정 가격
	req.URL.RawQuery = q.Encode()

	// 헤더 설정
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", s.token))
	req.Header.Set("appkey", s.appKey)
	req.Header.Set("appsecret", s.appSecret)
	req.Header.Set("tr_id", "FHKST01010400") // 주식 일/주/월 시세 조회 TR ID
	req.Header.Set("custtype", "P")          // 고객타입 - 개인: P, 법인: B

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
