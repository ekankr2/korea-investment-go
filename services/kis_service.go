package services

import (
	"net/http"
	"time"
)

// KISService는 한국투자증권 API 서비스의 공통 구조체입니다
type KISService struct {
	client    *http.Client
	baseURL   string
	appKey    string
	appSecret string
	token     string
	tokenExp  time.Time
}

// NewKISService는 KISService의 새 인스턴스를 생성합니다
func NewKISService(appKey, appSecret string) *KISService {
	return &KISService{
		client:    &http.Client{Timeout: 10 * time.Second},
		baseURL:   "https://openapi.koreainvestment.com:9443",
		appKey:    appKey,
		appSecret: appSecret,
	}
}
