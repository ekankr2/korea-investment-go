package services

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
)

// 주문 요청 바디 구조체
type OrderRequest struct {
	CANO         string `json:"CANO"`
	ACNT_PRDT_CD string `json:"ACNT_PRDT_CD"`
	PDNO         string `json:"PDNO"`
	ORD_DVSN     string `json:"ORD_DVSN"`
	ORD_QTY      string `json:"ORD_QTY"`
	ORD_UNPR     string `json:"ORD_UNPR"`
}

func (s *KISService) OrderCash(req OrderRequest, orderType string) (map[string]interface{}, error) {
	if err := s.GetAccessToken(); err != nil {
		return nil, err
	}

	req.CANO = s.accountNum[:8]
	req.ACNT_PRDT_CD = "01"

	var trID string
	switch orderType {
	case "buy":
		trID = "TTTC0802U"
	case "sell":
		trID = "TTTC0801U"
	default:
		return nil, errors.New("invalid order type")
	}

	bodyBytes, _ := json.Marshal(req)
	apiURL := s.baseURL + "/uapi/domestic-stock/v1/trading/order-cash"
	httpReq, _ := http.NewRequest("POST", apiURL, bytes.NewBuffer(bodyBytes))
	httpReq.Header.Set("content-type", "application/json; charset=UTF-8")
	httpReq.Header.Set("authorization", "Bearer "+s.token)
	httpReq.Header.Set("appkey", s.appKey)
	httpReq.Header.Set("appsecret", s.appSecret)
	httpReq.Header.Set("tr_id", trID)

	client := &http.Client{}
	resp, err := client.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return result, nil
}
