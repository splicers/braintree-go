package braintree

type PaymentMethod struct {
	CreditCard         *CreditCard    `xml:"credit-card,omitempty"`
	CustomerId         string         `xml:"customer-id,omitempty"`
	PaymentMethodNonce string         `xml:"payment-method-nonce,omitempty"`
	PaypalAccount      *PaypalAccount `xml:"paypal-account,omitempty"`
	Token              string         `xml:"token,omitempty"`
}
