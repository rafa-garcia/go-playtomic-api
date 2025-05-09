package models

// Payment represents payment information
type Payment struct {
	PaymentID               string  `json:"payment_id"`
	PaymentMethodID         string  `json:"payment_method_id"`
	PaymentMethodType       string  `json:"payment_method_type"`
	OnsitePaymentMethodType *string `json:"onsite_payment_method_type"`
	B2bBillingType          string  `json:"b2b_billing_type"`
	UserVat                 float64 `json:"user_vat"`
	TenantVat               float64 `json:"tenant_vat"`
	CommissionModel         string  `json:"commission_model"`
	RefundID                *string `json:"refund_id"`
	PaymentPrice            string  `json:"payment_price"`
	PaymentReference        *string `json:"payment_reference"`
	PayerID                 string  `json:"payer_id"`
	PaymentDate             string  `json:"payment_date"`
}
