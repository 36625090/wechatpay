package wechatpay

import (
	"context"
	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/services/refunddomestic"
)

type RefundService struct {
	appID string
	mchID string
	notifyURL string
	*refunddomestic.RefundsApiService
}

func (m *Client) RefundService() *RefundService {
	return &RefundService{
		appID:             m.appId,
		mchID:             m.mchID,
		notifyURL:         m.refundNotifyURL,
		RefundsApiService: &refunddomestic.RefundsApiService{Client: m.Client},
	}
}

func (m *RefundService) Refund(req refunddomestic.CreateRequest)(*refunddomestic.Refund, error)  {
	req.NotifyUrl = core.String(m.notifyURL)
	resp, _, err := m.RefundsApiService.Create(context.Background(), req)
	return resp, err
}

func (m *RefundService)Query(req refunddomestic.QueryByOutRefundNoRequest)(*refunddomestic.Refund, error){
	resp, _, err := m.RefundsApiService.QueryByOutRefundNo(context.Background(), req)
	return resp, err
}
