package wechatpay

import (
	"context"
	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/services/payments"
	"github.com/wechatpay-apiv3/wechatpay-go/services/payments/jsapi"
)

type JSAPIService struct {
	appID string
	mchID string
	notifyURL string
	*jsapi.JsapiApiService
}

func (m *Client) JSAPIService()*JSAPIService  {
	return &JSAPIService{
		appID: m.appId,
		mchID: m.mchID,
		notifyURL: m.paymentNotifyURL,
		JsapiApiService: &jsapi.JsapiApiService{
			Client: m.Client,

		},
	}
}

func (m *JSAPIService) PrepayPayment(request jsapi.PrepayRequest) (*jsapi.PrepayWithRequestPaymentResponse, error) {
	request.Appid = core.String(m.appID)
	request.Mchid = core.String(m.mchID)
	request.NotifyUrl = core.String(m.notifyURL)
	// 得到prepay_id，以及调起支付所需的参数和签名
	resp, _, err := m.JsapiApiService.PrepayWithRequestPayment(context.Background(), request)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (m *JSAPIService) QueryOrderById(transactionId string) (*payments.Transaction, error) {
	req := jsapi.QueryOrderByIdRequest{
		TransactionId: core.String(transactionId),
		Mchid:         core.String(m.mchID),
	}
	trans, _, err := m.JsapiApiService.QueryOrderById(context.Background(), req)
	return trans, err
}

func (m *JSAPIService) QueryOrderByOutTradeNo(outTradeNo string) (*payments.Transaction, error) {
	req := jsapi.QueryOrderByOutTradeNoRequest{
		OutTradeNo: core.String(outTradeNo),
		Mchid:         core.String(m.mchID),
	}
	trans, _, err := m.JsapiApiService.QueryOrderByOutTradeNo(context.Background(), req)
	return trans, err
}

func (m *JSAPIService) CloseOrder(outTradeNo string) error{
	req := jsapi.CloseOrderRequest{
		OutTradeNo: core.String(outTradeNo),
		Mchid:         core.String(m.mchID),
	}
	_, err := m.JsapiApiService.CloseOrder(context.Background(), req)
	return err
}

