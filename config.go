package wechatpay

type WechatClientConfig struct {
	AppID                      string
	MchID                      string
	MchAPIv3Key                string
	MchCertificateSerialNumber string
	PaymentNotifyURL           string
	RefundNotifyURL            string
	PrivateKey                 string
	PublicKey                  string
}
