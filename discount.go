package braintree

type DiscountList struct {
	XMLName   string     `xml:"discounts"`
	Discounts []Discount `xml:"discount"`
}

type Discount struct {
	XMLName string `xml:"discount"`
	Modification
	InheritedFromID string `xml:"inherited-from-id,omitempty"`
}

type Discounts struct {
	XMLName string      `xml:"discounts"`
	Add     interface{} `xml:"add,omitempty"`
	Update  interface{} `xml:"update,omitempty"`
	Remove  []string    `xml:"remove,omitempty"`
}
