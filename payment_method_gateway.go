package braintree

type PaymentMethodGateway struct {
	*Braintree
}

type pmCreateArgs struct {
	PaymentMethod *PaymentMethod `xml:"payment-method,omitempty"`
}

func (g *PaymentMethodGateway) Create(pm *PaymentMethod) (*PaymentMethod, error) {
	resp, err := g.execute("POST", "payment_methods", &pmCreateArgs{pm})
	if err != nil {
		return nil, err
	}
	switch resp.StatusCode {
	case 201:
		return resp.paymentMethod()
	}
	return nil, &invalidResponseError{resp}
}

func (g *PaymentMethodGateway) Update(pm *PaymentMethod) (*PaymentMethod, error) {
	resp, err := g.execute("PUT", "payment_methods/any/"+pm.Token, &pmCreateArgs{pm})
	if err != nil {
		return nil, err
	}
	switch resp.StatusCode {
	case 200:
		return resp.paymentMethod()
	}
	return nil, &invalidResponseError{resp}
}

func (g *PaymentMethodGateway) Find(token string) (*PaymentMethod, error) {
	resp, err := g.execute("GET", "payment_methods/any/"+token, nil)
	if err != nil {
		return nil, err
	}
	switch resp.StatusCode {
	case 200:
		return resp.paymentMethod()
	}
	return nil, &invalidResponseError{resp}
}

func (g *PaymentMethodGateway) Delete(pm *PaymentMethod) error {
	resp, err := g.execute("DELETE", "payment_methods/any/"+pm.Token, nil)
	if err != nil {
		return err
	}
	switch resp.StatusCode {
	case 200:
		return nil
	}
	return &invalidResponseError{resp}
}