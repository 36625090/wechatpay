package wechatpay

import (
	"context"
	"crypto/rsa"
	"fmt"
	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/core/option"
	"github.com/wechatpay-apiv3/wechatpay-go/utils"
)

type WechatPayClient interface {
	APPService() APPService
	NotifyService() (NotifyService, error)
	H5service() H5service
	NativeService() NativeService
	JSAPIService() JSAPIService
	RefundService() RefundService
}

type Client struct {
	*core.Client
	appId                      string
	mchID                      string
	mchAPIv3Key                string
	mchCertificateSerialNumber string
	paymentNotifyURL           string
	refundNotifyURL            string
	privateKey                 *rsa.PrivateKey
	publicKey                  *rsa.PublicKey
}

func NewWechatClient(cfg *Config) (WechatPayClient, error) {

	// 使用 utils 提供的函数从本地文件中加载商户私钥，商户私钥会用来生成请求的签名
	mchPrivateKey, err := utils.LoadPrivateKey(cfg.PrivateKey)
	if err != nil {
		return nil, fmt.Errorf("load private key failed: %v", err)
	}

	ctx := context.Background()
	// 使用商户私钥等初始化 client，并使它具有自动定时获取微信支付平台证书的能力
	opts := []core.ClientOption{
		option.WithWechatPayAutoAuthCipher(cfg.MchID, cfg.MchCertificateSerialNumber, mchPrivateKey, cfg.MchAPIv3Key),
	}

	client, err := core.NewClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("new wechat pay client err:%s", err)
	}

	wechatClient := &Client{
		Client:                     client,
		appId:                      cfg.AppID,
		mchID:                      cfg.MchID,
		mchAPIv3Key:                cfg.MchAPIv3Key,
		mchCertificateSerialNumber: cfg.MchCertificateSerialNumber,
		paymentNotifyURL:           cfg.PaymentNotifyURL,
		refundNotifyURL:            cfg.RefundNotifyURL,
		privateKey:                 mchPrivateKey,
	}

	if cfg.PublicKey != "" {
		mchPublicKey, err := utils.LoadPublicKey(cfg.PublicKey)
		if err != nil {
			return nil, fmt.Errorf("failed to load publicKey: %v", err)
		}
		wechatClient.publicKey = mchPublicKey
	}
	return wechatClient, nil
}
