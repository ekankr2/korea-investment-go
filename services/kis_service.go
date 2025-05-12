package services

import (
	"korea-investment-go/config"
	"net/http"
	"time"
)

type KISService struct {
	client     *http.Client
	baseURL    string
	appKey     string
	appSecret  string
	token      string
	tokenExp   time.Time
	accountNum string
}

func NewKISService() *KISService {
	cfg := config.GetConfig()

	return &KISService{
		client:     &http.Client{Timeout: 10 * time.Second},
		baseURL:    "https://openapi.koreainvestment.com:9443",
		appKey:     cfg.AppKey,
		appSecret:  cfg.AppSecret,
		accountNum: cfg.AccountNum,
	}
}
