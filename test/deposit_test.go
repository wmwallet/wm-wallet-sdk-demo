package test

import (
	"github.com/shopspring/decimal"
	"github.com/wmwallet/wm-wallet-sdk/deposit"
	"testing"
	localDeposit "wm-wallet-sdk-demo/deposit"
	"wm-wallet-sdk-demo/untils"
)

func TestDepositCreate(t *testing.T) {
	req := &deposit.CreateOrderReq{
		SourceId:   untils.GenerateUUID(),
		ChainId:    2,
		CoinId:     1,
		FiatAmount: decimal.NewFromInt(8),
		Symbol:     "USDT/CNY",
	}
	localDeposit.Create(req)
}

func TestDepositDetail(t *testing.T) {
	req := &deposit.GetDetailRequest{
		SourceId: "20241205195115111735",
	}
	localDeposit.Detail(req)
}

func TestDepositCancel(t *testing.T) {
	req := &deposit.CancelOrderReq{
		SourceId: "20241205195115111735",
	}
	localDeposit.Cancel(req)
}
