package kumex

import (
	"net/http"
	"strconv"
)

type LeverageModel struct {
	Symbol       string  `json:"symbol"`
	Leverage     int     `json:"leverage"`
	MaxRiskLimit float64 `json:"maxRiskLimit"`
}

// ChangeAutoDeposit  修改用户自动追加保证金状态
func (as *ApiService) ChangeAutoDeposit(params map[string]string) (*ApiResponse, error) {
	req := NewRequest(http.MethodPost, "/user-config/change-auto-deposit", params)
	return as.Call(req)
}

// AdjustLeverage  修改用户全局杠杆
func (as *ApiService) AdjustLeverage(symbol string, leverage int) (*ApiResponse, error) {
	req := NewRequest(http.MethodPost, "/user-config/adjust-leverage", map[string]string{
		"symbol":   symbol,
		"leverage": strconv.Itoa(leverage),
	})
	return as.Call(req)
}

func (as *ApiService) GetLeverage(symbol string) (*ApiResponse, error) {
	req := NewRequest(http.MethodGet, "/user-config/leverage", map[string]string{
		"symbol": symbol,
	})
	return as.Call(req)
}
