package wechatpay

import (
	"context"
	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/services/payments"
	"github.com/wechatpay-apiv3/wechatpay-go/services/payments/native"
)

type NativeService struct {
	appID string
	mchID string
	notifyURL string
	*native.NativeApiService
}


func (m *Client) NativeService()*NativeService  {
	return &NativeService{
		appID: m.appId,
		mchID: m.mchID,
		notifyURL: m.paymentNotifyURL,
		NativeApiService: &native.NativeApiService{
			Client: m.Client,
		},
	}
}

func (m *NativeService) PrepayPayment(request native.PrepayRequest) (*native.PrepayResponse, error) {
	request.Appid = core.String(m.appID)
	request.Mchid = core.String(m.mchID)
	request.NotifyUrl = core.String(m.notifyURL)
	// 得到prepay_id，以及调起支付所需的参数和签名
	resp, _, err := m.NativeApiService.Prepay(context.Background(), request)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
func (m *NativeService) QueryOrderById(transactionId string) (*payments.Transaction, error) {
	req := native.QueryOrderByIdRequest{
		TransactionId: core.String(transactionId),
		Mchid:         core.String(m.mchID),
	}
	trans, _, err := m.NativeApiService.QueryOrderById(context.Background(), req)
	return trans, err
}

func (m *NativeService) QueryOrderByOutTradeNo(outTradeNo string) (*payments.Transaction, error) {
	req := native.QueryOrderByOutTradeNoRequest{
		OutTradeNo: core.String(outTradeNo),
		Mchid:         core.String(m.mchID),
	}
	trans, _, err := m.NativeApiService.QueryOrderByOutTradeNo(context.Background(), req)
	return trans, err
}

func (m *NativeService) CloseOrder(outTradeNo string) error{
	req := native.CloseOrderRequest{
		OutTradeNo: core.String(outTradeNo),
		Mchid:         core.String(m.mchID),
	}
	_, err := m.NativeApiService.CloseOrder(context.Background(), req)
	return err
}
