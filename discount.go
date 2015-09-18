package braintree

type DiscountList struct {
	XMLName   string     `xml:"discounts"`
	Discounts []Discount `xml:"discount"`
}

type Discount struct {
	XMLName string `xml:"discount"`
	Modification
}

type AddedDiscount struct {
	XMLName         string `xml:"add"`
	InheritedFromID string `xml:"inherited-from-id,omitempty"`
	Quantity        int    `xml:"quantity"`
}

type UpdatedDiscount struct {
	XMLName         string `xml:"update"`
	InheritedFromID string `xml:"inherited-from-id,omitempty"`
	Quantity        int    `xml:"quantity"`
}

type Discounts struct {
	XMLName string             `xml:"discounts"`
	Add     []*AddedDiscount   `xml:"add,omitempty"`
	Update  []*UpdatedDiscount `xml:"update,omitempty"`
	Remove  []string           `xml:"remove,omitempty"`
}
