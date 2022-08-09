package kumex

import (
	"net/http"
)

// A PositionModel represents a position info.
type PositionModelV1 struct {
	Id                string `json:"id"`
	Symbol            string `json:"symbol"`
	AutoDeposit       bool   `json:"autoDeposit"`
	MaintMarginReq    string `json:"maintMarginReq"`
	RiskLimit         string `json:"riskLimit"`
	RealLeverage      string `json:"realLeverage"`
	CrossMode         bool   `json:"crossMode"`
	DelevPercentage   string `json:"delevPercentage"`
	OpeningTimestamp  string `json:"openingTimestamp"`
	CurrentTimestamp  string `json:"currentTimestamp"`
	CurrentQty        string `json:"currentQty"`
	CurrentCost       string `json:"currentCost"`
	CurrentComm       string `json:"currentComm"`
	UnrealisedCost    string `json:"unrealisedCost"`
	RealisedGrossCost string `json:"realisedGrossCost"`
	RealisedCost      string `json:"realisedCost"`
	IsOpen            bool   `json:"isOpen"`
	MarkPrice         string `json:"markPrice"`
	MarkValue         string `json:"markValue"`
	PosCost           string `json:"posCost"`
	PosCross          string `json:"posCross"`
	PosInit           string `json:"posInit"`
	PosComm           string `json:"posComm"`
	PosLoss           string `json:"posLoss"`
	PosMargin         string `json:"posMargin"`
	PosMaint          string `json:"posMaint"`
	MaintMargin       string `json:"maintMargin"`
	RealisedGrossPnl  string `json:"realisedGrossPnl"`
	RealisedPnl       string `json:"realisedPnl"`
	UnrealisedPnl     string `json:"unrealisedPnl"`
	UnrealisedPnlPcnt string `json:"unrealisedPnlPcnt"`
	UnrealisedRoePcnt string `json:"unrealisedRoePcnt"`
	AvgEntryPrice     string `json:"avgEntryPrice"`
	LiquidationPrice  string `json:"liquidationPrice"`
	BankruptPrice     string `json:"bankruptPrice"`
	SettleCurrency    string `json:"settleCurrency"`
	MaintainMargin    string `json:"maintainMargin"`
	RiskLimitLevel    string `json:"riskLimitLevel"`
}

type PositionModelV2 struct {
	Symbol                string  `json:"symbol"`
	Qty                   float64 `json:"qty"`
	Leverage              int     `json:"leverage"`
	MarginType            string  `json:"marginType"`
	Side                  string  `json:"side"`
	AutoDeposit           bool    `json:"autoDeposit"`
	EntryPrice            float64 `json:"entryPrice"`
	EntryValue            float64 `json:"entryValue"`
	Margin                float64 `json:"margin"`
	TotalMargin           float64 `json:"totalMargin"`
	LiquidationPrice      float64 `json:"liquidationPrice"`
	UnrealisedPnl         float64 `json:"unrealisedPnl"`
	MarkPrice             float64 `json:"markPrice"`
	RiskRate              float64 `json:"riskRate"`
	MaintenanceMarginRate float64 `json:"maintenanceMarginRate"`
	MaintenanceMargin     float64 `json:"maintenanceMargin"`
	AdlPercentile         int     `json:"adlPercentile"`
}

// Position Get Position Details.
func (as *ApiService) Position(symbol string) (*ApiResponse, error) {
	p := map[string]string{}
	if symbol != "" {
		p["symbol"] = symbol
	}
	req := NewRequest(http.MethodGet, "/symbol-position", p)
	return as.Call(req)
}

// Positions Get Position List.
func (as *ApiService) Positions() (*ApiResponse, error) {
	req := NewRequest(http.MethodGet, "/all-position", nil)
	return as.Call(req)
}

// AutoDepositStatus Enable/Disable of Auto-Deposit Margin.
func (as *ApiService) AutoDepositStatus(params map[string]string) (*ApiResponse, error) {
	req := NewRequest(http.MethodPost, "/position/margin/auto-deposit-status", params)
	return as.Call(req)
}

// DepositMargin Add Margin Manually.
func (as *ApiService) DepositMargin(params map[string]string) (*ApiResponse, error) {
	req := NewRequest(http.MethodPost, "/position/margin/deposit-margin", params)
	return as.Call(req)
}
