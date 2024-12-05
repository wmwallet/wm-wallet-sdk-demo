package test

import (
	"github.com/shopspring/decimal"
	"github.com/wmwallet/wm-wallet-sdk/withdraw"
	"testing"
	"wm-wallet-sdk-demo/untils"
	localWithdraw "wm-wallet-sdk-demo/withdraw"
)

func TestWithdrawCreate(t *testing.T) {
	req := &withdraw.CreateOrderReq{
		SourceId: untils.GenerateUUID(),
		ChainId:  2,
		CoinId:   1,
		Address:  "UQAB0bw8P7UPqmDzx5cA8nz_4YWWMGNin6gD7prFpd3QACX3",
		Tag:      "",
		Amount:   decimal.NewFromInt(4),
	}
	localWithdraw.Create(req)
}

func TestWithdrawDetail(t *testing.T) {
	req := &withdraw.GetDetailRequest{
		SourceId: "20241205200150747164",
	}
	localWithdraw.Detail(req)
}
