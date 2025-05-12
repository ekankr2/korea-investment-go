package services

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// GetAccountBalance 함수는 계좌 잔고를 조회합니다
func (s *KISService) GetAccountBalance(accountNo string) (map[string]interface{}, error) {
	// 토큰 확인 및 갱신
	if err := s.GetAccessToken(); err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%s/uapi/domestic-stock/v1/trading/inquire-balance", s.baseURL)

	// 요청 생성
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	// 쿼리 파라미터 추가
	q := req.URL.Query()
	q.Add("CANO", accountNo[:8])         // 계좌번호 앞 8자리
	q.Add("ACNT_PRDT_CD", accountNo[8:]) // 계좌상품코드
	q.Add("INQR_DVSN", "02")             // 조회구분(01: 대출일별, 02: 종목별)
	q.Add("UNIT_CLS", "01")              // 단위구분(01: 원화, 02: 달러)
	q.Add("AFHR_FLPR_YN", "Y")           // 앺장 포함여부(Y: 포함, N: 미포함)
	q.Add("OFL_YN", "N")                 // 오프라인여부(N: 미포함, Y: 포함)
	q.Add("UNPR_DVSN", "01")             // 단가 구분 - 01: 평균단가, 02: 기준가
	q.Add("FUND_STTL_ICLD_YN", "N")      // 펀드 결제분 포함 여부 - 일반적으로 "N"으로 설정
	q.Add("CTX_AREA_FK100", "")          // 연속조회검색조건100
	q.Add("CTX_AREA_NK100", "")          // 연속조회키100
	q.Add("FNCG_AMT_AUTO_RDPT_YN", "N")  // 융자금액 자동상환여부
	q.Add("PRCS_DVSN", "01")             // 처리구분 - 01: 조회, 02: 정정
	q.Add("COST_ICLD_YN", "N")           // 비용포함여부
	req.URL.RawQuery = q.Encode()

	// 헤더 설정
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", s.token))
	req.Header.Set("appkey", s.appKey)
	req.Header.Set("appsecret", s.appSecret)
	req.Header.Set("tr_id", "TTTC8434R") // 계좌 잔고 조회 TR ID
	req.Header.Set("custtype", "P")      // 고객타입 - 개인: P, 법인: B
	req.Header.Set("tr_cont", "")        // 연속거래 여부 - N: 초기, F: 최초, M: 중간, L: 마지막

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
