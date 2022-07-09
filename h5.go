package wechatpay

import (
	"context"
	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/services/payments"
	"github.com/wechatpay-apiv3/wechatpay-go/services/payments/h5"
)

type H5service struct {
	appID string
	mchID string
	notifyURL string
	*h5.H5ApiService
}
func (m *Client) H5service() *H5service {
	return &H5service{
		appID:         m.appId,
		mchID:         m.mchID,
		notifyURL:     m.paymentNotifyURL,
		H5ApiService: &h5.H5ApiService{Client: m.Client},
	}
}

func (m *H5service) PrepayPayment(request h5.PrepayRequest) (*h5.PrepayResponse, error) {
	request.Appid = core.String(m.appID)
	request.Mchid = core.String(m.mchID)
	request.NotifyUrl = core.String(m.notifyURL)
	// 得到prepay_id，以及调起支付所需的参数和签名
	resp, _, err := m.H5ApiService.Prepay(context.Background(), request)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
func (m *H5service) QueryOrderById(transactionId string) (*payments.Transaction, error) {
	req := h5.QueryOrderByIdRequest{
		TransactionId: core.String(transactionId),
		Mchid:         core.String(m.mchID),
	}
	trans, _, err := m.H5ApiService.QueryOrderById(context.Background(), req)
	return trans, err
}

func (m *H5service) QueryOrderByOutTradeNo(outTradeNo string) (*payments.Transaction, error) {
	req := h5.QueryOrderByOutTradeNoRequest{
		OutTradeNo: core.String(outTradeNo),
		Mchid:         core.String(m.mchID),
	}
	trans, _, err := m.H5ApiService.QueryOrderByOutTradeNo(context.Background(), req)
	return trans, err
}

func (m *H5service) CloseOrder(outTradeNo string) error{
	req := h5.CloseOrderRequest{
		OutTradeNo: core.String(outTradeNo),
		Mchid:         core.String(m.mchID),
	}
	_, err := m.H5ApiService.CloseOrder(context.Background(), req)
	return err
}
