package wechatpay

type Config struct {
	AppID                      string `json:"app_id"`
	MchID                      string `json:"mch_id"`
	MchAPIv3Key                string `json:"mch_api_v3_key"`
	MchCertificateSerialNumber string `json:"mch_certificate_serial_number"`
	PaymentNotifyURL           string `json:"payment_notify_url"`
	RefundNotifyURL            string `json:"refund_notify_url"`
	PrivateKey                 string `json:"private_key"`
	PublicKey                  string `json:"public_key"`
}
