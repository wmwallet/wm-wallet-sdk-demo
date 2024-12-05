package deposit

import (
	"github.com/wmwallet/wm-wallet-sdk/deposit"
	"log"
	"wm-wallet-sdk-demo/thirdparty"
	"wm-wallet-sdk-demo/untils"
)

var wmWalletSDK = thirdparty.NewWmWalletSDK()

func Create(req *deposit.CreateOrderReq) {
	log.Println("入参", untils.GetString(req))
	resp, err := wmWalletSDK.DepositOrderCreate(req)
	if err != nil {
		log.Println("出错", err)
		return
	}
	log.Println("出参", untils.GetString(resp))
}

func Detail(req *deposit.GetDetailRequest) {
	log.Println("入参", untils.GetString(req))
	resp, err := wmWalletSDK.DepositOrderDetail(req)
	if err != nil {
		log.Println("出错", err)
		return
	}
	log.Println("出参", untils.GetString(resp))
}

func Cancel(req *deposit.CancelOrderReq) {
	log.Println("入参", untils.GetString(req))
	resp, err := wmWalletSDK.DepositOrderCancel(req)
	if err != nil {
		log.Println("出错", err)
		return
	}
	log.Println("出参", untils.GetString(resp))
}
