package wechatpay

import (
	"context"
	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/services/payments"
	"github.com/wechatpay-apiv3/wechatpay-go/services/payments/app"
)

type APPService struct {
	appID     string
	mchID     string
	notifyURL string
	*app.AppApiService
}

func (m *Client) APPService() *APPService {
	return &APPService{
		appID:         m.appId,
		mchID:         m.mchID,
		notifyURL:     m.paymentNotifyURL,
		AppApiService: &app.AppApiService{Client: m.Client},
	}
}

func (m *APPService) PrepayPayment(request app.PrepayRequest) (*app.PrepayWithRequestPaymentResponse, error) {
	request.Appid = core.String(m.appID)
	request.Mchid = core.String(m.mchID)
	request.NotifyUrl = core.String(m.notifyURL)
	// 得到prepay_id，以及调起支付所需的参数和签名
	resp, _, err := m.AppApiService.PrepayWithRequestPayment(context.Background(), request)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (m *APPService) QueryOrderById(transactionId string) (*payments.Transaction, error) {
	req := app.QueryOrderByIdRequest{
		TransactionId: core.String(transactionId),
		Mchid:         core.String(m.mchID),
	}
	trans, _, err := m.AppApiService.QueryOrderById(context.Background(), req)
	return trans, err
}

func (m *APPService) QueryOrderByOutTradeNo(outTradeNo string) (*payments.Transaction, error) {
	req := app.QueryOrderByOutTradeNoRequest{
		OutTradeNo: core.String(outTradeNo),
		Mchid:      core.String(m.mchID),
	}
	trans, _, err := m.AppApiService.QueryOrderByOutTradeNo(context.Background(), req)
	return trans, err
}

func (m *APPService) CloseOrder(outTradeNo string) error {
	req := app.CloseOrderRequest{
		OutTradeNo: core.String(outTradeNo),
		Mchid:      core.String(m.mchID),
	}
	_, err := m.AppApiService.CloseOrder(context.Background(), req)
	return err
}
