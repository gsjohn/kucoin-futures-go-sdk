package kumex

import (
	"testing"
	"time"
)

func TestApiService_Position(t *testing.T) {
	t.SkipNow()

	s := NewApiServiceFromEnv()
	rsp, err := s.Position("XBTUSDM")
	if err != nil {
		t.Fatal(err)
	}
	o := &PositionModelV1{}
	if err := rsp.ReadData(o); err != nil {
		t.Fatal(err)
	}
	t.Log(ToJsonString(o))
	switch {
	case o.Id == "":
		t.Error("Empty key 'id'")
	case o.Symbol == "":
		t.Error("Empty key 'symbol'")
	case o.MarkPrice == "":
		t.Error("Empty key 'markPrice'")
	case o.MarkValue == "":
		t.Error("Empty key 'markValue'")
	case o.SettleCurrency == "":
		t.Error("Empty key 'SettleCurrency'")
	}
}

func TestApiService_Positions(t *testing.T) {
	t.SkipNow()

	s := NewApiServiceFromEnv()
	rsp, err := s.Positions()
	if err != nil {
		t.Fatal(err)
	}
	o := &PositionModelV1{}
	if err := rsp.ReadData(o); err != nil {
		t.Fatal(err)
	}
	t.Log(ToJsonString(o))
	switch {
	case o.Id == "":
		t.Error("Empty key 'id'")
	case o.Symbol == "":
		t.Error("Empty key 'symbol'")
	case o.MarkPrice == "":
		t.Error("Empty key 'markPrice'")
	case o.MarkValue == "":
		t.Error("Empty key 'markValue'")
	case o.SettleCurrency == "":
		t.Error("Empty key 'SettleCurrency'")
	}
}

func TestApiService_autoDepositStatus(t *testing.T) {
	t.SkipNow()

	s := NewApiServiceFromEnv()
	p := map[string]string{
		"symbol": "XBTUSDM",
		"status": "true",
	}
	rsp, err := s.AutoDepositStatus(p)
	if err != nil {
		t.Fatal(err)
	}
	o := &PositionModelV1{}
	if err := rsp.ReadData(o); err != nil {
		t.Fatal(err)
	}
	t.Log(ToJsonString(o))
	switch {
	case o.Id == "":
		t.Error("Empty key 'id'")
	}
}

func TestApiService_DepositMargin(t *testing.T) {
	t.SkipNow()

	s := NewApiServiceFromEnv()
	p := map[string]string{
		"symbol": "XBTUSDM",
		"margin": "0.1111",
		"bizNo":  IntToString(time.Now().UnixNano()),
	}
	rsp, err := s.AutoDepositStatus(p)
	if err != nil {
		t.Fatal(err)
	}
	o := &PositionModelV1{}
	if err := rsp.ReadData(o); err != nil {
		t.Fatal(err)
	}
	t.Log(ToJsonString(o))
	switch {
	case o.Id == "":
		t.Error("Empty key 'id'")
	}
}
