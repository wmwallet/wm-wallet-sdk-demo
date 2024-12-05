package thirdparty

import (
	"context"
	sdk "github.com/wmwallet/wm-wallet-sdk"
	"github.com/wmwallet/wm-wallet-sdk/deposit"
	"github.com/wmwallet/wm-wallet-sdk/withdraw"
)

const CustomerName = "test"

const Url = "http://wmwallet.pro"

type WmWalletSDK struct {
	wmWalletSDK *sdk.WmWalletClient
}

func NewWmWalletSDK() *WmWalletSDK {
	ops := []sdk.Option{
		sdk.WithSecretPath("wallet/public_key.pem"),
		sdk.WithCertPath("wallet/ca.crt", "wallet/client.crt", "wallet/client.key"),
		sdk.WithCustomer(CustomerName),
	}
	wmWalletSDK, err := sdk.Init(ops...)
	if err != nil {
		panic(err)
	}
	return &WmWalletSDK{
		wmWalletSDK: wmWalletSDK,
	}
}

func (wws *WmWalletSDK) DepositOrderCreate(req *deposit.CreateOrderReq) (resp *deposit.CreateOrderResp, err error) {
	d := deposit.NewDeposit(wws.wmWalletSDK, Url)
	resp, err = d.Create(context.Background(), req)
	if err != nil {
		return
	}
	return
}

func (wws *WmWalletSDK) DepositOrderDetail(req *deposit.GetDetailRequest) (resp *deposit.GetDetailResponse, err error) {
	d := deposit.NewDeposit(wws.wmWalletSDK, Url)
	resp, err = d.Detail(context.Background(), req)
	if err != nil {
		return
	}
	return
}

func (wws *WmWalletSDK) DepositOrderCancel(req *deposit.CancelOrderReq) (resp *deposit.CancelOrderResp, err error) {
	d := deposit.NewDeposit(wws.wmWalletSDK, Url)
	resp, err = d.Cancel(context.Background(), req)
	if err != nil {
		return
	}
	return
}

func (wws *WmWalletSDK) WithdrawOrderCreate(req *withdraw.CreateOrderReq) (resp *withdraw.CreateOrderResp, err error) {
	w := withdraw.NewWithdraw(wws.wmWalletSDK, Url)
	resp, err = w.Create(context.Background(), req)
	if err != nil {
		return
	}
	return
}

func (wws *WmWalletSDK) WithdrawOrderDetail(req *withdraw.GetDetailRequest) (resp *withdraw.GetDetailResponse, err error) {
	w := withdraw.NewWithdraw(wws.wmWalletSDK, Url)
	resp, err = w.Detail(context.Background(), req)
	if err != nil {
		return
	}
	return
}
