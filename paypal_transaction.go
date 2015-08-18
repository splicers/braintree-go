package braintree

type PaypalTransaction struct {
	PayerEmail string `xml:"payer_email,omitempty"`
}
