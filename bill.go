package wechatpay

import (
	"github.com/wechatpay-apiv3/wechatpay-go/services/transferbatch"
)

type BillService interface {
}

type billService struct {
	appID     string
	mchID     string
	notifyURL string
	*transferbatch.TransferDetailApiService
}
