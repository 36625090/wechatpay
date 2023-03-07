package wechatpay

type Config struct {
	AppID                      string `json:"app_id" hcl:"app_id"`
	MchID                      string `json:"mch_id" hcl:"mch_id"`
	MchAPIv3Key                string `json:"mch_api_v3_key" hcl:"mch_api_v3_key"`
	MchCertificateSerialNumber string `json:"mch_certificate_serial_number" hcl:"mch_certificate_serial_number"`
	PaymentNotifyURL           string `json:"payment_notify_url" hcl:"payment_notify_url"`
	RefundNotifyURL            string `json:"refund_notify_url" hcl:"refund_notify_url"`
	PrivateKey                 string `json:"private_key" hcl:"private_key"`
	PublicKey                  string `json:"public_key" hcl:"public_key"`
}
