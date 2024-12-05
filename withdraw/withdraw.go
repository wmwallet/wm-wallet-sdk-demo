package withdraw

import (
	"github.com/wmwallet/wm-wallet-sdk/withdraw"
	"log"
	"wm-wallet-sdk-demo/thirdparty"
	"wm-wallet-sdk-demo/untils"
)

var wmWalletSDK = thirdparty.NewWmWalletSDK()

func Create(req *withdraw.CreateOrderReq) {
	log.Println("入参", untils.GetString(req))
	resp, err := wmWalletSDK.WithdrawOrderCreate(req)
	if err != nil {
		log.Println("出错", err)
		return
	}
	log.Println("出参", untils.GetString(resp))
}

func Detail(req *withdraw.GetDetailRequest) {
	log.Println("入参", untils.GetString(req))
	resp, err := wmWalletSDK.WithdrawOrderDetail(req)
	if err != nil {
		log.Println("出错", err)
		return
	}
	log.Println("出参", untils.GetString(resp))
}
