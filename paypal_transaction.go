package braintree

type PaypalTransaction struct {
	PayerEmail string `xml:"payer-email,omitempty"`
	Token      string `xml:"token,omitempty"`
}
